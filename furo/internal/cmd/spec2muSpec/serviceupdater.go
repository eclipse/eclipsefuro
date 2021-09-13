package spec2muSpec

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/microservices"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"strings"
)

func updateAndStoreMicroServices(serviceItems map[string]*UTshadowNode) {
	muServicesPerFile := map[string][]*microservices.MicroService{}

	for _, shadowNode := range serviceItems {
		if shadowNode.edgeServiceNode != nil {

			if shadowNode.edgeMicroServiceNode == nil {
				// create muType because it does not exist
				shadowNode.edgeMicroServiceNode = &microservices.MicroServiceAst{
					SourceFile: viper.GetString("muSpec.dir") + "/" + shadowNode.edgeServiceNode.ServiceSpec.XProto.Package + "/" + shadowNode.edgeServiceNode.ServiceSpec.Name + ".services.yaml",
				}
			}

			if muServicesPerFile[shadowNode.edgeMicroServiceNode.SourceFile] == nil {
				muServicesPerFile[shadowNode.edgeMicroServiceNode.SourceFile] = []*microservices.MicroService{}
			}

			var serviceList []microservices.MicroRPC
			shadowNode.edgeServiceNode.ServiceSpec.Services.Map(func(iKey interface{}, iValue interface{}) {
				mRpc := microservices.MicroRPC{}
				rpc := iValue.(*specSpec.Rpc)
				mdLine := []string{}
				mdLine = append(mdLine, iKey.(string)+":")
				mdLine = append(mdLine, rpc.Deeplink.Method)
				mdLine = append(mdLine, rpc.Deeplink.Href)
				if rpc.Data.Request == "" {
					rpc.Data.Request = "google.protobuf.Empty"
				}
				mdLine = append(mdLine, rpc.Data.Request)
				mdLine = append(mdLine, ",")

				mdLine = append(mdLine, rpc.Data.Response)
				mdLine = append(mdLine, "#"+rpc.Description)

				mRpc.Md = strings.Join(mdLine, " ") //List: GET /samples google.protobuf.Empty, sample.SampleCollection #List samples with pagination.

				if shadowNode.edgeRequestTypeNode != nil {
					// we have qp
					qplist := orderedmap.New()

					// find the correct entry in shadowNode.edgeRequestTypeNode
					reqType := findRequestType(rpc.RpcName, shadowNode)
					reqType.TypeSpec.Fields.Map(func(iKey interface{}, iValue interface{}) {
						f := iValue.(*specSpec.Field) //*string:1 # A * before the type means required
						// ignore the Body field
						if rpc.Data.Bodyfield != iKey.(string) {
							fieldline := []string{}

							// maybe one day we want more on the query paramy??
							/*if f.Constraints["required"] != nil {
								fieldline = append(fieldline, "*")
							}
							if f.Meta.Readonly {
								fieldline = append(fieldline, "-")
							}
							if f.Meta.Repeated {
								fieldline = append(fieldline, "[]")
							}
							fieldline = append(fieldline, f.Type+":"+strconv.Itoa(int(f.XProto.Number)))
							/*
							*/
							fieldline = append(fieldline, f.Type)
							fieldline = append(fieldline, "#"+f.Description)
							qplist.Set(iKey, strings.Join(fieldline, " "))
						}

					})
					if qplist.Len() > 0 {
						mRpc.Qp = qplist
					}

				}

				serviceList = append(serviceList, mRpc)
			})

			muService := &microservices.MicroService{
				Name:        shadowNode.edgeServiceNode.ServiceSpec.Name,
				Methods:     serviceList,
				Target:      shadowNode.edgeServiceNode.ServiceSpec.XProto.Targetfile,
				Package:     shadowNode.edgeServiceNode.ServiceSpec.XProto.Package,
				Description: shadowNode.edgeServiceNode.ServiceSpec.Description,
			}

			// add type to "file"
			muServicesPerFile[shadowNode.edgeMicroServiceNode.SourceFile] = append(muServicesPerFile[shadowNode.edgeMicroServiceNode.SourceFile], muService)
		} else {
			// type is not in spec (shadowNode.edgeServiceNode does not exist)
			// todo: implement delete

		}
	}

	// store every item in muTypesPerFile (key is filename)
	for filename, muService := range muServicesPerFile {
		// save the stuff
		file, _ := yaml.Marshal(muService)
		if !util.DirExists(path.Dir(filename)) {
			util.MkdirRelative(path.Dir(filename))
		}
		err := ioutil.WriteFile("./"+filename, file, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func findRequestType(rpcName string, node *UTshadowNode) *typeAst.TypeAst {
	// node is in node.edgeRequestTypeNode array
	for _, n := range node.edgeRequestTypeNode {
		if n.TypeSpec.Name == rpcName+"FuroGrpcRqst" {
			return n
		}
	}
	return nil
}
