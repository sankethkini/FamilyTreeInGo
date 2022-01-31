package graph

import (
	"fmt"
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
	tests := []format{
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
	tests := []struct {
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
			want:     ErrNodeNotFound,
		},
	}
	mygraph := NewGraph()
	mygraph.AddNode("1", "1")
	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			got := mygraph.RemoveNode(tt.id)
			if !checkErrorsEquality(tt.want, got) {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}

func TestAllNodes(t *testing.T) {
	mygraph := NewGraph()
	node1 := mygraph.AddNode("1", "1")
	node2 := mygraph.AddNode("2", "2")

	tests := []struct {
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

			if len(got) != len(tt.nodes) {
				t.Errorf("expected %v got %v", tt.nodes, got)
			}
		})
	}
}

func TestAddDep(t *testing.T) {
	mygraph := NewGraph()
	node1 := mygraph.AddNode("1", "1")
	node2 := mygraph.AddNode("2", "2")

	tests := []struct {
		testname string
		parNodes []*node.Node
		chdNodes []*node.Node
		parID    string
		chdID    string
		wantErr  error
	}{
		{
			testname: "add dependency",
			parID:    "1",
			chdID:    "2",
			parNodes: []*node.Node{node1},
			chdNodes: []*node.Node{node2},
			wantErr:  nil,
		},
		{
			testname: "add dependency parent node not exits",
			parID:    "3",
			chdID:    "2",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  ErrNodeNotFound,
		},
		{
			testname: "add dependency",
			parID:    "1",
			chdID:    "3",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  ErrNodeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := mygraph.AddDependency(tt.parID, tt.chdID)
			eq := checkErrorsEquality(tt.wantErr, err)
			if !eq {
				t.Errorf("expected error %v got %v", tt.wantErr, err)
			}

			if tt.wantErr == nil {
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

func checkErrorsEquality(err1 error, err2 error) bool {
	if err2 != nil && err1 == nil {
		fmt.Println(2)
		return false
	}
	if err1 != nil {
		if err2 == nil {
			fmt.Println(2)
			return false
		} else if err1.Error() != err2.Error() {
			return false
		}
	}
	return true
}

func TestRemoveDep(t *testing.T) {
	mygraph := NewGraph()
	node1 := mygraph.AddNode("1", "1")
	node2 := mygraph.AddNode("2", "2")

	err := mygraph.AddDependency(node1.GetID(), node2.GetID())
	if err != nil {
		t.Errorf("cannot add dependency %v", err)
	}

	tests := []struct {
		testname string
		parID    string
		chdID    string
		parNodes []*node.Node
		chdNodes []*node.Node
		wantErr  error
	}{
		{
			testname: "remove dependency",
			parID:    "1",
			chdID:    "2",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  nil,
		},
		{
			testname: "remove dependency parent node not exists",
			parID:    "3",
			chdID:    "2",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  ErrNodeNotFound,
		},
		{
			testname: "remove dependency child node not exits",
			parID:    "1",
			chdID:    "3",
			parNodes: []*node.Node{},
			chdNodes: []*node.Node{},
			wantErr:  ErrNodeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			err := mygraph.RemoveDependency(tt.parID, tt.chdID)
			eq := checkErrorsEquality(tt.wantErr, err)
			if !eq {
				t.Errorf("expected error %v got %v", tt.wantErr, err)
			}

			if tt.wantErr == nil {
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
