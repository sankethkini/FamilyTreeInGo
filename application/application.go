package application

import (
	"github.com/pkg/errors"
	"github.com/sankethkini/FamilyTreeInGo/model/graph"
	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

// aliasing.
type data = map[string]interface{}

type MyApp struct {
	mygraph graph.IGraph
}

func NewApp() *MyApp {
	app := MyApp{mygraph: graph.NewGraph()}
	return &app
}

func createMsg(body interface{}) []data {
	var retmsg []data
	mp := make(map[string]interface{})
	mp["message"] = body
	retmsg = append(retmsg, mp)
	return retmsg
}

// parse nodes to represnt in form of message.
func ParseNodes(nd ...*node.Node) []data {
	retmsg := make([]data, 0, len(nd))
	for _, val := range nd {
		mp := make(map[string]interface{})
		mp["id"] = val.GetID()
		mp["Name"] = val.GetName()
		retmsg = append(retmsg, mp)
	}
	return retmsg
}

// AddNode functionality adds node to the graph.
func (app *MyApp) AddNode(name, id string) ([]data, error) {
	_, ok := app.mygraph.GetNode(id)

	if ok {
		return nil, errors.Wrap(ErrNodeExists, "cannot add node")
	}
	app.mygraph.AddNode(id, name)

	msg := createMsg("node added successfully")
	return msg, nil
}

// Parents function returns the parents of a node.
func (app *MyApp) Parents(id string) ([]data, error) {
	curnode, ok := app.mygraph.GetNode(id)

	if !ok {
		return nil, errors.Wrap(ErrNodeNotFound, "cannot find parents")
	}
	par := curnode.GetParents()
	msg := ParseNodes(par...)
	return msg, nil
}

// Children function returns the children of a node.
func (app *MyApp) Children(id string) ([]data, error) {
	curnode, ok := app.mygraph.GetNode(id)

	if !ok {
		return nil, errors.Wrap(ErrNodeNotFound, "cannot find children")
	}
	chd := curnode.GetChildren()
	msg := ParseNodes(chd...)
	return msg, nil
}

// Ancestors function returns the ancestors of a node.
func (app *MyApp) Ancestors(id string) ([]data, error) {
	curnode, ok := app.mygraph.GetNode(id)

	if !ok {
		return nil, errors.Wrap(ErrNodeNotFound, "cannot find ancestors")
	}

	par := curnode.GetParents()
	visited := make(map[string]bool)
	var res []*node.Node
	for _, val := range par {
		getAncestors(val, visited, &res)
	}

	msg := ParseNodes(res...)
	return msg, nil
}

// helper dfs to find all ancestors.
func getAncestors(cur *node.Node, visited map[string]bool, res *[]*node.Node) {
	if visited[cur.GetID()] {
		return
	}
	visited[cur.GetID()] = true
	*res = append(*res, cur)
	par := cur.GetParents()
	for _, val := range par {
		if !visited[val.GetID()] {
			getAncestors(val, visited, res)
		}
	}
}

// Descendants function returns descendants of a node.
func (app *MyApp) Descendants(id string) ([]data, error) {
	curnode, ok := app.mygraph.GetNode(id)
	if !ok {
		return nil, errors.Wrap(ErrNodeNotFound, "cannot find descendants")
	}

	chd := curnode.GetChildren()
	visited := make(map[string]bool)
	var res []*node.Node
	for _, val := range chd {
		getDescendants(val, visited, &res)
	}

	msg := ParseNodes(res...)
	return msg, nil
}

// helper dfs to find descendants.
func getDescendants(cur *node.Node, visited map[string]bool, res *[]*node.Node) {
	if visited[cur.GetID()] {
		return
	}
	visited[cur.GetID()] = true
	*res = append(*res, cur)
	chd := cur.GetChildren()
	for _, val := range chd {
		if !visited[val.GetID()] {
			getDescendants(val, visited, res)
		}
	}
}

// DeleteNode function deletes the node.
func (app *MyApp) DeleteNode(id string) ([]data, error) {
	for _, val := range app.mygraph.AllNodes() {
		val.RemoveChild(id)
		val.RemoveParent(id)
	}
	err := app.mygraph.RemoveNode(id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot delete node ")
	}

	msg := createMsg("node deleted successfully")
	return msg, nil
}

// DeleteDependency function deletes dependency between two nodes.
func (app *MyApp) DeleteDependency(parentID string, childID string) ([]data, error) {
	err := app.mygraph.RemoveDependency(parentID, childID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot remove dependency")
	}

	msg := createMsg("dependency deleted successfully")
	return msg, nil
}

// to check if cycle exists by checking if childnode is in the ancestors of parent node.
func (app *MyApp) checkCycle(parentID, childID string) bool {
	pnode, _ := app.mygraph.GetNode(parentID)
	visited := make(map[string]bool)
	var res []*node.Node
	getAncestors(pnode, visited, &res)

	for _, val := range res {
		if val.GetID() == childID {
			return true
		}
	}
	return false
}

// AddDependency function add dependency between two nodes.
func (app *MyApp) AddDependency(parentID, childID string) ([]data, error) {
	err := app.mygraph.AddDependency(parentID, childID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot add dependency")
	}

	if app.checkCycle(parentID, childID) {
		return nil, errors.Wrap(ErrCyclicDependency, "cannot add dependency")
	}

	err = app.mygraph.AddDependency(parentID, childID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot add dependency")
	}

	msg := createMsg("dependency added successfully")
	return msg, nil
}
