package pattern

// /*
// 	Реализовать паттерн «посетитель».
// Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
// 	https://en.wikipedia.org/wiki/Visitor_pattern
// */

// // Visitor описывает общий интерфейс для всех типов посетителей. Он объявляет набор методов,
// // отличающихся типом входящего параметра, которые нужны для запуска операции для всех типов
// // конкретных элементов.
// type Visitor interface {
// 	VisitSushiBar(p *SushiBar) string
// 	VisitPizzeria(p *Pizzeria) string
// 	VisitBurgerBar(p *BurgerBar) string
// }

// // Конкретные посетители реализуют какое-то осбенное поведение для всех типов элементов,
// // которые можно подать через методы интерфейса посетителя.
// // People implements the Visitor interface.
// type ConcreteVisitors struct {
// }

// // VisitSushiBar implements visit to SushiBar.
// func (v *ConcreteVisitors) VisitSushiBar(p *SushiBar) string {
// 	return p.BuySushi()
// }

// // VisitPizzeria implements visit to Pizzeria.
// func (v *ConcreteVisitors) VisitPizzeria(p *Pizzeria) string {
// 	return p.BuyPizza()
// }

// // VisitBurgerBar implements visit to BurgerBar.
// func (v *ConcreteVisitors) VisitBurgerBar(p *BurgerBar) string {
// 	return p.BuyBurger()
// }

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

// // Элемент описывает метод принятия посетителя. Этот метод должен иметь единственный
// // параметр, объявленный с типом интерфейса посетителя.
// type Element interface {
//	Buy() string
// 	Accept(v Visitor) string
// }

// // Конкретный элемент.
// // SushiBar implements the Place interface.
// type SushiBar struct {
// }

// // Buy implementation.
// func (s *SushiBar) Buy() string {
// 	return "Buy sushi..."
// }

// // Accept implementation.
// func (s *SushiBar) Accept(v Visitor) string {
// 	return v.VisitSushiBar(s)
// }

// // Конкретный элемент.
// // Pizzeria implements the Place interface.
// type Pizzeria struct {
// }

// // Buy implementation.
// func (p *Pizzeria) Buy() string {
// 	return "Buy pizza..."
// }

// // Accept implementation.
// func (p *Pizzeria) Accept(v Visitor) string {
// 	return v.VisitPizzeria(p)
// }

// // Конкретный элемент.
// // BurgerBar implements the Place interface.
// type BurgerBar struct {
// }

// // Buy implementation.
// func (b *BurgerBar) Buy() string {
// 	return "Buy burger..."
// }

// // Accept implementation.
// func (b *BurgerBar) Accept(v Visitor) string {
// 	return v.VisitBurgerBar(b)
// }
