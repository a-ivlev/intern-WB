package pattern

import "strings"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Builder — это паттерн проектирования, который позволяет поэтапно создавать сложные объекты
// с помощью четко определенной последовательности действий.

// Паттерн проектирования Builder решает такие проблемы, как:
// Класс (процесс строительства) создаёт различные представления сложного объекта.

// Builder (строитель) - абстрактный класс/интерфейс, который определяет все этапы, необходимые
// для производства сложного объекта-Страница. Как правило, здесь объявляются (абстрактно)
// все этапы (buildPart), а их реализация относится к классам конкретных строителей (ConcreteBuilder).
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
	GetPage() Page
}

// Director (распорядитель) - супервизионный класс, под конролем котрого строитель выполняет
// скоординированные этапы для создания объекта-Страница. Распорядитель обычно получает на вход
// строителя с этапами на выполнение в четком порядке для построения объекта-Страница.
type Director struct {
	builder Builder
}

// NewDirector метод конструктор создаёт нового Director (распорядитель).
func NewDirector(builder Builder) *Director {
	return &Director{
		builder,
	}
}

// CreatePage определяет чёткий порядок действий, по которому будет создаваться страница.
func (d *Director) CreatePage() Page {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
	return d.builder.GetPage()
}

// Page (Страница) - Класс, который определяет сложный объект, который мы пытаемся шаг
// за шагом сконструировать, используя простые объекты.
type Page struct {
	Content string
}

// ConcreteBuilder1 (конкретный строитель) класс-строитель, который предоставляет фактический код
// для создания объекта-продукта. У нас может быть несколько разных ConcreteBuilder-классов,
// каждый из которых реализует различную разновидность или способ создания объекта-Страница.
type ConcreteBuilder1 struct {
	page *Page
}

// NewconcreteBuilder1 функция конструктор создаёт concreteBuilder1.
func NewconcreteBuilder1() *ConcreteBuilder1 {
	return &ConcreteBuilder1{
		page: &Page{},
	}
}

// MakeHeader создаёт шапку страницы.
func (cb *ConcreteBuilder1) MakeHeader(header string) {
	var b strings.Builder
	b.Grow(17 + len(header))
	b.WriteString("<header>")
	b.WriteString(header)
	b.WriteString("</header>")
	cb.page.Content = b.String()
}

// MakeBody создаёт тело страницы.
func (cb *ConcreteBuilder1) MakeBody(body string) {
	var b strings.Builder
	b.Grow(13 + len(body) + len(cb.page.Content))
	b.WriteString(cb.page.Content)
	b.WriteString("<body>")
	b.WriteString(body)
	b.WriteString("</body>")
	cb.page.Content = b.String()
}

// MakeFooter создаёт заключительную часть страницы.
func (cb *ConcreteBuilder1) MakeFooter(footer string) {
	var b strings.Builder
	b.Grow(17 + len(footer) + len(cb.page.Content))
	b.WriteString(cb.page.Content)
	b.WriteString("<footer>")
	b.WriteString(footer)
	b.WriteString("</footer>")
	cb.page.Content = b.String()
}

// GetPage возвращает готовую страницу.
func (cb *ConcreteBuilder1) GetPage() Page {
	return *cb.page
}
