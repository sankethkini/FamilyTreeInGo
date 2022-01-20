package graph

import (
	"reflect"
	"testing"

	"github.com/sankethkini/FamilyTreeInGo/model/node"
)

type format struct {
	testname string
	id       string
	name     string
	res      node.INode
	want     bool
}

func TestAddGraph(t *testing.T) {

	var tests = []format{
		{
			testname: "adding node to graph",
			id:       "1",
			name:     "one",
			res:      node.NewNode("1", "one"),
			want:     true,
		},
	}
	mygraph := NewGraph()
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			got := mygraph.AddNode(tt.id, tt.name)
			if !reflect.DeepEqual(tt.res, got) {
				t.Errorf("wrong node added exp:%v got %v", got, tt.res)
			}

			_, val := mygraph.GetNode(tt.id)
			if val != tt.want {
				t.Errorf("node not added exp:%v got %v", val, tt.want)
			}
		})
	}
}

func TestRemoveGraph(t *testing.T) {

	var tests = []struct {
		testname string
		id       string
		name     string
		res      node.INode
		want     error
	}{
		{
			testname: "remove node from graph",
			id:       "1",
			name:     "one",
			want:     nil,
		},
		{
			testname: "remove node from graph which not exits",
			id:       "1",
			name:     "one",
			want:     NodeNotFoundErr,
		},
	}
	mygraph := NewGraph()
	mygraph.AddNode("1", "1")
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			got := mygraph.RemoveNode(tt.id)

			if got != tt.want {
				t.Errorf("node not removed exp:%v got %v", got, tt.want)
			}
		})
	}
}

func TestAllNodes(t *testing.T) {

	mygraph := NewGraph()
	node1 := mygraph.AddNode("1", "1")
	node2 := mygraph.AddNode("2", "2")

	var tests = []struct {
		testname string
		nodes    []node.INode
	}{
		{
			testname: "add nodes",
			nodes:    []node.INode{node1, node2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			got := mygraph.AllNodes()

			for i, _ := range got {
				if got[i] != tt.nodes[i] {
					t.Errorf("incorrect exp:%v got:%v", tt.nodes, got)
				}
			}
		})
	}
}

func TestAddDep(t *testing.T) {

	mygraph := NewGraph()
	node1 := mygraph.AddNode("1", "1")
	node2 := mygraph.AddNode("2", "2")

	var tests = []struct {
		testname string
		parNodes []node.INode
		chdNodes []node.INode
		parId    string
		chdId    string
		wantErr  error
	}{
		{
			testname: "add dependency",
			parId:    "1",
			chdId:    "2",
			parNodes: []node.INode{node1},
			chdNodes: []node.INode{node2},
			wantErr:  nil,
		},
		{
			testname: "add dependency parent node not exits",
			parId:    "3",
			chdId:    "2",
			parNodes: []node.INode{},
			chdNodes: []node.INode{},
			wantErr:  NodeNotFoundErr,
		},
		{
			testname: "add dependency",
			parId:    "1",
			chdId:    "3",
			parNodes: []node.INode{},
			chdNodes: []node.INode{},
			wantErr:  NodeNotFoundErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {

			err := mygraph.AddDependency(tt.parId, tt.chdId)
			if err != nil && tt.wantErr == nil {
				t.Errorf("exp %v got %v as error", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("exp %v got %v as error", tt.wantErr, err)
				} else {
					if tt.wantErr.Error() != err.Error() {
						t.Errorf("exp %v got %v as error", tt.wantErr, err)
					}
				}

			} else {
				gotPar := node2.GetParents()
				gotChd := node1.GetChildren()
				for i, _ := range gotPar {
					if gotPar[i] != tt.parNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.parNodes, gotPar)
					}
				}
				for i, _ := range gotChd {
					if gotChd[i] != tt.chdNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.chdNodes, gotChd)
					}
				}
			}

		})
	}
}

func TestRemoveDep(t *testing.T) {

	mygraph := NewGraph()
	node1 := mygraph.AddNode("1", "1")
	node2 := mygraph.AddNode("2", "2")

	mygraph.AddDependency(node1.GetId(), node2.GetId())

	var tests = []struct {
		testname string
		parId    string
		chdId    string
		parNodes []node.INode
		chdNodes []node.INode
		wantErr  error
	}{
		{
			testname: "remove dependency",
			parId:    "1",
			chdId:    "2",
			parNodes: []node.INode{},
			chdNodes: []node.INode{},
			wantErr:  nil,
		},
		{
			testname: "remove dependency",
			parId:    "3",
			chdId:    "2",
			parNodes: []node.INode{},
			chdNodes: []node.INode{},
			wantErr:  NodeNotFoundErr,
		},
		{
			testname: "remove dependency",
			parId:    "1",
			chdId:    "3",
			parNodes: []node.INode{},
			chdNodes: []node.INode{},
			wantErr:  NodeNotFoundErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := mygraph.RemoveDependency(tt.parId, tt.chdId)

			if err != nil && tt.wantErr == nil {
				t.Errorf("exp %v got %v as error", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("exp %v got %v as error", tt.wantErr, err)
				} else {
					if tt.wantErr.Error() != err.Error() {
						t.Errorf("exp %v got %v as error", tt.wantErr, err)
					}
				}

			} else {
				gotPar := node2.GetParents()
				gotChd := node1.GetChildren()
				for i, _ := range gotPar {
					if gotPar[i] != tt.parNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.parNodes, gotPar)
					}
				}
				for i, _ := range gotChd {
					if gotChd[i] != tt.chdNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.chdNodes, gotChd)
					}
				}
			}

		})
	}
}
