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
	res      *node.Node
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
		res      *node.Node
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
		nodes    []*node.Node
	}{
		{
			testname: "add nodes",
			nodes:    []*node.Node{node1, node2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			got := mygraph.AllNodes()

			for i := range got {
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
		parNodes []*node.Node
		chdNodes []*node.Node
		parId    string
		chdId    string
		wantErr  error
	}{
		{
			testname: "add dependency",
			parId:    "1",
			chdId:    "2",
			parNodes: []*node.Node{node1},
			chdNodes: []*node.Node{node2},
			wantErr:  nil,
		},
		{
			testname: "add dependency parent node not exits",
			parId:    "3",
			chdId:    "2",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  NodeNotFoundErr,
		},
		{
			testname: "add dependency",
			parId:    "1",
			chdId:    "3",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
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
				for i := range gotPar {
					if gotPar[i] != tt.parNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.parNodes, gotPar)
					}
				}
				for i := range gotChd {
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

	err := mygraph.AddDependency(node1.GetId(), node2.GetId())
	if err != nil {
		t.Errorf("cannot add dependency %v", err)
	}

	var tests = []struct {
		testname string
		parId    string
		chdId    string
		parNodes []*node.Node
		chdNodes []*node.Node
		wantErr  error
	}{
		{
			testname: "remove dependency",
			parId:    "1",
			chdId:    "2",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  nil,
		},
		{
			testname: "remove dependency",
			parId:    "3",
			chdId:    "2",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  NodeNotFoundErr,
		},
		{
			testname: "remove dependency",
			parId:    "1",
			chdId:    "3",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
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
				for i := range gotPar {
					if gotPar[i] != tt.parNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.parNodes, gotPar)
					}
				}
				for i := range gotChd {
					if gotChd[i] != tt.chdNodes[i] {
						t.Errorf("incorrect exp:%v got:%v", tt.chdNodes, gotChd)
					}
				}
			}

		})
	}
}
