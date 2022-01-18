package application

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestAddNodes(t *testing.T) {
	var test = []struct {
		testname string
		name     string
		id       string
		experr   error
		wantErr  bool
	}{
		{
			testname: "adding node",
			name:     "one",
			id:       "1",
			experr:   nil,
			wantErr:  false,
		},
		{
			testname: "adding node which exists",
			name:     "one",
			id:       "1",
			experr:   fmt.Errorf("node already exists %w", NodeExists),
			wantErr:  true,
		},
	}

	for _, tt := range test {
		t.Run(tt.testname, func(t *testing.T) {
			_, got := AddNode(tt.name, tt.id)

			if tt.wantErr && got == nil {
				t.Errorf("not got any error exp:%v got %v", tt.experr, got)
				if tt.experr.Error() != got.Error() {
					t.Errorf("node not added properly exp: %v got: %v", tt.experr, got)
				}
			}
		})
	}
}

func TestAddDependency(t *testing.T) {
	_, err := AddNode("four", "4")
	if err != nil {
		t.Errorf("node cant be added")
	}
	_, err = AddNode("five", "5")
	if err != nil {
		t.Errorf("node cant be added")
	}
	_, err = AddNode("six", "6")

	if err != nil {
		t.Errorf("node cant be added")
	}

	var test = []struct {
		testname string
		pid      string
		cid      string
		experr   error
		wantErr  bool
	}{
		{
			testname: "adding normal dependency",
			pid:      "4",
			cid:      "5",
			experr:   nil,
			wantErr:  false,
		},
		{
			testname: "parent node not exits",
			pid:      "10",
			cid:      "5",
			experr:   fmt.Errorf("parent node does not exists %w", NodeNotFound),
			wantErr:  true,
		},
		{
			testname: "cyclic dependency",
			pid:      "5",
			cid:      "4",
			experr:   fmt.Errorf("cannot add dependency %w", CyclicDependency),
			wantErr:  true,
		},
	}

	for _, tt := range test {
		t.Run(tt.testname, func(t *testing.T) {
			_, got := AddDependency(tt.pid, tt.cid)

			if tt.wantErr && got == nil {
				t.Errorf("not got any error exp:%v got %v", tt.experr, got)
				if tt.experr.Error() != got.Error() {
					t.Errorf("dependency not added properly exp: %v got: %v", tt.experr, got)
				}
			}
		})
	}
}

func checkEqual(d1, d2 []data) bool {
	ids1 := make([]string, 0)
	for _, val := range d1 {
		ids1 = append(ids1, val["id"].(string))
	}
	ids2 := make([]string, 0)
	for _, val := range d2 {
		ids2 = append(ids2, val["id"].(string))
	}
	sort.Strings(ids1)
	sort.Strings(ids2)
	return reflect.DeepEqual(ids1, ids2)
}
func TestCheckAscendentsAndDescendents(t *testing.T) {
	_, err := AddNode("eight", "8")
	if err != nil {
		t.Errorf("node cant be added")
	}
	_, err = AddNode("nine", "9")
	if err != nil {
		t.Errorf("node cant be added")
	}
	_, err = AddNode("ten", "10")
	if err != nil {
		t.Errorf("node cant be added")
	}

	_, err = AddDependency("8", "9")
	if err != nil {
		t.Errorf("depend cant be added")
	}
	_, err = AddDependency("9", "10")
	if err != nil {
		t.Errorf("depend cant be added")
	}

	n, _ := mygraph.GetNode("8")
	mp8 := ParseNodes(n)[0]

	n, _ = mygraph.GetNode("9")
	mp9 := ParseNodes(n)[0]

	n, _ = mygraph.GetNode("10")
	mp10 := ParseNodes(n)[0]

	var tests = []struct {
		testname string
		id       string
		err      error
		wantErr  bool
		asc      []data
		dsc      []data
	}{
		{
			testname: "get ancestor and descendants of root",
			id:       "8",
			err:      nil,
			wantErr:  false,
			asc:      make([]data, 0),
			dsc:      []data{mp9, mp10},
		},
		{
			testname: "get ancestor and descendants of a node",
			id:       "9",
			err:      nil,
			wantErr:  false,
			asc:      []data{mp8},
			dsc:      []data{mp10},
		},
		{
			testname: "get ancestor and descendants of a leaf",
			id:       "10",
			err:      nil,
			wantErr:  false,
			asc:      []data{mp9, mp8},
			dsc:      make([]data, 0),
		},
	}

	for _, tt := range tests {
		gotasc, err1 := Ancestors(tt.id)
		gotdsc, err := Descendants(tt.id)

		if err1 != tt.err && err != tt.err {
			t.Errorf("not got right result exp:%v got:%v", err, tt.err)
		}

		if !checkEqual(tt.asc, gotasc) {
			t.Errorf("not got right ascendants exp:%v got:%v", tt.asc, gotasc)
		}
		if !checkEqual(tt.dsc, gotdsc) {
			t.Errorf("not got right descendants exp:%v got:%v", tt.dsc, gotdsc)
		}
	}
}
