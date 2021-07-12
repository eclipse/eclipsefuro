package reqres

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type AST struct {
	Config            map[string]interface{}            `yaml:"config"` // contains the config of the spec project, this can be relevant
	InstalledServices map[string]*serviceAst.ServiceAst `yaml:"installedServices"`
	InstalledTypes    map[string]*typeAst.TypeAst       `yaml:"installedTypes"`
	Services          map[string]*serviceAst.ServiceAst `yaml:"services"`
	Types             map[string]*typeAst.TypeAst       `yaml:"types"`
	AllTypes          map[string]*typeAst.TypeAst       // build this after receiving
	AllServices       map[string]*serviceAst.ServiceAst // build this after receiving

}

type Request struct {
	Parameters   []string
	ParameterMap map[string]string
	AST          AST
	Debug        bool
}

// print what you want to the stderr console and not to the reqres
func (Request) Fprintln(i ...interface{}) {
	fmt.Fprintln(os.Stderr, i...)
}

// Does all the input handling, and marshalling for you
//
// To enable debuging add the arguments debug and debugfile to the call of this command, without furoc
// "debug debugfile=./sample/fullyaml.yaml"
//
// To create a debug file use the argument "debugfileout=./sample/fullyaml.yaml" with furoc
func NewRequester() (*Request, *Response) {
	req := &Request{}
	req.Parameters = os.Args
	req.ParameterMap = map[string]string{}

	res := NewResponser()
	req.Debug = false

	// make a param map for easy access
	// you still have to parse it
	for _, p := range req.Parameters[1:] {
		kv := strings.Split(p, "=")
		if len(kv) > 1 {
			req.ParameterMap[kv[0]] = kv[1]
		} else {
			req.ParameterMap[p] = p
		}
	}

	var data []byte
	var err error
	// this is for debuging your generator
	// call it with your-cmd debug debugfile=path/to/debuginput
	if _, debug := req.ParameterMap["debug"]; debug {
		var debugFile string

		if f, ok := req.ParameterMap["debugfile"]; ok {
			debugFile = f
		}
		data, err = ioutil.ReadFile(debugFile)

		if err != nil {
			log.Fatal(debugFile, err)
		}
		req.Debug = true
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
		// write debugfile
		if f, ok := req.ParameterMap["debugfileout"]; ok {
			ioutil.WriteFile(f, data, 0644)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	err = yaml.Unmarshal([]byte(data), &req.AST) //reads yaml and json because json is just a subset of yaml
	if err != nil {
		log.Fatal(err)
	}

	// build allTypes
	req.AST.AllTypes = map[string]*typeAst.TypeAst{}
	for n, t := range req.AST.Types {
		req.AST.AllTypes[n] = t
	}
	for n, t := range req.AST.InstalledTypes {
		req.AST.AllTypes[n] = t
	}
	// build allServices
	req.AST.AllServices = map[string]*serviceAst.ServiceAst{}
	for n, t := range req.AST.Services {
		req.AST.AllServices[n] = t
	}
	for n, t := range req.AST.InstalledServices {
		req.AST.AllServices[n] = t
	}

	return req, res
}
