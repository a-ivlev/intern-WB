// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).
package main

import (
	"fmt"
)

// Human структура описывает человека и имеет метод получения
// форматированной строки.
type Human struct {
	ID         int64 `json:"id"`
	Name       string    `json:"name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Age        int64     `json:"age"`
	Hobby      []string  `json:"hobby"`
}

// PrintHuman метод структуры Human формирует форматированную строку
// информации о конкретном человеке.
func (h Human) PrintHuman() string {
	return fmt.Sprintf("ID %d\nФамилия: %s\nИмя: %s\nОтчество: %s\nВозраст: %d\nХобби: %v\n",
		h.ID, h.LastName, h.Name, h.MiddleName, h.Age, h.Hobby)
}

// В структуру Action мы встроили структуру Human, а вместе с ней и все её методы.
// У разных структур методы могут называться одинакого. Так же у разных структур
// поля могут иметь одинаковые названия. Поэтому, чтобы обратиться к полю Name структуры
// Human, необходимо обращаться так Human.Name если мы обратимся просто к полю Name, тогда
// мы получим значение поля Name структуры Action. Тоже самое относиться и к методам.
type Action struct {
	ID   int64
	Name string
	Human
}

func (a Action) PrintHuman() string {
	return fmt.Sprintf("ID - %d Вид спорта - %s\nПобедитель: %s\n", a.ID, a.Name, a.Human.PrintHuman())
}

func main() {
	// Пример: нужно вывести в консоль информацию о победителе соревнований по разным видам спорта.
	win := Action{
		ID: 1,
		Name: "Плавание",
		Human: Human{
			ID:         5,
			Name:       "Иван",
			MiddleName: "Иваныч",
			LastName:   "Иванов",
			Age:        20,
			Hobby:      []string{"бегать по утрам", "кататься на лыжах", "читать"},
		},
	}

	// Выведет в STDOUT результат работы метода структуры Action.
	fmt.Println(win.PrintHuman())
		// ID - 1 Вид спорта - Плавание
		// Победитель: ID 5
		// Фамилия: Иванов
		// Имя: Иван
		// Отчество: Иваныч
		// Возраст: 20
		// Хобби: [бегать по утрам кататься на лыжах читать]

	// Если нужно вывести в консоль только структуру Human, тогда
	// необходимо обратиться к полю Human и затем к нужному методу.
	fmt.Println(win.Human.PrintHuman())
		// ID 5
		// Фамилия: Иванов
		// Имя: Иван
		// Отчество: Иваныч
		// Возраст: 20
		// Хобби: [бегать по утрам кататься на лыжах читать]
}
