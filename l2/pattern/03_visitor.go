package pattern

import "math"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Посетитель - это поведенческий паттерн проектирования, который позводяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Применяется:
	1. Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
	Посетитель позволяет применять одну и ту же операцию к объектам различный классов.

	2. Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции,
		но вы не хотите "засорять" классы такими операциями.
	Посетитель позволяет извлечь родственные операции из классов, составляющих структуру объектов, поместив их
	в один класс-посетитель.

	3. Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
	Посетитель позволяет определить поведение только для этих классов, оставив его пустым для всех остальных.
*/

// Shape интерфейс и его реализация, это стороняя библиотека, в которой мы не можем
// производить какие-либо изменения.
// Паттерн "Посетитель" позволяет нам реализовать дополнительный функционал, для этого
// нам достаточно добавить 1 метод accept(v Visitor) к каждой реализации интерфейса Shape.
type Shape interface {
	// GetShape один из методов сторонней библиотеки.
	GetShape() string
	// accept метод который необходимо добавить, для реализации
	// паттерна "Посетитель".
	accept(Visitor)
}

// Реализация интерфейса Shape фигурой "Квадрат".
type Square struct {
	Side int64
}

func (d *Square) GetShape() string {
	return "Shape is Square"
}

func (d *Square) accept(v Visitor) {
	v.VisitForSquare(d)
}

// Реализация интерфейса Shape фигурой "Круг".
type Circle struct {
	Radius int
}

func (c *Circle) GetShape() string {
	return "Shape is Circle"
}

func (c *Circle) accept(v Visitor) {
	v.VisitForCircle(c)
}

// Реализация интерфейса Shape фигурой "Прямоугольник".
type Rectangle struct {
	SideLeft  int
	SideRight int
}

func (r *Rectangle) GetShape() string {
	return "Shape is Rectangle"
}

func (r *Rectangle) accept(v Visitor) {
	v.VisitForRectangle(r)
}

// Visitor описывает общий интерфейс для всех типов посетителей.
// Он объявляет набор методов, отличающихся типом входящего параметра,
// которые нужны для запуска операции для всех типов конкретных элементов.
type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

// Element описывает метод принятия посетителя. Этот метод должен иметь
// единственный параметр, объявленный с типом интерфейса посетителя.
type Element interface {
	accept(v *Visitor)
}

// Конкретные посетители реализуют какое-то осбенное поведение для всех типов элементов,
// которые можно подать через методы интерфейса посетителя.
// В данном примере, считаем площадь фигур.
type AreaShape struct {
	area float64
}

// Конкретные элементы реализуют методы принятия посетителя. Цель этого метода -
// вызвать тот метод посещения, который соответствует типу этого элемента.

func (a *AreaShape) VisitForSquare(s *Square) {
	a.area = float64(s.Side * s.Side)
}

func (a *AreaShape) VisitForCircle(c *Circle) {
	a.area = math.Pi * float64(c.Radius*c.Radius)
}

func (a *AreaShape) VisitForRectangle(r *Rectangle) {
	a.area = float64(r.SideLeft * r.SideRight)
}

// func main() {
// 	square := &Square{Side: 3}
// 	circle := &Circle{Radius: 3}
// 	rectangle := &Rectangle{SideLeft: 2, SideRight: 3}

// 	areaSh := &AreaShape{}

// 	square.accept(areaSh)
// 	fmt.Println(areaSh.area)
// 	circle.accept(areaSh)
// 	fmt.Println(areaSh.area)
// 	rectangle.accept(areaSh)
// 	fmt.Println(areaSh.area)
// }
