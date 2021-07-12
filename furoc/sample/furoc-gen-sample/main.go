package main

import (
	furoc "github.com/eclipse/eclipsefuro/furoc/pkg/reqres"
	"gopkg.in/yaml.v3"
	"log"
)

func main() {

	// receive the request, all the stdin and mapping stuff is done here
	// if you need to debug your plugin, start it using following arguments:
	// debug debugfile=./sample/fullyaml.yaml
	// to create a debugfile
	// To create a debug file use "debugfileout=./sample/fullyaml.yaml" as an argument
	req, res := furoc.NewRequester()

	// create a responser, which can be used to add files and send the response back to furoc

	// use req.Fprintln(interface{}) if you want to print something to the console (or write to stderr)
	// you can not write to the console with fmt or log (because this goes to stdout)
	req.Fprintln("Sample plugin started")

	// the req object  contains
	//	Parameters   []string  a list of the given input parameters
	//	ParameterMap map[string]string the input parameters transformed to a map
	//	AST          AST

	for name, s := range req.AST.Services {

		// Using your own extension
		// when you have the custom extension "sampleExtension" in the service spec
		//
		//         extensions:
		//            sampleExtension:
		//                generate: sample
		// you can decode its content with furoc.DecodeExtension
		ext := &MyServiceSpecExtension{}
		furoc.DecodeExtension(s.ServiceSpec.Extensions, "sampleExtension", ext)

		// do something if generate was set in the extension
		if ext.Generate {

			// build your file
			fileContent, err := yaml.Marshal(s.ServiceSpec)
			if err != nil {
				log.Fatal(err)
			}

			// if your plugin needs to call another executable, you can use commandpipe.NewCommand()
			// create sample file
			readme := furoc.TargetFile{
				Filename: "/" + name + "/" + s.ServiceSpec.Name + ".md", // full qualified filename which will generated in :outputdir/
				Content:  fileContent,                                   //[]byte with content
			}

			// Add file to the responder
			res.AddFile(&readme)
		}
	}

	// send the response back to furoc
	res.SendResponse(req.Debug)
}

type MyServiceSpecExtension struct {
	Generate bool `yaml:"generate"` //field must be public
}
