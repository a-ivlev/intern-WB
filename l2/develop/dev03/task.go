package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// read читает данные из файла или из STDIN, построчно заносит их в слайс
// и подсчитывает количество повторяющихся строк.
func read(f io.ReadCloser) ([]string, error) {
	// counts карта - нужна для подсчёта количества повторяющихся строк.
	// Строка заноситься в карту, и счётчик увеличивается на единицу.
	counts := make(map[string]int)
	// res результирующий слайс в него добавляются прочитанные строки.
	var res []string
	input := bufio.NewScanner(f)
	for input.Scan() {
		// Подсчёт повторяемости строки.
		counts[input.Text()]++
		res = append(res, input.Text())
	}
	return res, nil
}

// sortByColumn реализует сортировку по выбранной колонке с учётом флага -h сортировать по числовому значению.
func sortByColumn(lines []string, ColumnNumber int, sortByNumber bool) {
	check := strings.Split(lines[0], " ")
	// Если заданный номер колонки меньше 1, считаем что нужно сортировать по 1 колонке.
	if ColumnNumber < 1 {
		ColumnNumber = 1
	}
	// Если заданный номер колонки больше колличества колонок, считаем что нужно сортировать по последней колонке.
	if ColumnNumber > len(check) {
		ColumnNumber = len(check)
	}
	sort.Sort(stringSlice{
		inputSlice:   lines,
		columnNumber: ColumnNumber,
		separate:     " ",
		sortByNumber: sortByNumber,
	})
}

// stringSlice хранит входные параметры, методы реализуют интерфейс sort,
// чтобы можно было прописать свою логику сортировки.
type stringSlice struct {
	inputSlice   []string
	columnNumber int
	separate     string
	sortByNumber bool
}

// Len возвращает длину исходного слайса.
func (s stringSlice) Len() int { return len(s.inputSlice) }

// Less проверяет переданные параметры сортировки, и сравнивает два значения.
func (s stringSlice) Less(i, j int) bool {
	// Проверка, если не задан номер колонки для сортировки,
	// сортируем по первой колонке.
	if s.inputSlice[i] == "" || s.inputSlice[j] == "" {
		s.columnNumber = 1
	}
	// Из слайса строк выбирается строка с индексом i и индексом j. Затем эти строки разбираются на поля по указанному
	// сепаратору. Переменной left присваивается разобранная строка полученная по индексу i.
	left := strings.Split(s.inputSlice[i], s.separate)
	// Переменной right присваивается разобранная строка полученная по индексу j.
	right := strings.Split(s.inputSlice[j], s.separate)

	// Условие выполняется если указан флаг -h — сортировать по числовому значению.
	if s.sortByNumber {
		var err error
		var leftInt, rightInt int
		// Из разобранной строки left, получаем значение поля указанного для сортировки, и переводим его значение
		// в тип int, и присваиваем переменной leftInt.
		leftInt, err = strconv.Atoi(left[s.columnNumber-1])
		// Проверка на ошибку корректность перевода.
		if err != nil {
			return left[s.columnNumber-1] < right[s.columnNumber-1]
		}
		// Из разобранной строки right, получаем значение поля указанного для сортировки, и переводим его значение
		// в тип int, и присваиваем переменной rightInt.
		rightInt, err = strconv.Atoi(right[s.columnNumber-1])
		// Проверка на ошибку корректность перевода.
		if err != nil {
			return left[s.columnNumber-1] < right[s.columnNumber-1]
		}
		// Возвращаем результат сравнения двух переменных.
		return leftInt < rightInt
	}
	// Возвращаем результат сравниваются значения полей указанных для сортировки.
	return left[s.columnNumber-1] < right[s.columnNumber-1]
}

// Swap меняет местами в слайсе две строки.
func (s stringSlice) Swap(i, j int) {
	s.inputSlice[i], s.inputSlice[j] = s.inputSlice[j], s.inputSlice[i]
}

// revers переворачивает в обратном порядке отсортированный слайс.
func revers(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// uniq удаляет повторяющиеся строки в уже отсортированном по заданному полю слайсе.
func uniq(lines []string) []string {
	// Итерируемся по слайсу.
	for i := 0; i < len(lines)-1; i++ {
		// Условие выполниться если найдётся повторяющаяся строка.
		if lines[i] == lines[i+1] {
			var j int
			// Дополнительный цикл подсчитывает количество повторяющихся строк.
			for j = i; lines[i] == lines[j] && j < len(lines)-1; j++ {
			}
			// Проверка выполняется если достигли конца списка строк.
			if j == len(lines)-1 {
				// Удаление повторяющихся строк до конца списка.
				lines = lines[:i+1]
				break
			}
			// Удаление повторяющихся строк если конец списка ещё не достигнут.
			lines = append(lines[:i+1], lines[j:]...)
			i = 0
		}
	}
	return lines
}

// Options Работа с флагами
type Options struct {
	ColumnNumber int  `short:"k" long:"key" default:"1" description:"указание колонки для сортировки"`
	SortByNumber bool `short:"n" long:"numeric-sort" description:"сортировать по числовому значению"`
	SortRevers   bool `short:"r" long:"reverse" description:"сортировать в обратном порядке"`
	SortUnique   bool `short:"u" long:"unique" description:"не выводить повторяющиеся строки"`
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
	// Проверка передачи в аргумент пути и названия файла.
	if len(args) > 0 {
		// Открываем переданный файл.
		f, err = os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdin
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Читаем данные из файла или если файл не указан с консоли.
	listStr, err := read(f)
	if err != nil {
		log.Fatal(err)
	}
	sortByColumn(listStr, options.ColumnNumber, options.SortByNumber)

	// Если установлен флаг -r переворачиваем отсортированный слайс.
	if options.SortRevers {
		revers(listStr)
	}

	var resSlice = make([]string, 0, len(listStr))
	// Проверка установки флага -q.
	if options.SortUnique {
		// усли флаг установлен, удаляем дублирующиеся строки.
		resSlice = uniq(listStr)
	} else {
		resSlice = listStr
	}

	for _, str := range resSlice {
		fmt.Println(str)
	}
}
