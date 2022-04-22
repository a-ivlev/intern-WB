// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func main() {
	arr1 := []int64{2, 8, 1, 10, 3, 5}
	arr2 := []int64{32, 0, 3, 1}

	res := Intersection(arr1, arr2)
	fmt.Println(res)
}

// Intersection находит пересечения множеств.
func Intersection(arr ...[]int64) []int64 {
	// Создаём map с помощью которой будем искать пересечения двух множеств.
	tmpMap := make(map[int64]uint64)
	// Результирующий массив в который записываются найденные пересечения.
	intersecArr := make([]int64, 0, 100)
	// Итерируемся по множествам.
	for _, sets := range arr {
		// Итерируемся по элементам множества.
		for _, elem := range sets {
			// Добавляем элементы в map и инкрементируем счётчик.
			tmpMap[elem]++
			// Если счётчик больше единицы, значит мы нашли пересечение.
			if tmpMap[elem] > 1 {
				// Добавляем пересечение в результирующий слайс.
				intersecArr = append(intersecArr, elem)
			}
		}
	}
	return intersecArr
}
