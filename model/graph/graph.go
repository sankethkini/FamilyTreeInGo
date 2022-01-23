package graph

import (
	"errors"

	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

var NodeNotFoundErr = errors.New("node not found")

type IGraph interface {
	AddNode(id, name string) *node.Node
	RemoveNode(id string) error
	AllNodes() []*node.Node
	GetNode(id string) (*node.Node, bool)
	AddDependency(parentId, childId string) error
	RemoveDependency(parentId, childId string) error
}

type graph struct {
	nodes map[string]*node.Node
}

func (g *graph) AllNodes() []*node.Node {
	var res []*node.Node
	for _, val := range g.nodes {
		res = append(res, val)
	}
	return res
}
func (g *graph) AddNode(id, name string) *node.Node {
	nd := node.NewNode(id, name)
	g.nodes[id] = nd
	return g.nodes[id]
}

func (g *graph) GetNode(id string) (*node.Node, bool) {
	val, ok := g.nodes[id]
	return val, ok
}

func (g *graph) RemoveNode(id string) error {
	_, ok := g.nodes[id]
	if !ok {
		return NodeNotFoundErr
	}
	delete(g.nodes, id)
	return nil
}

func (g *graph) AddDependency(parentId, childId string) error {
	_, ok := g.nodes[parentId]
	if !ok {
		return NodeNotFoundErr
	}
	_, ok = g.nodes[childId]
	if !ok {
		return NodeNotFoundErr
	}
	parentNode := g.nodes[parentId]
	childNode := g.nodes[childId]
	parentNode.AddChild(childNode)
	childNode.AddParent(parentNode)
	return nil
}

func (g *graph) RemoveDependency(parentId, childId string) error {
	_, ok := g.nodes[parentId]
	if !ok {
		return NodeNotFoundErr
	}
	_, ok = g.nodes[childId]
	if !ok {
		return NodeNotFoundErr
	}

	parentNode := g.nodes[parentId]
	childNode := g.nodes[childId]
	parentNode.RemoveChild(childId)
	childNode.RemoveParent(parentId)
	return nil
}

func NewGraph() IGraph {
	cur := graph{}
	cur.nodes = make(map[string]*node.Node)
	return &cur
}
