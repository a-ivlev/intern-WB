package pattern

import (
	"encoding/xml"
	"fmt"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Visitor описывает общий интерфейс для всех типов посетителей. Он объявляет набор методов,
// отличающихся типом входящего параметра, которые нужны для запуска операции для всех типов
// конкретных элементов.
type Visitor interface {
	VisitForDot(d *Dot) string
	VisitForCircle(c *Circle) string
	VisitForRectangle(r *Rectangle) string
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
func (v *XMLExportVisitor) VisitDot(d *Dot) string {
	xmlDot, err := xml.Marshal(d)
	if err != nil {
		return ""
	}
	return string(xmlDot)
}

// VisitPizzeria implements visit to Pizzeria.
func (v *XMLExportVisitor) VisitCircle(c *Circle) string {
	xmlCircle, err := xml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(xmlCircle)
}

// VisitBurgerBar implements visit to BurgerBar.
func (v *XMLExportVisitor) VisitRectangle(r *Rectangle) string {
	xmlRectangle, err := xml.Marshal(r)
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


// // Клиетном зачастую выступает коллекция или сложный составной объект,
// // например, дерево Компоновщика. Зачастую клиент не привязан к конкретным
// // классам элементов, работая с ними через общий интерфейс элементов.
// type Client struct {
// 	elements []Element
// }

// // Add appends Place to the collection.
// func (c *Client) Add(e Element) {
// 	c.elements = append(c.elements, e)
// }

// // Accept implements a visit to all places in the city.
// func (c *Client) Accept(v Visitor) string {
// 	var result string
// 	for _, p := range c.elements {
// 		result += p.Accept(v)
// 	}
// 	return result
// }