package reqres

// use this to make a reqres for furoc

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type TargetFile struct {
	Filename string // full qualified name from out dir. You can start with /
	Content  []byte // the file content
}

type Response pluginpb.CodeGeneratorResponse

// Creates a responser which holds the files you want to send back to protoc.
func NewResponser() *Response {
	return &Response{
		File: []*pluginpb.CodeGeneratorResponse_File{},
	}
}

// Add a file to the responser, duplicate filename checks are done in furoc.
func (r *Response) AddFile(file *TargetFile) {

	cntnt := string(file.Content)
	// create sample file
	f := &pluginpb.CodeGeneratorResponse_File{
		Name:    &file.Filename, // full qualified filename which will generated in :outputdir/
		Content: &cntnt,
	}

	r.File = append(r.File, f)
}

// Send the encoded message response back to furoc
func (r *Response) SendResponse(debug bool) {
	if debug {
		// do the writes directly when debuging is enabled
		for _, file := range r.File {
			if util.DirExists("debug_out") {
				fname := path.Join("debug_out", *file.Name)
				util.MkdirRelative(path.Dir(fname))
				ioutil.WriteFile(fname, []byte(*file.Content), 0644)
			} else {
				log.Fatal("Dir does not exist: ", "debug_out")
			}

		}
	} else {

		// encode and send the reqres
		res := pluginpb.CodeGeneratorResponse(*r)
		marshalled, err := proto.Marshal(&res)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(marshalled)
	}
}
