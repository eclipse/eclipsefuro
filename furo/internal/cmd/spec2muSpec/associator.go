package spec2muSpec

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/enumAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/microenums"
	"github.com/eclipse/eclipsefuro/furo/pkg/microservices"
	"github.com/eclipse/eclipsefuro/furo/pkg/microtypes"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/spf13/viper"
	"strings"
)

type UTshadowNode struct {
	IdentifierP          string // package part used to idenitify all edges
	IdentifierT          string // typename part used to idenitify all edges
	edgeRequestTypeNode  []*typeAst.TypeAst
	edgeTypeNode         *typeAst.TypeAst
	edgeEnumNode         *enumAst.EnumAst
	edgeServiceNode      *serviceAst.ServiceAst
	edgeMicroTypeNode    *microtypes.MicroTypeAst
	edgeMicroEnumNode    *microenums.MicroEnumAst
	edgeMicroServiceNode *microservices.MicroServiceAst
}

type UTShadowList struct {
	Items                 []*UTshadowNode
	TypeItemsByName       map[string]*UTshadowNode
	EnumItemsByName       map[string]*UTshadowNode
	ServiceItemsByName    map[string]*UTshadowNode
	ServiceRequestsByName map[string]*UTshadowNode
}

func NewUTShadowList() UTShadowList {
	return UTShadowList{
		Items:                 []*UTshadowNode{},
		TypeItemsByName:       map[string]*UTshadowNode{},
		EnumItemsByName:       map[string]*UTshadowNode{},
		ServiceItemsByName:    map[string]*UTshadowNode{},
		ServiceRequestsByName: map[string]*UTshadowNode{},
	}
}

func (s *UTShadowList) AddTypeNode(fullTypeName string, ast *typeAst.TypeAst) *UTshadowNode {

	if viper.GetString("muSpec.requestTypeSuffix") != "" && strings.HasSuffix(fullTypeName, viper.GetString("muSpec.requestTypeSuffix")) {
		return s.AddRequestTypeNode(ast)
	}

	return s.AddRegularTypeNode(fullTypeName, ast)
}

func (s *UTShadowList) AddEnumNode(fullEnumName string, ast *enumAst.EnumAst) *UTshadowNode {

	return s.AddRegularEnumNode(fullEnumName, ast)
}

func (s *UTShadowList) AddRegularTypeNode(fullTypeName string, ast *typeAst.TypeAst) *UTshadowNode {
	var node *UTshadowNode
	// find item by name, nok => create
	if s.TypeItemsByName[fullTypeName] == nil {
		node = s.CreateShadowTypeItem(fullTypeName)
	} else {
		node = s.TypeItemsByName[fullTypeName]
	}
	node.edgeTypeNode = ast
	return node
}

func (s *UTShadowList) AddRegularEnumNode(fullEnumName string, ast *enumAst.EnumAst) *UTshadowNode {
	var node *UTshadowNode
	// find item by name, nok => create
	if s.EnumItemsByName[fullEnumName] == nil {
		node = s.CreateShadowEnumItem(fullEnumName)
	} else {
		node = s.EnumItemsByName[fullEnumName]
	}
	node.edgeEnumNode = ast
	return node
}

func (s *UTShadowList) AddRequestTypeNode(ast *typeAst.TypeAst) *UTshadowNode {
	fullTypeName := ast.TypeSpec.XProto.Package + "." + ast.TypeSpec.Name[0:len(ast.TypeSpec.Name)-len(viper.GetString("muSpec.requestTypeSuffix"))]

	var node *UTshadowNode
	// find item by name, nok => create

	node = s.ServiceRequestsByName[fullTypeName]

	if node == nil {
		node = &UTshadowNode{

			edgeRequestTypeNode: []*typeAst.TypeAst{},
		}
	}
	node.edgeRequestTypeNode = append(node.edgeRequestTypeNode, ast)
	return node
}
func (s *UTShadowList) AddMicroTypeNode(ast *microtypes.MicroTypeAst) *UTshadowNode {
	// check for names without package

	fullTypeName := ast.Package + "." + ast.Type

	var node *UTshadowNode
	// find item by name, nok => create
	if s.TypeItemsByName[fullTypeName] == nil {
		node = s.CreateShadowTypeItem(fullTypeName)
	} else {
		node = s.TypeItemsByName[fullTypeName]
	}
	node.edgeMicroTypeNode = ast
	return node
}
func (s *UTShadowList) AddMicroEnumNode(ast *microenums.MicroEnumAst) *UTshadowNode {
	// check for names without package

	fullEnumName := ast.Package + "." + ast.Type

	var node *UTshadowNode
	// find item by name, nok => create
	if s.EnumItemsByName[fullEnumName] == nil {
		node = s.CreateShadowEnumItem(fullEnumName)
	} else {
		node = s.EnumItemsByName[fullEnumName]
	}
	node.edgeMicroEnumNode = ast
	return node
}

func (s *UTShadowList) AddMicroServiceNode(ast *microservices.MicroServiceAst) *UTshadowNode {
	var node *UTshadowNode
	fullTypeName := ast.Package + "." + ast.Name

	// find item by name, nok => create
	if s.ServiceItemsByName[fullTypeName] == nil {
		node = s.CreateShadowServiceItem(fullTypeName)
	} else {
		node = s.ServiceItemsByName[fullTypeName]
	}
	node.edgeMicroServiceNode = ast
	return node
}
func (s *UTShadowList) AddServiceNode(ast *serviceAst.ServiceAst) *UTshadowNode {
	var node *UTshadowNode
	fullTypeName := ast.ServiceSpec.XProto.Package + "." + ast.ServiceSpec.Name
	// find item by name, nok => create
	if s.ServiceItemsByName[fullTypeName] == nil {
		node = s.CreateShadowServiceItem(fullTypeName)
	} else {
		node = s.ServiceItemsByName[fullTypeName]
	}

	node.edgeServiceNode = ast

	ast.ServiceSpec.Services.Map(func(iKey interface{}, iValue interface{}) {
		rpc := iValue.(*specSpec.Rpc)
		s.ServiceRequestsByName[ast.ServiceSpec.XProto.Package+"."+rpc.RpcName] = node
	})

	return node
}

// get all microtypes without any connection to something
func (s *UTShadowList) GetUnconnectedMicroTypes() []*microtypes.MicroTypeAst {
	l := []*microtypes.MicroTypeAst{}
	for _, item := range s.Items {
		if item.edgeMicroTypeNode != nil &&
			item.edgeTypeNode == nil &&
			item.edgeServiceNode == nil &&
			item.edgeRequestTypeNode == nil {
			l = append(l, item.edgeMicroTypeNode)
		}
	}
	return l
}

func (s *UTShadowList) CreateShadowTypeItem(name string) *UTshadowNode {
	n := &UTshadowNode{}
	s.TypeItemsByName[name] = n
	s.Items = append(s.Items, n)
	ta := strings.Split(name, ".")

	n.IdentifierP = strings.Join(ta[0:len(ta)-1], ".")
	n.IdentifierT = strings.Join(ta[len(ta)-1:], ".")
	return n
}

func (s *UTShadowList) CreateShadowEnumItem(name string) *UTshadowNode {
	n := &UTshadowNode{}
	s.EnumItemsByName[name] = n
	s.Items = append(s.Items, n)
	ta := strings.Split(name, ".")

	n.IdentifierP = strings.Join(ta[0:len(ta)-1], ".")
	n.IdentifierT = strings.Join(ta[len(ta)-1:], ".")
	return n
}

func (s *UTShadowList) CreateShadowServiceItem(name string) *UTshadowNode {
	n := &UTshadowNode{}
	s.ServiceItemsByName[name] = n
	s.Items = append(s.Items, n)
	ta := strings.Split(name, ".")

	n.IdentifierP = strings.Join(ta[0:len(ta)-1], ".")
	n.IdentifierT = strings.Join(ta[len(ta)-1:], ".")
	return n
}
