package main

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGrep(t *testing.T) {
	testCases := map[string]struct {
		inputText    io.ReadCloser
		inputFindStr string
		options      Options
		wantAllInput []string
		wantResIdx   []int
		wantCount    int
	}{
		"test-1 one result": {
			inputText: ioutil.NopCloser(strings.NewReader(`
Наша Таня громко плачет,
Уронила в речку мячик,
Тише Танечка не плач,
Не утонет в речке мяч.
`)),
			inputFindStr: "мячик",
			options:      Options{},
			wantAllInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			wantResIdx:   []int{2},
			wantCount:    5,
		},
		"test-2 two result": {
			inputText: ioutil.NopCloser(strings.NewReader(`
Наша Таня громко плачет,
Уронила в речку мячик,
Тише Танечка не плач,
Не утонет в речке мяч.
`)),
			inputFindStr: "мяч",
			wantAllInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			wantResIdx:   []int{2, 4},
			wantCount:    5,
		},
		"test-3 no result": {
			inputText: ioutil.NopCloser(strings.NewReader(`
Наша Таня громко плачет,
Уронила в речку мячик,
Тише Танечка не плач,
Не утонет в речке мяч.
`)),
			inputFindStr: "мяч",
			wantAllInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				Fixed: true,
			},
			wantResIdx: nil,
			wantCount:  5,
		},
		"test-4 one result": {
			inputText: ioutil.NopCloser(strings.NewReader(`
Наша Таня громко плачет,
Уронила в речку мячик,
Тише Танечка не плач,
Не утонет в речке мяч.
`)),
			inputFindStr: "Уронила в речку мячик,",
			wantAllInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				Fixed: true,
			},
			wantResIdx: []int{2},
			wantCount:  5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotAllInput, gotResIndx, gotCount := grep(tc.inputText, tc.inputFindStr, tc.options)
			diff := cmp.Diff(tc.wantAllInput, gotAllInput)
			if diff != "" {
				t.Fatalf(diff)
			}
			diff = cmp.Diff(tc.wantResIdx, gotResIndx)
			if diff != "" {
				t.Fatalf(diff)
			}
			diff = cmp.Diff(tc.wantCount, gotCount)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestPrintRes(t *testing.T) {
	testCases := map[string]struct {
		allInput   []string
		resIdx     []int
		options    Options
		wantResult []string
	}{
		"test-1 no flags": {
			allInput:   []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			resIdx:     []int{2},
			wantResult: []string{"Уронила в речку мячик,"},
		},
		"test-2 flag line-num": {
			allInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				LineNum: true,
			},
			resIdx:     []int{2},
			wantResult: []string{"3 Уронила в речку мячик,"},
		},
		"test-3 flags line-num and before = 1": {
			allInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				LineNum: true,
				Before:  2,
			},
			resIdx:     []int{3},
			wantResult: []string{"2 Наша Таня громко плачет,", "3 Уронила в речку мячик,", "4 Тише Танечка не плач,"},
		},
		"test-4 flags line-num and after 2": {
			allInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				LineNum: true,
				After:   2,
			},
			resIdx:     []int{1},
			wantResult: []string{"2 Наша Таня громко плачет,", "3 Уронила в речку мячик,", "4 Тише Танечка не плач,"},
		},
		"test-5 flags line-num and after 2": {
			allInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				LineNum: true,
				After:   2,
			},
			resIdx:     []int{3},
			wantResult: []string{"4 Тише Танечка не плач,", "5 Не утонет в речке мяч."},
		},
		"test-5 flags line-num and context 2": {
			allInput: []string{"", "Наша Таня громко плачет,", "Уронила в речку мячик,", "Тише Танечка не плач,", "Не утонет в речке мяч."},
			options: Options{
				LineNum: true,
				Context: 2,
			},
			resIdx:     []int{3},
			wantResult: []string{"2 Наша Таня громко плачет,", "3 Уронила в речку мячик,", "4 Тише Танечка не плач,", "5 Не утонет в речке мяч."},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotResult := printRes(tc.allInput, tc.resIdx, tc.options)
			diff := cmp.Diff(tc.wantResult, gotResult)
			if diff != "" {
				t.Fatalf(diff)
			}

		})
	}
}
