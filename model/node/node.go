package node

type Node struct {
	Id       string
	Name     string
	Children map[string]*Node
	Parents  map[string]*Node
}

func (n *Node) AddChild(c *Node) bool {
	if _, ok := n.Children[c.Id]; ok {
		return false
	}
	n.Children[c.Id] = c
	return true
}

func (n *Node) AddParent(c *Node) bool {
	if _, ok := n.Parents[c.Id]; ok {
		return false
	}
	n.Parents[c.Id] = c
	return true
}

func (n *Node) RemoveChild(c *Node) bool {
	if _, ok := n.Children[c.Id]; !ok {
		return false
	}
	delete(n.Children, c.Id)
	return true
}

func (n *Node) RemoveParent(c *Node) bool {
	if _, ok := n.Parents[c.Id]; !ok {
		return false
	}
	delete(n.Parents, c.Id)
	return true
}

func NewNode(id, name string) *Node {
	node := Node{}
	node.Id = id
	node.Name = name
	node.Children = make(map[string]*Node)
	node.Parents = make(map[string]*Node)
	return &node
}

func (n *Node) GetParents() []*Node {
	var p []*Node
	for _, val := range n.Parents {
		p = append(p, val)
	}
	return p
}

func (n *Node) GetChildren() []*Node {
	var p []*Node
	for _, val := range n.Children {
		p = append(p, val)
	}
	return p
}
