package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAnagramSearch(t *testing.T) {
	testCases := map[string]struct {
		dictionary  *[]string
		wantAnagram *map[string]*[]string
		wantErr     error
	}{
		"test-1 success": {
			dictionary: &[]string{"пятка", "слиток", "тяпка", "листок", "тряпка", "Свисток", "пятак", "столик"},
			wantAnagram: &map[string]*[]string{
				"пятка":  {"пятак", "пятка", "тяпка"},
				"слиток": {"листок", "слиток", "столик"},
			},
			wantErr: nil,
		},
		"test-2 no anagram": {
			dictionary:  &[]string{"тяпка", "листок", "тряпка", "Свисток", "пятёрка"},
			wantAnagram: &map[string]*[]string{},
			wantErr:     nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := anagramSearch(tc.dictionary)
			diff := cmp.Diff(tc.wantAnagram, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
