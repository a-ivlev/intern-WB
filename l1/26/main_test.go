package main

import "testing"

func TestUniqString(t *testing.T) {
	testCases := map[string]struct {
		inputStr string
		wantBool bool
	}{
		"test 1 lowercase unique string": {
			inputStr: "yfgcdkhswqфывапро",
			wantBool: true,
		},
		"test 2 unique string": {
			inputStr: "yFgCDkhswqфыВаПро",
			wantBool: true,
		},
		"test 3 non-unique string": {
			inputStr: "yFgCDkhsSwq",
			wantBool: false,
		},
		"test 4 non-unique string": {
			inputStr: "фыВаПроп",
			wantBool: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotBool := UniqString(tc.inputStr)
			if tc.wantBool != gotBool {
            t.Fatalf("%s: expected: %v, got: %v", name, tc.wantBool, gotBool)
        }
		})
	}
}

func TestUniqSortString(t *testing.T) {
	casesTest := map[string]struct {
		inputStr string
		wantBool bool
	}{
		"test 1 lowercase unique string": {
			inputStr: "yfgcdkhswqфывапро",
			wantBool: true,
		},
		"test 2 unique string": {
			inputStr: "yFgCDkhswqфыВаПро",
			wantBool: true,
		},
		"test 3 non-unique string": {
			inputStr: "yFgCDkhsSwq",
			wantBool: false,
		},
		"test 4 non-unique string": {
			inputStr: "фыВаПроп",
			wantBool: false,
		},
	}

	for name, tc := range casesTest {
		t.Run(name, func(t *testing.T) {
			gotBool := UniqSortString(tc.inputStr)
			if tc.wantBool != gotBool {
            t.Fatalf("%s: expected: %v, got: %v", name, tc.wantBool, gotBool)
        }
		})
	}
}