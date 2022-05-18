package main

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	testCases := map[string]struct {
		inputText  io.ReadCloser
		wantOutput []string
		wantErr    error
	}{
		"test-1 one result": {
			inputText: ioutil.NopCloser(strings.NewReader(`
123 asd 72
567 foo 12
123 asd 72
321 abc 57
789 xyz 100
789 xyz 100
789 xyz 100
123 asd 72
`)),
			wantOutput: []string{"", "123 asd 72", "567 foo 12", "123 asd 72", "321 abc 57", "789 xyz 100", "789 xyz 100", "789 xyz 100", "123 asd 72"},
			wantErr:    nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := read(tc.inputText)
			diff := cmp.Diff(tc.wantOutput, got)
			if diff != "" {
				t.Fatalf(diff)
			}
			if tc.wantErr != gotErr {
				t.Fatalf("%s: expected err: %v, got err: %v", name, tc.wantErr, gotErr)
			}
		})
	}
}

func TestSortByColumn(t *testing.T) {
	testCases := map[string]struct {
		inputLines []string
		columNum   int
		byNumber   bool
		wantOutput []string
	}{
		"test-1 sort by 1 column": {
			inputLines: []string{"123 asd 72", "567 foo 12", "123 asd 72", "321 abc 57", "789 xyz 100", "789 xyz 100", "789 xyz 100", "123 asd 72"},
			columNum:   1,
			byNumber:   false,
			wantOutput: []string{"123 asd 72", "123 asd 72", "123 asd 72", "321 abc 57", "567 foo 12", "789 xyz 100", "789 xyz 100", "789 xyz 100"},
		},
		"test-2 sort by 2 column": {
			inputLines: []string{"123 asd 72", "567 foo 12", "123 asd 72", "321 abc 57", "789 xyz 100", "789 xyz 100", "789 xyz 100", "123 asd 72"},
			columNum:   2,
			byNumber:   false,
			wantOutput: []string{"321 abc 57", "123 asd 72", "123 asd 72", "123 asd 72", "567 foo 12", "789 xyz 100", "789 xyz 100", "789 xyz 100"},
		},
		"test-3 sort by 3 column number true": {
			inputLines: []string{"123 asd 72", "567 foo 12", "123 asd 72", "321 abc 57", "789 xyz 100", "789 xyz 100", "789 xyz 100", "123 asd 72"},
			columNum:   3,
			byNumber:   true,
			wantOutput: []string{"567 foo 12", "321 abc 57", "123 asd 72", "123 asd 72", "123 asd 72", "789 xyz 100", "789 xyz 100", "789 xyz 100"},
		},
		"test-4 sort by 3 column number false": {
			inputLines: []string{"123 asd 72", "567 foo 12", "123 asd 72", "321 abc 57", "789 xyz 100", "789 xyz 100", "789 xyz 100", "123 asd 72"},
			columNum:   3,
			byNumber:   false,
			wantOutput: []string{"789 xyz 100", "789 xyz 100", "789 xyz 100", "567 foo 12", "321 abc 57", "123 asd 72", "123 asd 72", "123 asd 72"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			sortByColumn(tc.inputLines, tc.columNum, tc.byNumber)
			diff := cmp.Diff(tc.wantOutput, tc.inputLines)
			if diff != "" {
				t.Fatalf(diff)
			}

		})
	}
}

func TestRevers(t *testing.T) {
	testCases := map[string]struct {
		inputLines []string
		columNum   int
		byNumber   bool
		wantOutput []string
	}{
		"test-1 sort by 1 column": {
			inputLines: []string{"123 asd 72", "123 asd 72", "123 asd 72", "321 abc 57", "567 foo 12", "789 xyz 100", "789 xyz 100", "789 xyz 100"},
			wantOutput: []string{"789 xyz 100", "789 xyz 100", "789 xyz 100", "567 foo 12", "321 abc 57", "123 asd 72", "123 asd 72", "123 asd 72"},
		},
		"test-2 sort by 2 column": {
			inputLines: []string{"321 abc 57", "123 asd 72", "123 asd 72", "123 asd 72", "567 foo 12", "789 xyz 100", "789 xyz 100", "789 xyz 100"},
			wantOutput: []string{"789 xyz 100", "789 xyz 100", "789 xyz 100", "567 foo 12", "123 asd 72", "123 asd 72", "123 asd 72", "321 abc 57"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			revers(tc.inputLines)
			diff := cmp.Diff(tc.wantOutput, tc.inputLines)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
