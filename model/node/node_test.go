package node

import "testing"

type nodeargs struct {
	n1 *node
	n2 *node
}

type format struct {
	name  string
	args  nodeargs
	want  bool
	nodes []string
}

var node1 *node = NewNode("1", "one")
var node2 *node = NewNode("2", "two")

func TestAddChild(t *testing.T) {
	var tests = []format{
		{
			name:  "adding child",
			args:  nodeargs{n1: node1, n2: node2},
			want:  true,
			nodes: []string{"2"},
		},
		{
			name:  "adding child that already exists",
			args:  nodeargs{n1: node1, n2: node2},
			want:  false,
			nodes: []string{"2"},
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
				if val.GetId() != tt.nodes[i] {
					t.Errorf("expected %v got %v", tt.want, ngot)
				}
			}
		})
	}
}

func TestAddParent(t *testing.T) {
	var tests = []format{
		{
			name:  "adding parent",
			args:  nodeargs{n1: node1, n2: node2},
			want:  true,
			nodes: []string{"2"},
		},
		{
			name:  "adding parent which already exists",
			args:  nodeargs{n1: node1, n2: node2},
			want:  false,
			nodes: []string{"2"},
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
				if val.GetId() != tt.nodes[i] {
					t.Errorf("expected %v got %v", tt.want, ngot)
				}
			}
		})
	}
}

func TestRemoveChild(t *testing.T) {
	var tests = []format{
		{
			name: "remving child",
			args: nodeargs{n1: node1, n2: node2},
			want: true,
		},
		{
			name: "removing child that doesnot exists",
			args: nodeargs{n1: node1, n2: node2},
			want: false,
		},
	}

	node1.AddChild(node2)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.n1.RemoveChild(tt.args.n2.Id)
			if got != tt.want {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}

func TestRemoveParent(t *testing.T) {
	var tests = []format{
		{
			name: "remving parent",
			args: nodeargs{n1: node1, n2: node2},
			want: true,
		},
		{
			name: "removing parent that doesnot exists",
			args: nodeargs{n1: node1, n2: node2},
			want: false,
		},
	}

	node1.AddParent(node2)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.n1.RemoveParent(tt.args.n2.Id)
			if got != tt.want {
				t.Errorf("expected %v got %v", tt.want, got)
			}
		})
	}
}
