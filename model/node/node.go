package node

type INode interface {
	AddChild(c INode) bool
	AddParent(c INode) bool
	RemoveChild(c string) bool
	RemoveParent(c string) bool
	GetParents() []INode
	GetChildren() []INode
	GetId() string
	GetName() string
}

type node struct {
	Id       string
	Name     string
	Children map[string]INode
	Parents  map[string]INode
}

func (n *node) AddChild(c INode) bool {
	if _, ok := n.Children[c.GetId()]; ok {
		return false
	}
	n.Children[c.GetId()] = c
	return true
}

func (n *node) AddParent(c INode) bool {
	if _, ok := n.Parents[c.GetId()]; ok {
		return false
	}
	n.Parents[c.GetId()] = c
	return true
}

func (n *node) RemoveChild(c string) bool {
	if _, ok := n.Children[c]; !ok {
		return false
	}
	delete(n.Children, c)
	return true
}

func (n *node) RemoveParent(c string) bool {
	if _, ok := n.Parents[c]; !ok {
		return false
	}
	delete(n.Parents, c)
	return true
}

func (n *node) GetParents() []INode {
	var p []INode
	for _, val := range n.Parents {
		p = append(p, val)
	}
	return p
}

func (n *node) GetChildren() []INode {
	var p []INode
	for _, val := range n.Children {
		p = append(p, val)
	}
	return p
}

func (n *node) GetId() string {
	return n.Id
}
func (n *node) GetName() string {
	return n.Name
}

func NewNode(id, name string) *node {
	node := node{}
	node.Id = id
	node.Name = name
	node.Children = make(map[string]INode)
	node.Parents = make(map[string]INode)
	return &node
}
