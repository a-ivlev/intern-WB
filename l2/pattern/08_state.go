package pattern

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Состояние - это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от
	своего состояния. Из вне создаётся впечатление, что изменился целиком весь объект.

	Основная идея в том, что программа	может находиться в нескольких состояниях, которые сменяют друг друга. Набор этих
	состояний, а также переходов между ними, предопределён и конечен. Находясь в разных состояниях, программа может
	по-разному реагировать на одни и те же события, которые	происходят с ней.
	Такой подход можно применять и к отдельным объектам, например объект документ, может принимать	три состояния:
	черновик, модерация или опубликован. В каждом из этих состояний, метод опубликовать будет работать по-разному:
	- из черновика он отправит документ на модерацию,
	- из модерации в публикацию, но при условии, что это сделал	администратор,
	- в опубликованном состоянии метод не будет делать ничего.
	Машину состояний чаще всего реализуют с помощью множества условных операторов if либо switch, которые проверяют
	текущее состояние объекта и выполняют соответствующее поведение.
	Паттерн состояние, предлагает создать отдельные объекты, для каждого состояния, в котором может прибывать объект, а
	затем внести туда поведение соответствующее этим состояниям. В место того, чтобы хранить код всех состояний,
	первоначальный объект называемый контекстом, будет содержать ссылку на один из объектов состояний и делегировать ему
	работу, зависящую от состояния. Благодаря тому, что объект состояний будет иметь общий интерфейс, контекст сможет
	делегировать работу, не привязываясь к его классу.
	Когда код объекта содержит множество больших похожих друг на друга условных операторов, которые выбирают поведение в
	зависимости от текущих значений полей этого объекта, паттерн предлагает переместить каждую ветку такого условного
	оператора в собственный объект, тип, тут можно поселить и все поля связанные с этим состоянием. Когда вы сосзнательно
	используете табличную модель состояния построенную на условных операторах, но вынуждены мирится с дублированием кода
	для похохожих состояний и переходов, паттерн состояние позволяет реализовать иерархическую машину состояний
	базирующуюся на наследовании, вы можете наследовать похожие состояния от одного родительского объекта и вынести туда
	весь дублирующийся код.

	Преймущества:
	1. Избавляет от множества больших условных операторов состояний.
	2. Концентрирует в одном месте код связанный с определённым состоянием.
	3. Упрошает код контекса.

	Недостатки:
	1. Может не оправдано усложнить код, если состояний мало и они редко меняются.
*/

/*
	Давайте применим паттерн проектирования "Состояние" в контексте торговых автоматов. Автомат может выдавать, только
	один товар и может пребывать только в одном из четырёх состояний:
	1. hasItem (имеет товар)
	2. noItem (не имеет товар)
	3. itemRequested (выдаёт товар)
	4. hasMoney (получил деньги)

	Торговый автомат может иметь различные действия:
	1. Выбрать товар
	2. Оплатить товар
	3. Выдать товар

	В нашем примере, автомат может быть в одном из множества состояний, которые непрерывно меняются. Предположим, что
	автомат находиться в режиме itemRequested (Выдача товара). Как только произойдёт действие "Оплата товара", он сразу
	же перейдёт в состояние hasMoney.
*/

var (
	ErrItemOut       = errors.New("item out of stock")
	ErrNoItemPresent = errors.New("no item present")
	ErrSelectItem    = errors.New("please select item first")
	ErrItemRequest   = errors.New("item already request")
	ErrInsertMoney   = errors.New("inserted money is less")
	ErrItemDispense  = errors.New("item dispense in progress")
)

// Интерфейс состояния
type state interface {
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
	addItem(count int) error
}

// Контекст
type vendingMachine struct {
	hasItem       state
	noItem        state
	itemRequested state
	hasMoney      state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount int, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &hasItemState{
		vendingMachine: v,
	}

	noItemState := &noItemState{
		vendingMachine: v,
	}

	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}

	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.noItem = noItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState

	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

// Конкретное состояние "товара нет".
type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
	return ErrItemOut
}

func (i *noItemState) insertMoney(money int) error {
	return ErrItemOut
}

func (i *noItemState) dispenseItem() error {
	return ErrItemOut
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

// Конкретное состояние "товар есть".
type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return ErrNoItemPresent
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) insertMoney(money int) error {
	return ErrSelectItem
}

func (i *hasItemState) dispenseItem() error {
	return ErrSelectItem
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

// Конкретное состояние "выбор товара".
type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return ErrItemRequest
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("%s. Please insert %d", ErrInsertMoney, i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return ErrItemRequest
}

func (i *itemRequestedState) addItem(count int) error {
	return ErrItemDispense
}

// Конкретное состояние "оплата товара".
type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() error {
	return ErrItemDispense
}

func (i *hasMoneyState) insertMoney(money int) error {
	return ErrItemOut
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func (i *hasMoneyState) addItem(count int) error {
	return ErrItemDispense
}

func main() {
	vendingMachine := newVendingMachine(1, 10)

	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err := vendingMachine.addItem(2); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}
