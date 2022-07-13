package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// ErrInvalidStr некорректная строка.
var ErrInvalidStr = errors.New("invalid string")

// unpackStr
func unpackStr(str string) (string, error) {
	// формируем новую строку.
	var res strings.Builder
	// Резервируем объём памяти под новую строку, равный двойной длине строки.
	res.Grow(len(str) * 2)
	// Преобразуем исходную строку в руны.
	runStr := []rune(str)

	// Проверка на пустую строку.
	if len(runStr) == 0 {
		return "", nil
	}

	// Если строка начинается с цифры, возвращаем ошибку.
	if unicode.IsDigit(runStr[0]) {
		return "", ErrInvalidStr
	}

	// Итерируемся по символам.
	for i := 0; i < len(runStr); i++ {
		// Реализация escape - последовательностей, если символ '\' пропускаем его.
		if runStr[i] == '\\' {
			i++
		}

		switch {
		// Добавляем в новую строку последний символ.
		case i+1 == len(str):
			res.WriteRune(runStr[i])
		// Добавляем в новую строку символ, если он строковый или второй '\'.
		case unicode.IsLetter(runStr[i+1]) || runStr[i+1] == '\\':
			res.WriteRune(runStr[i])
		// Если текущий символ цифра.
		case unicode.IsDigit(runStr[i+1]):
			number := new([]rune)
			var j int
			// Так как число может быть многозначным, получаем его строковое представление в цикле.
			for j = i + 1; j < len(str); j++ {
				// Цикл продолжается пока не будет найден строковый символ.
				if unicode.IsLetter(runStr[j]) {
					break
				}
				// Формируем строковое представление числа.
				*number = append(*number, runStr[j])
			}
			// Переводим полученную строку в число.
			count, _ := strconv.Atoi(string(*number))
			// В цикле добавляем символ, полученное количество раз.
			for x := 1; x <= count; x++ {
				res.WriteRune(runStr[i])
			}
			// Присваиваем i новое значение индекса.
			i = j - 1
		}
	}
	return res.String(), nil
}

func main() {
	res, err := unpackStr("qwe\\4\\5")
	if err != nil {
		log.Fatalf("%s: %s", "error unpuclStr", err)
	}
	fmt.Println(res)
}
