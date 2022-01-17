package enumAst

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Enumlist struct {
	EnumsByName          map[string]*EnumAst
	InstalledEnumsByName map[string]*EnumAst
	AllAvailabeEnums     map[string]*EnumAst // installed and project enums together
	SpecDir              string
}

type EnumAst struct {
	SpecDir  string // the base path, like specDir or dependencies/x.y.com/specDir
	Path     string // relative path of spec file to SpecDir
	FileName string
	Enumspec specSpec.Enum
}

var Format = "json"

// set the storage format
func (l *Enumlist) setStorageFormat(format string) {
	Format = format
}

// loads a spec directory and installed specs to the typelist
func (l *Enumlist) LoadEnumSpecsFromDir(specDir string) {
	l.EnumsByName = loadEnumSpecsFromDir(specDir)

	if l.AllAvailabeEnums == nil {
		l.AllAvailabeEnums = map[string]*EnumAst{}
	}
	for k, t := range l.EnumsByName {
		l.AllAvailabeEnums[k] = t
	}
}

// loads a spec directory and installed specs to the typelist
func (l *Enumlist) LoadInstalledEnumSpecsFromDir(specDir ...string) {
	// create map if it does not exist
	if l.InstalledEnumsByName == nil {
		l.InstalledEnumsByName = map[string]*EnumAst{}
	}
	for _, dir := range specDir {
		tlist := loadEnumSpecsFromDir(dir)
		for tname, v := range tlist {
			l.InstalledEnumsByName[tname] = v
		}
	}

	if l.AllAvailabeEnums == nil {
		l.AllAvailabeEnums = map[string]*EnumAst{}
	}
	for k, t := range l.InstalledEnumsByName {
		l.AllAvailabeEnums[k] = t
	}
}

func loadEnumSpecsFromDir(specDir string) (enumsMap map[string]*EnumAst) {
	enumsMap = map[string]*EnumAst{}
	err := filepath.Walk(specDir,
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(fpath, "enum.spec") {
				filename := path.Base(fpath)
				sdlen := len(strings.Split(specDir, "/"))
				if strings.HasPrefix(specDir, "./") {
					sdlen--
				}

				relativePath := path.Dir(strings.Join(strings.Split(fpath, "/")[sdlen:], "/"))

				AstType := &EnumAst{
					Path:     relativePath, // store Path without specDir
					SpecDir:  specDir,      // store Path without specDir
					FileName: filename,
					Enumspec: readAndUnmarshalSpec(fpath),
				}

				enumsMap[strings.Join([]string{AstType.Enumspec.XProto.Package, AstType.Enumspec.Type}, ".")] = AstType
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return enumsMap
}

func readAndUnmarshalSpec(fpath string) (s specSpec.Enum) {
	dataBytes, readError := ioutil.ReadFile(fpath)
	if readError != nil {
		log.Fatal(readError)
	}
	parseError := yaml.Unmarshal(dataBytes, &s) //reads yaml and json because json is just a subset of yaml
	if parseError != nil {
		fmt.Println(fpath + ":1:1")
		log.Fatal(parseError)
	}

	if s.Values != nil {
		// convert fields from yaml.Node to Field type
		for pair := s.Values.Oldest(); pair != nil; pair = pair.Next() {
			fieldYamlNode := pair.Value.(*yaml.Node)
			var AstField *uint32
			fieldYamlNode.Decode(&AstField)
			pair.Value = AstField
		}
	}

	return s
}

// Stores the spec to disc
func (a *EnumAst) Save(specDir string) {
	filepath := path.Dir(path.Join(specDir, a.Path, a.FileName))
	filename := path.Join(filepath, a.FileName)

	var d []byte
	var err error
	switch Format {
	case "json":
		d, err = a.ToJson()
		break
	case "yaml":
		d, err = a.ToYaml()
		break
	}
	if err != nil {
		panic(err)
	}
	util.MkdirRelative(filepath)
	err = ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		panic(err)
	}
}

func (a *EnumAst) ToJson() (d []byte, err error) {
	d, err = json.MarshalIndent(a.Enumspec, "", " ")
	return d, err
}

// returns unindented json
func (a *EnumAst) ToJsonFlat() (d []byte, err error) {
	d, err = json.Marshal(a.Enumspec)
	return d, err
}

func (a *EnumAst) ToYaml() (d []byte, err error) {
	d, err = yaml.Marshal(&a.Enumspec)
	return d, err
}

// stores the typelist to the spec directory
func (l *Enumlist) SaveAllEnumSpecsToDir(specDir string) {
	for _, typeAst := range l.EnumsByName {
		typeAst.Save(specDir)
	}
}

//
func (l *Enumlist) ResolveProtoImportForType(typeName string, pkg string) (imp string, typeFound bool) {
	if strings.HasPrefix(typeName, "stream ") {
		typeName = typeName[7:]
	}
	if strings.HasPrefix(typeName, "[] ") {
		typeName = typeName[3:]
	}

	fqTypeName := l.ResolveFullQualifiedTypeName(typeName, pkg)
	// remove leading dot
	if strings.HasPrefix(fqTypeName, ".") {
		fqTypeName = fqTypeName[1:len(fqTypeName)]
	}

	//check on installed and spec tpelist
	imp = ""
	if l.EnumsByName[fqTypeName] == nil && l.InstalledEnumsByName[fqTypeName] == nil {
		return imp, false
	}
	if l.EnumsByName[fqTypeName] != nil {
		imp = l.EnumsByName[fqTypeName].GetProtoTarget()
	}
	if l.InstalledEnumsByName[fqTypeName] != nil {
		imp = l.InstalledEnumsByName[fqTypeName].GetProtoTarget()
	}
	return imp, true
}

func (a *EnumAst) GetProtoTarget() (proto string) {
	protoFile := a.Enumspec.XProto.Targetfile
	return path.Join(a.Path, protoFile)
}

// Deletes the spec from disk and removes the element from List
func (l *Enumlist) DeleteType(typename string) {
	// delete the file
	filepath := path.Join(viper.GetString("specDir"), l.EnumsByName[typename].Path, l.EnumsByName[typename].FileName)
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DELETED", filepath)
	}

	delete(l.EnumsByName, typename)

}

// full name from c++ style name
func (l Enumlist) ResolveFullQualifiedTypeName(typename string, pkg string) string {
	// absolut type given, nothing special to do
	if strings.HasPrefix(typename, ".") {
		// type starts from root, just remove the .
		return typename[1:len(typename)]
	}

	pathArr := strings.Split(pkg, ".")
	// if we are in type a.b.c.d and want type x.y we look for
	// a.b.c.d.x.y
	// a.b.c.x.y
	// a.b.x.y
	// a.x.y
	// x.y
	for i := len(pathArr); i >= 0; i-- {
		sub := strings.Join(pathArr[0:i], ".")
		ftype := sub + "." + typename

		if l.AllAvailabeEnums[ftype] != nil {
			// match
			return ftype
			i = 0
		}
		// we are at root
		if i == 0 && strings.HasPrefix(ftype, ".") {
			// remove .
			return ftype[1:len(ftype)]
		}
	}

	return typename
}
