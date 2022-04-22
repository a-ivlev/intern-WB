package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDelElem(t *testing.T) {
	casesTest := map[string]struct {
		arr     []int64
		idxDel  int
		wantArr []int64
		errWant error
	}{
		"test 1 deleting i-th element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  2,
			wantArr: []int64{1, 2, 4, 5},
		},
		"test 2 deleting first element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  0,
			wantArr: []int64{2, 3, 4, 5},
		},
		"test 3 deleting last element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  4,
			wantArr: []int64{1, 2, 3, 4},
		},
	}

	for name, tc := range casesTest {
		t.Run(name, func(t *testing.T) {
			gotArr := DelElem(tc.arr, tc.idxDel)
			diff := cmp.Diff(tc.wantArr, gotArr)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestDelElem2(t *testing.T) {
	casesTest := map[string]struct {
		arr     []int64
		idxDel  int
		wantArr []int64
		errWant error
	}{
		"test 1 deleting i-th element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  2,
			wantArr: []int64{1, 2, 4, 5},
		},
		"test 2 deleting first element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  0,
			wantArr: []int64{2, 3, 4, 5},
		},
		"test 3 deleting last element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  4,
			wantArr: []int64{1, 2, 3, 4},
		},
	}

	for name, tc := range casesTest {
		t.Run(name, func(t *testing.T) {
			gotArr := DelElem2(tc.arr, tc.idxDel)
			diff := cmp.Diff(tc.wantArr, gotArr)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestDelElemEnumeration(t *testing.T) {
	testCases := map[string]struct {
		arr     []int64
		idxDel  int
		wantArr []int64
		errWant error
	}{
		"test 1 deleting i-th element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  2,
			wantArr: []int64{1, 2, 4, 5},
		},
		"test 2 deleting first element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  0,
			wantArr: []int64{2, 3, 4, 5},
		},
		"test 3 deleting last element": {
			arr:     []int64{1, 2, 3, 4, 5},
			idxDel:  4,
			wantArr: []int64{1, 2, 3, 4},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotArr := DelElemEnumeration(tc.arr, tc.idxDel)
			diff := cmp.Diff(tc.wantArr, gotArr)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
