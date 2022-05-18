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
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// read читает данные из файла или из STDIN, построчно заносит их в слайс
// и подсчитывает количество повторяющихся строк.
func goCut(f io.ReadCloser, options Options) {
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// res слайс добавляются указанные поля строк.
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := procLine(input.Text(), options)
		if line == "" && options.Separated {
			continue
		}
		fmt.Println(line)
	}
}

// procLine разбивает строку на поля по разделителю, формирует новую строку из заданных полей и возвращает её.
func procLine(line string, options Options) string {
	// resLine строка с добавленными полями.
	var resLine strings.Builder

	// Разбиваем входную строку на поля по разделителю.
	splitLines := strings.Split(line, options.Delimiter)

	// Проверяем если в строке нет разделителя, тогда вернётся исходная строка.
	if splitLines[0] == line {
		if options.Separated {
			return ""
		}
		return line
	}
	// Строку разбили на поля по разделителю.
	// Циклом пробегаемся по перечисленным в флаге -f полям.
	for i, numField := range options.Fields {
		// Проверка на существование поля.
		if i == len(splitLines)-2 {
			break
		}
		// Если поле не первое, тогда добавляем разделитель.
		if i != 0 {
			// Если поле не последнее, добавляем разделитель.
			resLine.WriteString(options.Delimiter)
		}

		// Добавляем нужные поля в новую строку.
		resLine.WriteString(splitLines[numField-1])
	}
	return resLine.String()
}

// Options Работа с флагами.
type Options struct {
	Fields    []int  `short:"f" long:"fields" description:"выбрать поля (колонки)"`
	Delimiter string `short:"d" long:"delimiter" default:"\t" description:"использовать другой разделитель"`
	Separated bool   `short:"s" long:"separated" description:"только строки с разделителем"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

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

	var f *os.File
	if len(args) > 0 {
		f, err = os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdin
	}
	defer f.Close()

	goCut(f, options)
}
