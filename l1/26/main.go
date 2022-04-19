// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	str := "SvhfDkeYWqs"
	fmt.Println(UniqString(str))
	fmt.Println(UniqSortString(str))

}

// 1 Способ решения
// UniqString переводит строку в нижний регистр. Итерируется по строке,
// проверяет есть ли в map новый символ, если нет, тогда добавляет новые символ в map.
// Если есть, возвращает false. Если строка закончилась и все символы уникальны,
// возвращается true.
func UniqString(str string) bool {
	str = strings.ToLower(str)

	checkMap := make(map[rune]struct{}, len(str))
	for _, simvol := range str {
		if _, ok := checkMap[simvol]; !ok {
			checkMap[simvol] = struct{}{}
			continue
		}
		return false
	}
	return true
}

// Способ 2.
// UniqSortString переводит строку в нижний регистр, сортирует строку и проверяет,
// находятся рядом 2 одинаковых символа. Если находятся возвращается false, если рядом
// нет одинаковых символов возвращается true.
func UniqSortString(str string) bool {
	str = strings.ToLower(str)
	sortStr := []rune(str)
	sort.Slice(sortStr, func(i, j int) bool {
		return sortStr[i] < sortStr[j]
	})
	for i, sim := range sortStr {
		if i > 0 && sortStr[i-1] == sim {
			return false
		}
	}
	return true
}