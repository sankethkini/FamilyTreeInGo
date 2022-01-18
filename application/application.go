package application

import (
	"fmt"

	"github.com/sankethkini/FamilyTreeInGo/model/graph"
	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

//global variable for graph
var mygraph *graph.Graph

//aliasing
type data = map[string]interface{}

func init() {
	mygraph = graph.NewGraph()
}

func createMsg(msg string, body interface{}) []data {
	var retmsg []data
	mp := make(map[string]interface{})
	mp[msg] = body
	retmsg = append(retmsg, mp)
	return retmsg
}

func ParseNodes(nd ...*node.Node) []data {
	var retmsg []data
	for _, val := range nd {
		mp := make(map[string]interface{})
		mp["id"] = val.Id
		mp["Name"] = val.Name
		retmsg = append(retmsg, mp)
	}
	return retmsg
}

func AddNode(name, id string) ([]data, error) {
	_, ok := mygraph.GetNode(id)

	if ok {
		return nil, fmt.Errorf("node already exists %w", NodeExists)
	}
	mygraph.AddNode(id, name)

	msg := createMsg("message", "node added successfuly")
	return msg, nil
}

func Parents(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		return nil, fmt.Errorf("node does not exists %w", NodeNotFound)
	}
	parents := curnode.GetParents()
	msg := ParseNodes(parents...)
	return msg, nil
}

func Children(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		return nil, fmt.Errorf("node does not exists %w", NodeNotFound)
	}
	children := curnode.GetChildren()
	msg := ParseNodes(children...)
	return msg, nil
}

func Ancestors(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		return nil, fmt.Errorf("node does not exists %w", NodeNotFound)
	}

	parents := curnode.GetParents()
	visited := make(map[string]bool)
	var res []*node.Node
	for _, val := range parents {
		getAncestors(val, visited, &res)
	}

	msg := ParseNodes(res...)
	return msg, nil
}

func getAncestors(cur *node.Node, visited map[string]bool, res *[]*node.Node) {
	if visited[cur.Id] {
		return
	}
	visited[cur.Id] = true
	*res = append(*res, cur)
	parents := cur.GetParents()
	for _, val := range parents {
		if !visited[val.Id] {
			getAncestors(val, visited, res)
		}
	}

}

func Descendants(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)
	if !ok {
		return nil, fmt.Errorf("node does not exists %w", NodeNotFound)
	}

	children := curnode.GetChildren()
	visited := make(map[string]bool)
	var res []*node.Node
	for _, val := range children {
		getDescendants(val, visited, &res)
	}

	msg := ParseNodes(res...)
	return msg, nil
}

func getDescendants(cur *node.Node, visited map[string]bool, res *[]*node.Node) {
	if visited[cur.Id] {
		return
	}
	visited[cur.Id] = true
	*res = append(*res, cur)
	parents := cur.GetChildren()
	for _, val := range parents {
		if !visited[val.Id] {
			getDescendants(val, visited, res)
		}
	}

}

func DeleteNode(id string) []data {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		msg := createMsg("message", "node does not exists")
		return msg
	}

	for _, val := range mygraph.Nodes {
		val.RemoveChild(curnode)
		val.RemoveParent(curnode)
	}
	mygraph.RemoveNode(id)

	msg := createMsg("message", "node deleted successfuly")
	return msg
}

func DeleteDependency(parentid string, childid string) []data {
	parentnode, ok := mygraph.GetNode(parentid)

	if !ok {
		msg := createMsg("message", "parent node does not exists")
		return msg
	}

	childnode, ok := mygraph.GetNode(childid)

	if !ok {
		msg := createMsg("message", "child node does not exists")
		return msg
	}

	parentnode.RemoveChild(childnode)
	childnode.RemoveParent(parentnode)

	msg := createMsg("message", "dependency deleted successfuly")
	return msg
}

func checkCycle(parentid, childid string) bool {
	parentnode, _ := mygraph.GetNode(parentid)
	visited := make(map[string]bool)
	var res []*node.Node
	getAncestors(parentnode, visited, &res)

	for _, val := range res {
		if val.Id == childid {
			return true
		}
	}
	return false
}

func AddDependency(parentid, childid string) ([]data, error) {
	parentnode, ok := mygraph.GetNode(parentid)

	if !ok {
		return nil, fmt.Errorf("parent node does not exists %w", NodeNotFound)
	}

	childnode, ok := mygraph.GetNode(childid)

	if !ok {
		return nil, fmt.Errorf("child node does not exists %w", NodeNotFound)
	}

	if checkCycle(parentid, childid) {
		return nil, fmt.Errorf("cannot add dependency %w", CyclicDependency)
	}

	parentnode.AddChild(childnode)
	childnode.AddParent(parentnode)

	msg := createMsg("message", "dependency added successfuly")
	return msg, nil
}
