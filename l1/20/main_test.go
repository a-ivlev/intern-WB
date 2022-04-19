package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReversWords(t *testing.T) {
	testCases := map[string]struct {
		inputWords string
		wantWords  string
	}{
		"test 1": {
			inputWords: "snow dog sun",
			wantWords:  "sun dog snow",
		},
		"test 2": {
			inputWords: "Hello 世界",
			wantWords:  "世界 Hello",
		},
		"test 3": {
			inputWords: "Hello",
			wantWords:  "Hello",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotWords := ReversWords(tc.inputWords)
			diff := cmp.Diff(tc.wantWords, gotWords)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}