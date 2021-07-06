package spec2muSpec

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/microservices"
	"github.com/eclipse/eclipsefuro/furo/pkg/microtypes"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"strings"
)

type UTshadowNode struct {
	IdentifierP            string // package part used to idenitify all edges
	IdentifierT            string // typename part used to idenitify all edges
	edgeEntityTypeNode     *typeAst.TypeAst
	edgeCollectionTypeNode *typeAst.TypeAst
	edgeRequestTypeNode    []*typeAst.TypeAst
	edgeTypeNode           *typeAst.TypeAst
	edgeServiceNode        *serviceAst.ServiceAst
	edgeMicroTypeNode      *microtypes.MicroTypeAst
	edgeMicroServiceNode   *microservices.MicroServiceAst
}

type UTShadowList struct {
	Items                 []*UTshadowNode
	TypeItemsByName       map[string]*UTshadowNode
	ServiceItemsByName    map[string]*UTshadowNode
	ServiceRequestsByName map[string]*UTshadowNode
}

func NewUTShadowList() UTShadowList {
	return UTShadowList{
		Items:                 []*UTshadowNode{},
		TypeItemsByName:       map[string]*UTshadowNode{},
		ServiceItemsByName:    map[string]*UTshadowNode{},
		ServiceRequestsByName: map[string]*UTshadowNode{},
	}
}

func (s *UTShadowList) AddTypeNode(fullTypeName string, ast *typeAst.TypeAst) *UTshadowNode {
	// Triage -> regular | entity | collection | request
	if strings.HasSuffix(fullTypeName, "Entity") {
		return s.AddEntityTypeNode(fullTypeName[0:len(fullTypeName)-6], ast)
	}

	if strings.HasSuffix(fullTypeName, "Collection") {
		return s.AddCollectionTypeNode(fullTypeName[0:len(fullTypeName)-10], ast)
	}

	if strings.HasSuffix(fullTypeName, "Request") {
		return s.AddRequestTypeNode(ast)
	}

	return s.AddRegularTypeNode(fullTypeName, ast)
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
func (s *UTShadowList) AddEntityTypeNode(fullTypeName string, ast *typeAst.TypeAst) *UTshadowNode {
	var node *UTshadowNode
	// find item by name, nok => create
	if s.TypeItemsByName[fullTypeName] == nil {
		node = s.CreateShadowTypeItem(fullTypeName)
	} else {
		node = s.TypeItemsByName[fullTypeName]
	}
	node.edgeEntityTypeNode = ast
	return node
}
func (s *UTShadowList) AddCollectionTypeNode(fullTypeName string, ast *typeAst.TypeAst) *UTshadowNode {
	var node *UTshadowNode
	// find item by name, nok => create
	if s.TypeItemsByName[fullTypeName] == nil {
		node = s.CreateShadowTypeItem(fullTypeName)
	} else {
		node = s.TypeItemsByName[fullTypeName]
	}
	node.edgeCollectionTypeNode = ast
	return node

}
func (s *UTShadowList) AddRequestTypeNode(ast *typeAst.TypeAst) *UTshadowNode {
	fullTypeName := ast.TypeSpec.XProto.Package + "." + ast.TypeSpec.Name[0:len(ast.TypeSpec.Name)-7]

	var node *UTshadowNode
	// find item by name, nok => create

	node = s.ServiceRequestsByName[fullTypeName]

	node.edgeRequestTypeNode = append(node.edgeRequestTypeNode, ast)
	return node
}
func (s *UTShadowList) AddMicroTypeNode(ast *microtypes.MicroTypeAst) *UTshadowNode {
	// check for names without package

	fullTypeName := ast.Package + "." + ast.Type

	if strings.HasSuffix(fullTypeName, "Entity") {
		return nil
	}
	if strings.HasSuffix(fullTypeName, "Collection") {
		return nil
	}

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
			item.edgeEntityTypeNode == nil &&
			item.edgeCollectionTypeNode == nil &&
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
func (s *UTShadowList) CreateShadowServiceItem(name string) *UTshadowNode {
	n := &UTshadowNode{}
	s.ServiceItemsByName[name] = n
	s.Items = append(s.Items, n)
	ta := strings.Split(name, ".")

	n.IdentifierP = strings.Join(ta[0:len(ta)-1], ".")
	n.IdentifierT = strings.Join(ta[len(ta)-1:], ".")
	return n
}
