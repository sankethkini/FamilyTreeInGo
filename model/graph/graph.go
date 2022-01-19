package graph

import (
	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

type Graph struct {
	Nodes map[string]*node.Node
}

func (g *Graph) AddNode(id, name string) *node.Node {
	nd := node.NewNode(id, name)
	g.Nodes[id] = nd
	return g.Nodes[id]
}

func (g *Graph) GetNode(id string) (*node.Node, bool) {
	val, ok := g.Nodes[id]
	return val, ok
}

func (g *Graph) RemoveNode(id string) bool {
	_, ok := g.Nodes[id]
	delete(g.Nodes, id)
	return ok
}
func NewGraph() *Graph {
	cur := Graph{}
	cur.Nodes = make(map[string]*node.Node)
	return &cur
}
