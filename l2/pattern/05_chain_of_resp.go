package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Цепочка вызовов - это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли
	передавать запрос дальше по цепи.

	Применимость:
	1. Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
	какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	С помощью "Цепочки вызовов" вы можете связать потенциальных обработчиков в одну цепь и при получении запроса
	поочерёдно спрашивать каждого из них, не хочет ли он обработать запрос.

	2. Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	"Цепочки вызовов" позволяет запускать обработчиков последовательно один за другим в том порядке, в котором они
	находятся в цепочке.

	3. Когда набор объектов, способных обработать запрос, должен задаваться динамически.
	В любой момент вы можете вмешаться в существующую цепочку и переназначить связи так, чтобы убрать или добавить
	новое звено.
*/

// Пациент пришедший за медицинской помощью.
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Интерфейс обработчика
type department interface {
	execute(*patient)
	setNext(department)
}

// Конкретный обработчик приёмное отделение.
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Printf("Patient %s registration already done", p.name)
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// Конкретный обработчик доктор.
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// Конкретный обработчик аптека.
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// Конкретный обработчик касса.
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
		c.next.execute(p)
		return
	}
	fmt.Println("Cashier getting money from patient")
	p.paymentDone = true
	c.next.execute(p)
}

func (c *cashier) setNext(next department) {
	c.next = next
}

//func main() {
//	// Цепочку событий задаём начиная с последнего элемента.
//
//	// Посешение кассы для оплаты за оказанные услуги.
//	cashier := &cashier{}
//
//	// После доктора пациент приходит в аптеку за лекарствами.
//	medical := &medical{}
//	medical.setNext(cashier)
//
//	// Из приёмного отделения пациент направляется к доктору.
//	doctor := &doctor{}
//	doctor.setNext(medical)
//
//	// Приёмное отделение
//	reception := &reception{}
//	reception.setNext(doctor)
//
//	// Создаём пациента клиники.
//	patient := &patient{name: "Vasily"}
//
//	// Визит пациента в клинику.
//	reception.execute(patient)
//}
