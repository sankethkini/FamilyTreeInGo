package node

import "testing"

type nodeargs struct {
	n1 *Node
	n2 *Node
}

type format struct {
	name  string
	args  nodeargs
	want  bool
	nodes []*Node
}

var node1 *Node = NewNode("1", "one")
var node2 *Node = NewNode("2", "two")

func TestAddchild(t *testing.T) {
	var tests = []format{
		{
			name:  "adding child",
			args:  nodeargs{n1: node1, n2: node2},
			want:  true,
			nodes: []*Node{node2},
		},
		{
			name:  "adding child",
			args:  nodeargs{n1: node1, n2: node2},
			want:  false,
			nodes: []*Node{node2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.n1.AddChild(tt.args.n2)
			if got != tt.want {
				t.Errorf("expected %v got %v", tt.want, got)
			}
			ngot := tt.args.n1.GetChildren()
			for i, val := range ngot {
				if val != tt.nodes[i] {
					t.Errorf("expected %v got %v", tt.want, ngot)
				}
			}
		})
	}
}

func TestAddparent(t *testing.T) {
	var tests = []format{
		{
			name:  "adding parent",
			args:  nodeargs{n1: node1, n2: node2},
			want:  true,
			nodes: []*Node{node2},
		},
		{
			name:  "adding parent",
			args:  nodeargs{n1: node1, n2: node2},
			want:  false,
			nodes: []*Node{node2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.n1.AddParent(tt.args.n2)
			if got != tt.want {
				t.Errorf("expected %v got %v", tt.want, got)
			}
			ngot := tt.args.n1.GetParents()
			for i, val := range ngot {
				if val != tt.nodes[i] {
					t.Errorf("expected %v got %v", tt.want, ngot)
				}
			}
		})
	}
}
