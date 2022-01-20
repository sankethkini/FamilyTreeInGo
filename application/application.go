package application

import (
	"fmt"

	"github.com/sankethkini/FamilyTreeInGo/model/graph"
	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

//global variable for graph
var mygraph graph.IGraph

//aliasing
type data = map[string]interface{}

//initializing graph
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

//parse nodes to represnt in form of message
func ParseNodes(nd ...node.INode) []data {
	var retmsg []data
	for _, val := range nd {
		mp := make(map[string]interface{})
		mp["id"] = val.GetId()
		mp["Name"] = val.GetName()
		retmsg = append(retmsg, mp)
	}
	return retmsg
}

//AddNode functionality adds node to the graph
func AddNode(name, id string) ([]data, error) {
	_, ok := mygraph.GetNode(id)

	if ok {
		return nil, fmt.Errorf("node cannot be added %w", NodeExistsErr)
	}
	mygraph.AddNode(id, name)

	msg := createMsg("message", "node added successfuly")
	return msg, nil
}

//Parents function returns the parents of a node
func Parents(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		return nil, fmt.Errorf("cannot find parents %w", NodeNotFoundErr)
	}
	par := curnode.GetParents()

	msg := ParseNodes(par...)
	return msg, nil
}

//Children function returns the children of a node
func Children(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		return nil, fmt.Errorf("cannot find children %w", NodeNotFoundErr)
	}
	chd := curnode.GetChildren()
	msg := ParseNodes(chd...)
	return msg, nil
}

//Ancestors function returns the ancestors of a node
func Ancestors(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)

	if !ok {
		return nil, fmt.Errorf("cannot find ancestors %w", NodeNotFoundErr)
	}

	par := curnode.GetParents()
	visited := make(map[string]bool)
	var res []node.INode
	for _, val := range par {
		getAncestors(val, visited, &res)
	}

	msg := ParseNodes(res...)
	return msg, nil
}

//helper dfs to find all ancestors
func getAncestors(cur node.INode, visited map[string]bool, res *[]node.INode) {
	if visited[cur.GetId()] {
		return
	}
	visited[cur.GetId()] = true
	*res = append(*res, cur)
	par := cur.GetParents()
	for _, val := range par {
		if !visited[val.GetId()] {
			getAncestors(val, visited, res)
		}
	}

}

//Descendants function returns descendants of a node
func Descendants(id string) ([]data, error) {
	curnode, ok := mygraph.GetNode(id)
	if !ok {
		return nil, fmt.Errorf("cannot find descendants %w", NodeNotFoundErr)
	}

	chd := curnode.GetChildren()
	visited := make(map[string]bool)
	var res []node.INode
	for _, val := range chd {
		getDescendants(val, visited, &res)
	}

	msg := ParseNodes(res...)
	return msg, nil
}

//helper dfs to find descendants
func getDescendants(cur node.INode, visited map[string]bool, res *[]node.INode) {
	if visited[cur.GetId()] {
		return
	}
	visited[cur.GetId()] = true
	*res = append(*res, cur)
	chd := cur.GetChildren()
	for _, val := range chd {
		if !visited[val.GetId()] {
			getDescendants(val, visited, res)
		}
	}

}

//DeleteNode function deletes the node
func DeleteNode(id string) []data {
	_, ok := mygraph.GetNode(id)

	if !ok {
		msg := createMsg("message", "node does not exists")
		return msg
	}

	for _, val := range mygraph.AllNodes() {
		val.RemoveChild(id)
		val.RemoveParent(id)
	}
	mygraph.RemoveNode(id)

	msg := createMsg("message", "node deleted successfuly")
	return msg
}

//DeleteDependency function deletes dependency between two nodes
func DeleteDependency(parentId string, childId string) ([]data, error) {

	err := mygraph.RemoveDependency(parentId, childId)
	if err != nil {
		return nil, fmt.Errorf("cannot remove dependency %w", err)
	}

	msg := createMsg("message", "dependency deleted successfuly")
	return msg, nil
}

//to check if cycle exists by checking if childnode is in the ancestors of parent node
func checkCycle(parentId, childId string) bool {
	pnode, _ := mygraph.GetNode(parentId)
	visited := make(map[string]bool)
	var res []node.INode
	getAncestors(pnode, visited, &res)

	for _, val := range res {
		if val.GetId() == childId {
			return true
		}
	}
	return false
}

//AddDependency function add dependency between two nodes
func AddDependency(parentId, childId string) ([]data, error) {

	err := mygraph.AddDependency(parentId, childId)
	if err != nil {
		return nil, fmt.Errorf("cannot add dependency %w", err)
	}

	if checkCycle(parentId, childId) {
		return nil, fmt.Errorf("cannot add dependency %w", CyclicDependencyErr)
	}

	mygraph.AddDependency(parentId, childId)

	msg := createMsg("message", "dependency added successfuly")
	return msg, nil
}
