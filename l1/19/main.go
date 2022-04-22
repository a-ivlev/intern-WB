// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "главрыба"
	fmt.Println(ReversString(str))
}


func ReversString(str string) string {
	runStr := []rune(str)
	var revStr strings.Builder
	revStr.Grow(len(str))

	for i := len(runStr)-1; i >= 0; i-- {
		revStr.WriteRune(runStr[i])
	}
	
	return revStr.String()
}

func ReversString2(str string) string {
	runStr := []rune(str)
	revStr := make([]rune, 0, len(runStr))

	for i := len(runStr)-1; i >= 0; i-- {
		revStr = append(revStr, runStr[i]) 
	}

	return string(revStr)
}