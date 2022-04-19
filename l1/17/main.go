// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptySlice          = errors.New("passed an empty slice")
	ErrNumberOut           = errors.New("number out of range")
	ErrElemNotFound        = errors.New("element not found")
)

// BinarySearch осуществляет поиск в упорядоченном слайсе из n-элементов целых чисел, заданного числа.
// BinarySearch возвращает индекс найденного элемента в слайсе и ошибку. Индекс -1 означает, что
// заданного элемента в данном слайсе не обнаружено.
func BinarySearch(arr []int64, x int64) (int, error) {
	if len(arr) == 0 {
		return -1, ErrEmptySlice
	}

	first := 0
	last := len(arr) - 1
	mid := 0

	if (arr[first] > x) || (arr[last] < x) {
		return -1, ErrNumberOut
	}

	switch {
	// Проверяем благоприятный случай. Искомый элемент первый элемент слайса.
	case arr[first] == x:
		return first, nil
	// Проверяем благоприятный случай. Искомый элемент последний элемент слайса.
	case arr[last] == x:
		return last, nil
	}

	for first <= last {
		mid = first + (last - first) / 2
		if arr[mid] == x {
			return mid, nil
		}
		if arr[mid] < x {
			first = mid+1
			continue
		}
		if arr[mid] > x {
			last = mid-1
		}
		
	}

	return -1, ErrElemNotFound
}

func main() {

	c := []int64{1, 3, 5, 7, 9, 10, 13, 15, 25, 27, 31, 45, 55, 71, 72, 78, 81, 86, 92, 95}

	fmt.Println("Будем искать число 27 находящееся в середине слайса. Индекс в слайсе 8.")
	idx, err := BinarySearch(c, 25)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Поиск проходил в следующем слайсе", c)
	fmt.Printf("Искомый элемент %d найден, он находиться по индексу %d в слайсе.\nОбщее количество элементов в переданном слайсе %d.\n", c[idx], idx, len(c))

}
