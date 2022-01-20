package graph

import (
	"errors"

	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

var NodeNotFoundErr = errors.New("node not found")

type IGraph interface {
	AddNode(id, name string) node.INode
	RemoveNode(id string) error
	AllNodes() []node.INode
	GetNode(id string) (node.INode, bool)
	AddDependency(parentId, childId string) error
	RemoveDependency(parentId, childId string) error
}

type graph struct {
	Nodes map[string]node.INode
}

func (g *graph) AllNodes() []node.INode {
	var res []node.INode
	for _, val := range g.Nodes {
		res = append(res, val)
	}
	return res
}
func (g *graph) AddNode(id, name string) node.INode {
	nd := node.NewNode(id, name)
	g.Nodes[id] = nd
	return g.Nodes[id]
}

func (g *graph) GetNode(id string) (node.INode, bool) {
	val, ok := g.Nodes[id]
	return val, ok
}

func (g *graph) RemoveNode(id string) error {
	_, ok := g.Nodes[id]
	if !ok {
		return NodeNotFoundErr
	}
	delete(g.Nodes, id)
	return nil
}

func (g *graph) AddDependency(parentId, childId string) error {
	_, ok := g.Nodes[parentId]
	if !ok {
		return NodeNotFoundErr
	}
	_, ok = g.Nodes[childId]
	if !ok {
		return NodeNotFoundErr
	}
	parentNode := g.Nodes[parentId]
	childNode := g.Nodes[childId]
	parentNode.AddChild(childNode)
	childNode.AddParent(parentNode)
	return nil
}

func (g *graph) RemoveDependency(parentId, childId string) error {
	_, ok := g.Nodes[parentId]
	if !ok {
		return NodeNotFoundErr
	}
	_, ok = g.Nodes[childId]
	if !ok {
		return NodeNotFoundErr
	}

	parentNode := g.Nodes[parentId]
	childNode := g.Nodes[childId]
	parentNode.RemoveChild(childId)
	childNode.RemoveParent(parentId)
	return nil
}

func NewGraph() IGraph {
	cur := graph{}
	cur.Nodes = make(map[string]node.INode)
	return &cur
}
