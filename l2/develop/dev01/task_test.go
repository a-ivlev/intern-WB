package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNtpTime(t *testing.T) {
	testCases := map[string]struct {
		wantErr error
	}{
		"test err == nil": {
			wantErr: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := ntpTime()
			diff := cmp.Diff(tc.wantErr, err)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
