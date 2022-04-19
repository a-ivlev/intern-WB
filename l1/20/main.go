// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
//
package main

import (
	"fmt"
	"strings"
)

func ReversWords(s string) string {
	words := strings.Split(s, " ")
	var revWords strings.Builder

	for i := len(words)-1; i >= 0; i-- {
		revWords.WriteString(words[i])
		if i == 0 {		
			break
		}
		revWords.WriteString(" ")
	}

	return revWords.String()
}

func main() {
	s := "snow dog sun"
	
	fmt.Println(ReversWords(s))


}