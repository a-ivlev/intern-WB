package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Product struct {
	Name  string
	Price int64
}

type Shop struct {
	Name     string
	Products []Product
}

// Sell метод является реализацией паттерна "фасад". В нём происходит сложная бизнес-логика
// по проверке баланса на карте клиента, по проверке остатков товара.
func(s Shop) Sell(user User, product Product) error {
	fmt.Printf("[Магазин %s] запрос в банк об остатке средств на карте %s остаток: %d у клиента %s\n", s.Name, user.Card.Number, user.GetBalance, user.Name)

	for _, prod := range s.Products {
		if prod.Name == product.Name {
			if prod.Price > user.GetBalance() {
				return errors.New("у пользователя не достаточно средств на карте")
			}
			fmt.Printf("Товар %s приобретён!", product.Name)
			return nil
		}
	}
	return errors.New("запрашиваемый товар не найден")
}

type Bank struct {
	Name string
	Cards []Card
}

func(b Bank) CheckBalance(cardNumber string) (int64, error) {
	fmt.Printf("[Банк] Проверяет баланс по карте %s\n", cardNumber)
	for _, card := range b.Cards {
		if card.Number == cardNumber {
			return card.Balance, nil
		}
	}
	return 0, errors.New("нет такой карты")
}

type Card struct {
	Number    string
	Balance int64
	Bank    *Bank
}

func(c *Card) CheckBalance() (int64, error) {
	var err error
	c.Balance, err = c.Bank.CheckBalance(c.Number)
	if err != nil {
		return 0, err
	}
	return c.Balance, nil
}

type User struct {
	Name string
	Card *Card
}

func(u User) GetBalance() int64 {
	balance, err := u.Card.CheckBalance()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return balance
}

// type cpu struct{
//   void Freeze();
//   void Jump(long position);
//   void Execute();
// }
// func(c cpu) Freeze() {}
// func(c cpu) Jump(long position) {}
// func(c cpu) Execute() {}

// type hardDrive struct{
//   char* Read(long lba, int size);
// };

// type memory struct{
//   void load(long position, char* data);
// };

// func(m memory) load() {}

// type ComputerFacade struct {
//   cpu *cpu

// }

// func(cf ComputerFacade) Start() {
// 	cpu_.Freeze();
//     memory_.Load(kBootAddress, hard_drive_.Read(kBootSector, kSectorSize));
//     cpu_.Jump(kBootAddress);
//     cpu_.Execute();
// }

// func NewComputerFacade() *ComputerFacade {
// 	return &ComputerFacade{}
// }
