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

	var tests = []format{
		{
			testname: "remove node from graph",
			id:       "1",
			name:     "one",
			want:     true,
		},
		{
			testname: "remove node from graph which not exits",
			id:       "1",
			name:     "one",
			want:     false,
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
