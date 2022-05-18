package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnpackStr(t *testing.T) {
	testCases := map[string]struct {
		inputStr string
		want     string
		wantErr  error
	}{
		"test-1 letter and number string": {
			inputStr: "a4bc2d5e",
			want:     "aaaabccddddde",
			wantErr:  nil,
		},
		"test-2 letter string": {
			inputStr: "abcd",
			want:     "abcd",
			wantErr:  nil,
		},
		"test-3 error input": {
			inputStr: "45",
			want:     "",
			wantErr:  ErrInvalidStr,
		},
		"test-4 empty string": {
			inputStr: "",
			want:     "",
			wantErr:  nil,
		},
		"test-5 escape sequences qwe\\4\\5": {
			inputStr: "qwe\\4\\5",
			want:     "qwe45",
			wantErr:  nil,
		},
		"test-6 escape sequences qwe\\45": {
			inputStr: "qwe\\45",
			want:     "qwe44444",
			wantErr:  nil,
		},
		"test-7 escape sequences qwe\\\\5": {
			inputStr: "qwe\\\\5",
			want:     "qwe\\\\\\\\\\",
			wantErr:  nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := unpackStr(tc.inputStr)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
			if tc.wantErr != gotErr {
				t.Fatalf("%s: expected err: %v, got err: %v", name, tc.wantErr, gotErr)
			}
		})
	}
}
