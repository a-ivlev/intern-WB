package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProcLine(t *testing.T) {
	testCases := map[string]struct {
		inputLine string
		options   Options
		delimiter string
		fields    []int
		wantLine  string
	}{
		"test-1 two fields": {
			inputLine: "Winter: white: snow: frost",
			options: Options{
				Delimiter: ":",
				Fields:    []int{1, 3},
			},
			wantLine: "Winter: snow",
		},
		"test-2 one fields": {
			inputLine: "Winter: white: snow: frost",
			options: Options{
				Delimiter: ":",
				Fields:    []int{2},
			},
			wantLine: " white",
		},
		"test-3 no one field": {
			inputLine: "Winter: white: frost",
			options: Options{
				Delimiter: ":",
				Fields:    []int{2, 4},
			},
			wantLine: " white",
		},
		"test-4 no separated": {
			inputLine: "Spring grass warm",
			options: Options{
				Delimiter: ":",
				Fields:    []int{2, 4},
			},
			wantLine: "Spring grass warm",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotLine := procLine(tc.inputLine, tc.options)
			diff := cmp.Diff(tc.wantLine, gotLine)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

//func TestRead(t *testing.T) {
//	testCases := map[string]struct {
//		inputText  io.ReadCloser
//		wantOutput []string
//		wantErr    error
//	}{
//		"test-1 one result": {
//			inputText: ioutil.NopCloser(strings.NewReader(`
//Winter: white: snow: frost
//Spring: green: grass: warm
//Summer: colorful: blossom: hot
//Autumn: yellow: leaves: cool
//`)),
//			wantOutput: []string{"", "123 asd 72", "567 foo 12", "123 asd 72", "321 abc 57", "789 xyz 100", "789 xyz 100", "789 xyz 100", "123 asd 72"},
//			wantErr:    nil,
//		},
//	}
//
//	for name, tc := range testCases {
//		t.Run(name, func(t *testing.T) {
//			got, gotErr := read(tc.inputText)
//			diff := cmp.Diff(tc.wantOutput, got)
//			if diff != "" {
//				t.Fatalf(diff)
//			}
//			if tc.wantErr != gotErr {
//				t.Fatalf("%s: expected err: %v, got err: %v", name, tc.wantErr, gotErr)
//			}
//		})
//	}
//}
