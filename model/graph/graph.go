package graph

import (
	"errors"

	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

var ErrNodeNotFound = errors.New("node not found")

type IGraph interface {
	AddNode(id, name string) *node.Node
	RemoveNode(id string) error
	AllNodes() []*node.Node
	GetNode(id string) (*node.Node, bool)
	AddDependency(parentID, childID string) error
	RemoveDependency(parentID, childID string) error
}

type graph struct {
	nodes map[string]*node.Node
}

func (g *graph) AllNodes() []*node.Node {
	res := make([]*node.Node, 0, len(g.nodes))
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
		return ErrNodeNotFound
	}
	delete(g.nodes, id)
	return nil
}

func (g *graph) AddDependency(parentID, childID string) error {
	_, ok := g.nodes[parentID]
	if !ok {
		return ErrNodeNotFound
	}
	_, ok = g.nodes[childID]
	if !ok {
		return ErrNodeNotFound
	}
	parentNode := g.nodes[parentID]
	childNode := g.nodes[childID]
	parentNode.AddChild(childNode)
	childNode.AddParent(parentNode)
	return nil
}

func (g *graph) RemoveDependency(parentID, childID string) error {
	_, ok := g.nodes[parentID]
	if !ok {
		return ErrNodeNotFound
	}
	_, ok = g.nodes[childID]
	if !ok {
		return ErrNodeNotFound
	}

	parentNode := g.nodes[parentID]
	childNode := g.nodes[childID]
	parentNode.RemoveChild(childID)
	childNode.RemoveParent(parentID)
	return nil
}

func NewGraph() IGraph {
	cur := graph{}
	cur.nodes = make(map[string]*node.Node)
	return &cur
}
