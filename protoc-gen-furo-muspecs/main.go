package main

import (
	"github.com/eclipse/eclipsefuro/protoc-gen-furo-muspecs/pkg/generator"
	"github.com/eclipse/eclipsefuro/protoc-gen-furo-muspecs/pkg/protoast"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"os"
)

func main() {
	// os.Stdin will contain data which will unmarshal into the following object:
	// https://godoc.org/github.com/golang/protobuf/protoc-gen-go/plugin#CodeGeneratorRequest
	req := &pluginpb.CodeGeneratorRequest{}

	// disable next line to debug from file instead of stdin
	data, err := ioutil.ReadAll(os.Stdin)

	// enable next line to save stdin to a file. This file can be used for debugging.
	// ioutil.WriteFile("protocdata", data, 755)

	// debug mode
	// enable next line to read the file instead using stdin.
	// data, err := ioutil.ReadFile("sample/protos/protocdata")

	if err != nil {
		panic(err)
	}

	proto.Unmarshal(data, req)

	if err != nil {
		panic(err)
	}

	Ast := protoast.NewProtoAST(req)

	err = generator.Generate(Ast)
	if err != nil {
		panic(err)
	}

	marshalled, err := proto.Marshal(Ast.Response)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
