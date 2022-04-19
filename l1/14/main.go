// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)


func PrintType(value interface{}) string {
	switch value.(type) {
	case int:
		return "Type: int"
	case string:
		return "Type: string"
	case chan int:
		return "Type: chan int"
	case bool:
		return "Type: bool"
	default:
		return "Unknown variable type"	
	}
}

func main() {
	a, b, c, d := 1, "hello word!", true, make(chan int)
	
	fmt.Println(PrintType(a))
	fmt.Println(PrintType(b))
	fmt.Println(PrintType(c))
	fmt.Println(PrintType(d))
	

	fmt.Println(reflect.TypeOf(d))
}