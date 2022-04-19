package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReversString(t *testing.T) {
	testCases := map[string]struct {
		inputStr string
		wantStr  string
	}{
		"test 1": {
			inputStr: "главрыба",
			wantStr:  "абырвалг",
		},
		"test 2": {
			inputStr: "шалаш",
			wantStr:  "шалаш",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotStr := ReversString(tc.inputStr)
			diff := cmp.Diff(tc.wantStr, gotStr)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
