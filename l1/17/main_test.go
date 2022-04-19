package main

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := map[string]struct{
		name string
		arr []int64
		x int64
		want int
		wantErr error
	}{
		"test 1": {
			name: "поиск числа находящегося в начале слайса",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 5,
			want: 2,
			wantErr: nil,
		},
		"test 2": {
			name: "поиск числа находящегося в середине слайса",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 25,
			want: 8,
			wantErr: nil,
		},
		"test 3": {
			name: "поиск числа находящегося в конце слайса",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 86,
			want: 17,
			wantErr: nil,
		},
		"test 4": {
			name: "поиск числа находящегося в середине слайса индекс которого определяется срузу.",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 27,
			want: 9,
			wantErr: nil,
		},
		"test 5": {
			name: "поиск числа которого нет в переданом слайсе.",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 30,
			want: -1,
			wantErr: ErrElemNotFound,
		},
		"test 6": {
			name: "поиск числа которое больше, последнего числа слайса.",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 105,
			want: -1,
			wantErr: ErrNumberOut,
		},
		"test 7": {
			name: "поиск числа которое меньше, первого числа слайса.",
			arr: []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95},
			x: 0,
			want: -1,
			wantErr: ErrNumberOut,
		},
		"test 8": {
			name: "передан пустой слайс.",
			arr: []int64{},
			x: 25,
			want: -1,
			wantErr: ErrEmptySlice,
		},
	}

	for name, tc := range testCases {
		got, gotErr := BinarySearch(tc.arr, tc.x)
        if tc.want != got {
            t.Fatalf("%s %s: expected: %v, got: %v", name, tc.name, tc.want, got)
        }
		if tc.wantErr != gotErr {
			t.Fatalf("%s %s: expected err: %v, got err: %v", name, tc.name, tc.wantErr, gotErr)
		}
	}
}
