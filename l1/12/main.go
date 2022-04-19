// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
package main

import "fmt"

func main() {
	arrStr := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём множество на базе карты в которой ключом будут строки.
	// Значение нам не важно поэтому используем пустую структуру.
	setStr := make(map[string]struct{}, len(arrStr))

	for _, str := range arrStr {
		if _, ok := setStr[str]; ok {
			continue
		}
		setStr[str] = struct{}{}
	}

	fmt.Println(setStr)
}