package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProcessCmd(t *testing.T) {
	testCases := map[string]struct {
		cmd     string
		a       string
		b       string
		want    string
		wantErr error
	}{
		"test 1": {
			cmd:     "sum",
			a:       "2500000000000000000000",
			b:       "2500000000000000000000",
			want:    "5000000000000000000000",
			wantErr: nil,
		},
		"test 2": {
			cmd:     "subtract",
			a:       "5000000000000000000000",
			b:       "2000000000000000000000",
			want:    "3000000000000000000000",
			wantErr: nil,
		},
		"test 3": {
			cmd:     "multiply",
			a:       "5000000000000000000000",
			b:       "3",
			want:    "15000000000000000000000",
			wantErr: nil,
		},
		"test 4": {
			cmd:     "divide",
			a:       "15000000000000000000000",
			b:       "5",
			want:    "3000000000000000000000",
			wantErr: nil,
		},
		"test 5": {
			cmd:     "divide",
			a:       "15000000000000000000000",
			b:       "0",
			want:    "",
			wantErr: ErrDivZero,
		},
	}

		for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := processCmd(tc.a, tc.cmd, tc.b)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
			if tc.wantErr != err {
				t.Errorf("%s: want error %s, got error %s", name, tc.wantErr, err)
			}
		})
	}
}
