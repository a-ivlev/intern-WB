package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuicksort(t *testing.T) {
	testCases := map[string]struct {
		inputArr []int
		wantArr  []int
	}{
		"test 1": {
			inputArr: []int{3, 4, 1, 2, 5, 7, -1, 0},
			wantArr:  []int{-1, 0, 1, 2, 3, 4, 5, 7},
		},
		"test 2": {
			inputArr: []int{17, -24, 20, -15, -21, -3, 13},
			wantArr:  []int{-24, -21, -15, -3, 13, 17, 20},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			Quicksort(tc.inputArr)
			diff := cmp.Diff(tc.wantArr, tc.inputArr)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
