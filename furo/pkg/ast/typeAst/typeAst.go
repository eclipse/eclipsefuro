package typeAst

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/enumAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Typelist struct {
	TypesByName          map[string]*TypeAst
	InstalledTypesByName map[string]*TypeAst
	AllAvailabeTypes     map[string]*TypeAst // installed and project types together
	SpecDir              string
}

type TypeAst struct {
	SpecDir  string // the base path, like specDir or dependencies/x.y.com/specDir
	Path     string // relative path of spec file to SpecDir
	FileName string
	TypeSpec specSpec.Type
}

var Format = "json"

// set the storage format
func (l *Typelist) setStorageFormat(format string) {
	Format = format
}

// loads a spec directory and installed specs to the typelist
func (l *Typelist) LoadTypeSpecsFromDir(specDir string) {
	l.TypesByName = loadTypeSpecsFromDir(specDir)

	if l.AllAvailabeTypes == nil {
		l.AllAvailabeTypes = map[string]*TypeAst{}
	}
	for k, t := range l.TypesByName {
		l.AllAvailabeTypes[k] = t
	}
}

// loads a spec directory and installed specs to the typelist
func (l *Typelist) LoadInstalledTypeSpecsFromDir(specDir ...string) {
	// create map if it does not exist
	if l.InstalledTypesByName == nil {
		l.InstalledTypesByName = map[string]*TypeAst{}
	}
	for _, dir := range specDir {
		tlist := loadTypeSpecsFromDir(dir)
		for tname, v := range tlist {
			l.InstalledTypesByName[tname] = v
		}
	}

	if l.AllAvailabeTypes == nil {
		l.AllAvailabeTypes = map[string]*TypeAst{}
	}
	for k, t := range l.InstalledTypesByName {
		l.AllAvailabeTypes[k] = t
	}
}

func loadTypeSpecsFromDir(specDir string) (typesMap map[string]*TypeAst) {
	typesMap = map[string]*TypeAst{}
	err := filepath.Walk(specDir,
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(fpath, "type.spec") {
				filename := path.Base(fpath)
				sdlen := len(strings.Split(specDir, "/"))
				if strings.HasPrefix(specDir, "./") {
					sdlen--
				}

				relativePath := path.Dir(strings.Join(strings.Split(fpath, "/")[sdlen:], "/"))

				AstType := &TypeAst{
					Path:     relativePath, // store Path without specDir
					SpecDir:  specDir,      // store Path without specDir
					FileName: filename,
					TypeSpec: readAndUnmarshalSpec(fpath),
				}

				typesMap[strings.Join([]string{AstType.TypeSpec.XProto.Package, AstType.TypeSpec.Type}, ".")] = AstType
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return typesMap
}

func readAndUnmarshalSpec(fpath string) (s specSpec.Type) {
	dataBytes, readError := ioutil.ReadFile(fpath)
	if readError != nil {
		log.Fatal(readError)
	}
	parseError := yaml.Unmarshal(dataBytes, &s) //reads yaml and json because json is just a subset of yaml
	if parseError != nil {
		fmt.Println(fpath + ":1:1")
		log.Fatal(parseError)
	}

	if s.Fields != nil {
		// convert fields from yaml.Node to Field type
		for pair := s.Fields.Oldest(); pair != nil; pair = pair.Next() {
			fieldYamlNode := pair.Value.(*yaml.Node)
			var AstField *specSpec.Field
			fieldYamlNode.Decode(&AstField)
			pair.Value = AstField
		}
	}

	return s
}

// Stores the spec to disc
func (a *TypeAst) Save(specDir string) {
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

func (a *TypeAst) ToJson() (d []byte, err error) {
	d, err = json.MarshalIndent(a.TypeSpec, "", " ")
	return d, err
}

// returns unindented json
func (a *TypeAst) ToJsonFlat() (d []byte, err error) {
	d, err = json.Marshal(a.TypeSpec)
	return d, err
}

func (a *TypeAst) ToYaml() (d []byte, err error) {
	d, err = yaml.Marshal(&a.TypeSpec)
	return d, err
}

// stores the typelist to the spec directory
func (l *Typelist) SaveAllTypeSpecsToDir(specDir string) {
	for _, typeAst := range l.TypesByName {
		typeAst.Save(specDir)
	}
}

//
func (l *Typelist) ResolveProtoImportForType(typeName string, pkg string) (imp string, typeFound bool) {
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
	if l.TypesByName[fqTypeName] == nil && l.InstalledTypesByName[fqTypeName] == nil {
		return imp, false
	}
	if l.TypesByName[fqTypeName] != nil {
		imp = l.TypesByName[fqTypeName].GetProtoTarget()
	}
	if l.InstalledTypesByName[fqTypeName] != nil {
		imp = l.InstalledTypesByName[fqTypeName].GetProtoTarget()
	}
	return imp, true
}

func (a *TypeAst) GetProtoTarget() (proto string) {
	protoFile := a.TypeSpec.XProto.Targetfile
	return path.Join(a.Path, protoFile)
}

// updates the imports on each type
func (l *Typelist) UpdateImports(enumlist *enumAst.Enumlist) {
	for t, v := range l.TypesByName {
		self, _ := l.ResolveProtoImportForType(t, v.TypeSpec.XProto.Package)
		imports := []string{}
		v.TypeSpec.Fields.Map(func(iFieldname interface{}, iField interface{}) {
			field := iField.(*specSpec.Field)
			typeToImport := field.Type
			// map imports
			if strings.HasPrefix(typeToImport, "map") {
				regex := regexp.MustCompile(`,([^>]*)`)
				matches := regex.FindStringSubmatch(typeToImport)
				typeToImport = strings.TrimSpace(matches[1])
			}

			// string, uint,... does not need to be imported
			f := strings.Split(typeToImport, ".")
			// todo: do a better check, sometimes typos exists. stirng is not a valid type ==> string
			if len(f) > 1 {
				i, found := l.ResolveProtoImportForType(typeToImport, v.TypeSpec.XProto.Package)
				if found && i != self {
					imports = append(imports, i)
				} else {
					if i != self {
						en, enFound := enumlist.ResolveProtoImportForType(typeToImport, v.TypeSpec.XProto.Package)
						if enFound && en != self {
							imports = append(imports, en)
						} else {
							if en != self {
								fmt.Println(util.ScanForStringPosition(typeToImport, path.Join(viper.GetString("specDir"), l.TypesByName[t].Path, l.TypesByName[t].FileName))+":Import", typeToImport, "not found in type", t, "on field", iFieldname)
								fmt.Println(util.ScanForStringPosition(typeToImport, viper.GetString("muSpec.types"))+":Import not found. Check your muSpec types if you came from there. Field:", iFieldname)
							}
						}

					}
				}
			}

		})
		// remove duplicate imports and sort them alphabetical
		imports = distinctStringArray(imports)
		sort.Strings(imports)
		v.TypeSpec.XProto.Imports = imports
	}
}
func distinctStringArray(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Deletes the spec from disk and removes the element from List
func (l *Typelist) DeleteType(typename string) {
	// delete the file
	filepath := path.Join(viper.GetString("specDir"), l.TypesByName[typename].Path, l.TypesByName[typename].FileName)
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DELETED", filepath)
	}

	delete(l.TypesByName, typename)

}

// full name from c++ style name
func (l Typelist) ResolveFullQualifiedTypeName(typename string, pkg string) string {
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

		if l.AllAvailabeTypes[ftype] != nil {
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
