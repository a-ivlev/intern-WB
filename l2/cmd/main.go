package main

import (
	"encoding/xml"
	"fmt"
)

// var (
// 	bank = pattern.Bank{
// 		Name: "Alfa",
// 		Cards: []pattern.Card{},
// 	}

// 	card1 = pattern.Card{
// 		Number: "123456789",
// 		Balance: 200,
// 		Bank: &bank,
// 	}

// 	card2 = pattern.Card{
// 		Number: "987654321",
// 		Balance: 50,
// 		Bank: &bank,
// 	}

// 	user = pattern.User{
// 		Name: "Покупатель-1",
// 		Card: &card1,
// 	}

// 	user2 = pattern.User{
// 		Name: "Покупатель-2",
// 		Card: &card2,
// 	}

// 	prod = pattern.Product{
// 		Name: "Товар-1",
// 		Price: 100,
// 	}

// 	shop = pattern.Shop{
// 		Name: "Shop-1",
// 		Products: []pattern.Product{
// 			prod,
// 		}
// 	}
// )

// func main() {
// 	//println(concat2codegen("Hello ", "word!"))
// 	//fmt.Fprintf(&buf, "%d:%d, ", i+1, p)

// 	s1 := "Hello"
// 	s2 := "word!"

// 	var bstr strings.Builder
// 	bstr.Grow(len(s1) + len(s2))
// 	fmt.Fprintf(&bstr, "%s %s\n", s1, s2)
// 	print(bstr.String())
// }

// func concat2codegen(s1, s2 string) string { return s1 + s2 }

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(wg sync.WaitGroup, i int) {
// 			fmt.Println(i)
// 			wg.Done()
// 		}(wg, i)
// 	}
// 	wg.Wait()
// 	fmt.Println("exit")
// }

// Visitor описывает общий интерфейс для всех типов посетителей. Он объявляет набор методов,
// отличающихся типом входящего параметра, которые нужны для запуска операции для всех типов
// конкретных элементов.
type Visitor interface {
	VisitForDot(*Dot) string
	VisitForCircle(*Circle) string
	VisitForRectangle(*Rectangle) string
}

// Элемент описывает метод принятия посетителя. Этот метод должен иметь единственный
// параметр, объявленный с типом интерфейса посетителя.
type Element interface {
	Accept(v Visitor) string
}

// Конкретные посетители реализуют какое-то осбенное поведение для всех типов элементов,
// которые можно подать через методы интерфейса посетителя.
// People implements the Visitor interface.
type XMLExportVisitor struct {
	xml string
}

// VisitSushiBar implements visit to SushiBar.
func (v *XMLExportVisitor) VisitForDot(s *Dot) string {
	xmlDot, err := xml.Marshal(s)
	if err != nil {
		return ""
	}
	return string(xmlDot)
}

// VisitPizzeria implements visit to Pizzeria.
func (v *XMLExportVisitor) VisitForCircle(s *Circle) string {
	xmlCircle, err := xml.Marshal(s)
	if err != nil {
		return ""
	}
	return string(xmlCircle)
}

// VisitBurgerBar implements visit to BurgerBar.
func (v *XMLExportVisitor) VisitForRectangle(s *Rectangle) string {
	xmlRectangle, err := xml.Marshal(s)
	if err != nil {
		return ""
	}
	return string(xmlRectangle)
}

type Shape interface {
	getType()
	Accept(v Visitor) string
}

// SushiBar implements the Place interface.
type Dot struct {
	X int64
	Y int64
}

// BuySushi implementation.
func (d *Dot) getType() string {
	return "Dot"
}

// Accept implementation.
func (d *Dot) Accept(v Visitor) string {
	return v.VisitForDot(d)
}

// Pizzeria implements the Place interface.
type Circle struct {
	X int64
	Y int64
	Radius int64 
}

// BuyPizza implementation.
func (c *Circle) getType() string {
	return "Circle"
}

// Accept implementation.
func (c *Circle) Accept(v Visitor) string {
	return v.VisitForCircle(c)
}

// BurgerBar implements the Place interface.
type Rectangle struct {
	SideLeft int64
	SideRight int64
}

// Shape implementation.
func (r *Rectangle) getType() string {
	return "Rectangle"
}

// Accept implementation.
func (r *Rectangle) Accept(v Visitor) string {
	return v.VisitForRectangle(r)
}

func main() {
	dot := &Dot{1,2}
	circle := &Circle{1,2,3}
	rectangle := Rectangle{5,10}

	xmlVisitor := &XMLExportVisitor{}
	fmt.Println(dot.Accept(xmlVisitor))
	fmt.Println(circle.Accept(xmlVisitor))
	fmt.Println(rectangle.Accept(xmlVisitor))
}