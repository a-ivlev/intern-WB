// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
//
package main

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrNotCommand = errors.New("command not implemented")
	ErrDivZero    = errors.New("dividing by zero")
	ErrNotNumber = errors.New("error is not a number")
)

func sum(a, b *big.Int) (string, error) {
	res := new(big.Int)

	bigInt := res.Add(a, b)
	str := fmt.Sprintf("%v", bigInt)

	return str, nil
}

func subtract(a, b *big.Int) (string, error) {
	res := new(big.Int)

	bigInt := res.Sub(a, b)
	str := fmt.Sprintf("%v", bigInt)

	return str, nil
}

func multiply(a, b *big.Int) (string, error) {
	res := new(big.Int)

	bigInt := res.Mul(a, b)
	str := fmt.Sprintf("%v", bigInt)

	return str, nil
}

func divide(a, b *big.Int) (string, error) {
	res := new(big.Int)
	//zero := big.NewInt(0)
	if fmt.Sprintf("%v", b) == "0" {
		return "", ErrDivZero
	}

	bigInt := res.Div(a, b)
	str := fmt.Sprintf("%v", bigInt)

	return str, nil
}

// мапа соответствий названия действия функции-обработчику.
var cmdMap = map[string]func(*big.Int, *big.Int) (string, error){
	"sum":      sum,
	"subtract": subtract,
	"multiply": multiply,
	"divide":   divide,
}

func processCmd(a string, cmd string, b string) (string, error) {

	num1 := new(big.Int)
	if _, ok := num1.SetString(a, 10); !ok {
		return "", ErrNotNumber
	}

	num2 := new(big.Int)
	if _, ok := num2.SetString(b, 10); !ok {
		return "", ErrNotNumber
	}

	if cmdFunc, ok := cmdMap[cmd]; ok {
		return cmdFunc(num1, num2)
	}
	return "", ErrNotCommand
}

func main() {
	cmd := "sum"
	a := "2500000000000000000000"
	b := "2500000000000000000000"

	res, err := processCmd(a, cmd, b)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(res)

	cmd = "subtract"

	a = "5000000000000000000000"
	b = "500000000000000000000"

	res, err = processCmd(a, cmd, b)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(res)

	 cmd = "multiply"

	a = "5000000000000000000000"
	b = "3"

	res, err = processCmd(a, cmd,  b)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(res)

	cmd = "divide"

	a = "5000000000000000000000"
	b = "1000000000000000000000"

	res, err = processCmd(a, cmd, b)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(res)

	cmd = "divide"

	a = "5000000000000000000000"
	b = "0"

	res, err = processCmd(a, cmd, b)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Println(res)
}
