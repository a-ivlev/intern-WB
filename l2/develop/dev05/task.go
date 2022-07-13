package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Options Работа с флагами.
type Options struct {
	After      int  `short:"A" long:"after" description:"печатать +N строк после совпадения"`
	Before     int  `short:"B" long:"before" description:"печатать +N строк до совпадения"`
	Context    int  `short:"C" long:"context" description:"(A+B) печатать ±N строк вокруг совпадения"`
	Count      bool `short:"c" long:"count" description:"количество строк"`
	IgnoreCase bool `short:"i" long:"ignore-case" description:"игнорировать регистр"`
	Invert     bool `short:"v" long:"invert" description:"вместо совпадения, исключать"`
	Fixed      bool `short:"F" long:"fixed" description:"точное совпадение со строкой, не паттерн"`
	LineNum    bool `short:"n" long:"line-num" description:"печатать номер строки"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

// grep читает строки построчно и заносит их в слайс, проверяет на совпадение и подсчитывает количество строк.
// Возвращает два слайса и общее количество прочитанных строк. Первый это полный слайс allInput,
// в котором находятся все прочитанные строки. Второй слайс resIdx содержит индексы строк, по которым найдены совпадения.
func grep(f io.ReadCloser, findStr string, options Options) (allInput []string, resIdx []int, count int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		// allInput в этот слайс добавляются все прочитанные строки.
		allInput = append(allInput, line)
		// ignoreCase если этот флаг установлен, мы переводим строки в нижний регистр,
		// и сравниваем строки. Так мы игнорируем регистр входной и искомой строки.
		if options.IgnoreCase && !options.Fixed {
			line = strings.ToLower(line)
			findStr = strings.ToLower(findStr)
		}
		// Проверяем содержит ли текущая строка искомую подстроку.
		if strings.Contains(line, findStr) && !options.Fixed {
			// Если совпадение найдено добавляем индекс строки в res слайс.
			// res слайс для хранения индексов строк в которых обнаружено совпадение с искомой строкой.
			resIdx = append(resIdx, count)
		}
		if line == findStr && options.Fixed {
			// res слайс для хранения индексов строк в которых обнаружено совпадение с искомой строкой.
			resIdx = append(resIdx, count)
		}
		// count переменная которая хранит значение количества строк.
		count++
	}
	return
}

// printRes с учётом полученных флагов подготавливает результат для вывода в STDOUT.
func printRes(allInput []string, resIdx []int, options Options) []string {
	var resStrings []string
	// Условие выполниться если установлен флаг -v вместо совпадения, исключать.
	if options.Invert {
		var j int
		// Итерируемся по слайсу который содержит все прочитанные строки.
		for i, elem := range allInput {
			switch {
			// Когда установлен флаг -v нам необходимо не выводить строки в которых найдены совпадения.
			// Когда текущий индекс совпдает со значением из слайса resIdx (хранит индексы строк в которых найдены
			// совпадения) мы не добавляем это значение в итоговый слайс, и переходим к проверке следующего значения.
			case i == resIdx[j]:
				if j < len(resIdx)-1 {
					j++
				}
			case options.LineNum:
				resStrings = append(resStrings, fmt.Sprintf("%d %s", i+1, elem))
				//fmt.Println(i+1, elem)
			default:
				resStrings = append(resStrings, elem)
				//fmt.Println(elem)
			}
		}
		return resStrings
	}

	var startIdx, stopIdx int
	for _, idx := range resIdx {
		switch {
		case options.Before > 0:
			if startIdx < idx-options.Before {
				startIdx = idx - options.Before
			}
			if startIdx < 0 {
				startIdx = 0
			}
			stopIdx = idx
		case options.After > 0:
			startIdx = idx
			if stopIdx < idx+options.After {
				stopIdx = idx + options.After
			}
			if stopIdx > len(allInput)-1 {
				stopIdx = len(allInput) - 1
			}
		case options.Context > 0:
			startIdx = idx - options.Context
			if startIdx < 0 {
				startIdx = 0
			}

			if stopIdx < idx+options.Context {
				stopIdx = idx + options.Context
			}
			if stopIdx > len(allInput)-1 {
				stopIdx = len(allInput) - 1
			}
		default:
			startIdx = idx
			stopIdx = idx
			//resStrings = append(resStrings, fmt.Sprintf("%s\n", allInput[idx]))
			//fmt.Println(allInput[idx])
		}
		for ; startIdx <= stopIdx; startIdx++ {
			if options.LineNum {
				resStrings = append(resStrings, fmt.Sprintf("%d %s", startIdx+1, allInput[startIdx]))
				//fmt.Println(idx+1, allInput[idx])
			} else {
				resStrings = append(resStrings, allInput[startIdx])
				//fmt.Println(allInput[idx])
			}
		}
	}

	return resStrings
}

func main() {
	// Работа с флагами.
	var args []string
	var err error

	if args, err = parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	inputFilename := args[0]
	findStr := args[1]

	f, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	allInput, resIdx, count := grep(f, findStr, options)
	fmt.Println(resIdx)

	resSlise := printRes(allInput, resIdx, options)
	for _, line := range resSlise {
		fmt.Println(line)
	}

	if options.Count {
		fmt.Println("count read line:", count)
	}
}
