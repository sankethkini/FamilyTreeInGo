package node

type INode interface {
	AddChild(c *Node) bool
	AddParent(c *Node) bool
	RemoveChild(c string) bool
	RemoveParent(c string) bool
	GetParents() []*Node
	GetChildren() []*Node
	GetId() string
	GetName() string
}

type Node struct {
	id       string
	name     string
	children map[string]*Node
	parents  map[string]*Node
}

func (n *Node) AddChild(c *Node) bool {
	if _, ok := n.children[c.GetID()]; ok {
		return false
	}
	n.children[c.GetID()] = c
	return true
}

func (n *Node) AddParent(c *Node) bool {
	if _, ok := n.parents[c.GetID()]; ok {
		return false
	}
	n.parents[c.GetID()] = c
	return true
}

func (n *Node) RemoveChild(c string) bool {
	if _, ok := n.children[c]; !ok {
		return false
	}
	delete(n.children, c)
	return true
}

func (n *Node) RemoveParent(c string) bool {
	if _, ok := n.parents[c]; !ok {
		return false
	}
	delete(n.parents, c)
	return true
}

func (n *Node) GetParents() []*Node {
	p := make([]*Node, 0, len(n.parents))
	for _, val := range n.parents {
		p = append(p, val)
	}
	return p
}

func (n *Node) GetChildren() []*Node {
	p := make([]*Node, 0, len(n.children))
	for _, val := range n.children {
		p = append(p, val)
	}
	return p
}

func (n *Node) GetID() string {
	return n.id
}

func (n *Node) GetName() string {
	return n.name
}

func NewNode(id, name string) *Node {
	node := Node{}
	node.id = id
	node.name = name
	node.children = make(map[string]*Node)
	node.parents = make(map[string]*Node)
	return &node
}
