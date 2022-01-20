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
			experr:   fmt.Errorf("node already exists %w", NodeExistsErr),
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
			experr:   fmt.Errorf("parent node does not exists %w", NodeNotFoundErr),
			wantErr:  true,
		},
		{
			testname: "cyclic dependency",
			pid:      "5",
			cid:      "4",
			experr:   fmt.Errorf("cannot add dependency %w", CyclicDependencyErr),
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
		aerr     error
		derr     error
		perr     error
		cerr     error
		wantErr  bool
		asc      []data
		dsc      []data
		par      []data
		chd      []data
	}{
		{
			testname: "get ancestor and descendants of root",
			id:       "8",
			aerr:     nil,
			derr:     nil,
			perr:     nil,
			cerr:     nil,
			wantErr:  false,
			asc:      make([]data, 0),
			dsc:      []data{mp9, mp10},
			chd:      []data{mp9},
			par:      []data{},
		},
		{
			testname: "get ancestor and descendants of a node",
			id:       "9",
			aerr:     nil,
			derr:     nil,
			perr:     nil,
			cerr:     nil,
			wantErr:  false,
			asc:      []data{mp8},
			dsc:      []data{mp10},
			par:      []data{mp8},
			chd:      []data{mp10},
		},
		{
			testname: "get ancestor and descendants of a leaf",
			id:       "10",
			aerr:     nil,
			derr:     nil,
			perr:     nil,
			cerr:     nil,
			wantErr:  false,
			asc:      []data{mp9, mp8},
			dsc:      []data{},
			par:      []data{mp9},
			chd:      []data{},
		},
		{
			testname: "get ancestor and descendants of a node that not exists",
			id:       "60",
			aerr:     fmt.Errorf("cannot find ancestors %w", NodeNotFoundErr),
			derr:     fmt.Errorf("cannot find descendants %w", NodeNotFoundErr),
			perr:     fmt.Errorf("cannot find parents %w", NodeNotFoundErr),
			cerr:     fmt.Errorf("cannot find children %w", NodeNotFoundErr),
			wantErr:  true,
			asc:      []data{mp9, mp8},
			dsc:      []data{},
			par:      []data{mp9},
			chd:      []data{},
		},
	}

	for i, tt := range tests {
		gotAsc, err1 := Ancestors(tt.id)
		gotDsc, err := Descendants(tt.id)
		gotPar, err2 := Parents(tt.id)
		gotChd, err3 := Children(tt.id)

		if tt.wantErr && err1.Error() != tt.aerr.Error() && err.Error() != tt.derr.Error() && err2.Error() != tt.perr.Error() && err3.Error() != tt.cerr.Error() {
			t.Errorf("not got right result exp:%v got:%v", err, tt.aerr)
		}

		if tt.wantErr {
			continue
		}

		if !checkEqual(tt.asc, gotAsc) {
			t.Errorf("%d not got right ascendants exp:%v got:%v", i, tt.asc, gotAsc)
		}
		if !checkEqual(tt.dsc, gotDsc) {
			t.Errorf("not got right descendants exp:%v got:%v", tt.dsc, gotDsc)
		}

		if !checkEqual(tt.par, gotPar) {
			t.Errorf("%d not got right parents exp:%v got:%v", i, tt.dsc, gotDsc)
		}

		if !checkEqual(tt.chd, gotChd) {
			t.Errorf("not got right children exp:%v got:%v", tt.dsc, gotDsc)
		}

	}
}

func TestDeleteNode(t *testing.T) {
	_, err := AddNode("one8", "18")
	if err != nil {
		t.Errorf("node cant be added")
	}

	var test = []struct {
		testname string
		id       string
		msg      string
	}{
		{
			testname: "node deleted",
			id:       "18",
			msg:      "node deleted successfuly",
		},
		{
			testname: "node does not exists",
			id:       "18",
			msg:      "node does not exists",
		},
	}

	for _, tt := range test {
		t.Run(tt.testname, func(t *testing.T) {
			got := DeleteNode(tt.id)
			if got[0]["message"] != tt.msg {
				t.Errorf("not deleted exp: %v got: %v", tt.msg, got)
			}
		})
	}
}

func TestDeleteDependency(t *testing.T) {
	_, err := AddNode("one7", "17")
	if err != nil {
		t.Errorf("node cant be added")
	}
	_, err = AddNode("one9", "19")
	if err != nil {
		t.Errorf("node cant be added")
	}

	AddDependency("17", "18")

	var test = []struct {
		testname string
		pid      string
		cid      string
		msg      string
		wantErr  error
	}{
		{
			testname: "node deleted",
			pid:      "17",
			cid:      "19",
			msg:      "dependency deleted successfuly",
			wantErr:  nil,
		},
		{
			testname: "node does not exists",
			pid:      "20",
			cid:      "17",
			msg:      "parent node does not exists",
			wantErr:  fmt.Errorf("cannot remove dependency %w", NodeNotFoundErr),
		},
		{
			testname: "node does not exists",
			pid:      "17",
			cid:      "27",
			msg:      "child node does not exists",
			wantErr:  fmt.Errorf("cannot remove dependency %w", NodeNotFoundErr),
		},
	}

	for _, tt := range test {
		t.Run(tt.testname, func(t *testing.T) {
			got, err := DeleteDependency(tt.pid, tt.cid)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("exp %v got %v as error", tt.wantErr, err)
				} else {
					if tt.wantErr.Error() != err.Error() {
						t.Errorf("exp %v got %v as error", tt.wantErr, err)
					}
				}
			} else {
				if got[0]["message"] != tt.msg {
					t.Errorf("not deleted exp: %v got: %v", tt.msg, got)
				}
			}

		})
	}
}
