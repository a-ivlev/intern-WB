// Удалить i-ый элемент из слайса.
package main

import "fmt"

func main() {

	arr := []int64{1, 2, 3, 4, 5}
	idxDel := 2

	got := DelElem(arr, idxDel)

	fmt.Println(got)

}
// Способ 1
func DelElem(arr []int64, idx int) []int64 {
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}

// Способ 2
func DelElem2(arr []int64, idx int) []int64 {
	arr1 := make([]int64, 0, len(arr)-1)
	arr1 = arr[:idx]
	arr1 = append(arr1, arr[idx+1:]...)
	return arr1
}

// Способ 2
func DelElemEnumeration(arr []int64, idx int) []int64 {
	arr1 := make([]int64, 0, len(arr)-1)
	for i, elem := range arr {
		if i == idx {
			continue
		}
		arr1 = append(arr1, elem)
	}
	return arr1
}
