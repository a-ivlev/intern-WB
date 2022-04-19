package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Entity struct {
	ID    string
	Order Order
}

// Order структура для входящих и исходящих данных.
type Order struct {
	OrderUID          string 	`json:"order_uid"`
	TrackNumber       string 	`json:"track_number"`
	Entry             string 	`json:"entry"`
	Delivery          		 	`json:"delivery"`
	Payment           		 	`json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int64     `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (o Order) Value() ([]byte, error) {
	return json.Marshal(o)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (o *Order) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &o)
}

// Delivery структура для входящих и исходящих данных.
type Delivery struct {
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Zip     string    `json:"zip"`
	City    string    `json:"city"`
	Address string    `json:"address"`
	Region  string    `json:"region"`
	Email   string    `json:"email"`
}

// Payment структура для входящих и исходящих данных.
type Payment struct {
	Transaction  string    `json:"transaction"`
	RequestID    string    `json:"request_id"`
	Currency     string    `json:"currency"`
	Provider     string    `json:"provider"`
	Amount       int64     `json:"amount"`
	PaymentDt    uint64    `json:"payment_dt"`
	Bank         string    `json:"bank" db:"bank"`
	DeliveryCost int64     `json:"delivery_cost"`
	GoodsTotal   int64     `json:"goods_total"`
	CustomFee    int64     `json:"custom_fee"`
}

// Item структура для входящих и исходящих данных.
type Item struct {
	ChrtID      int64     `json:"chrt_id"`
	TrackNumber string    `json:"track_number"`
	Price       int64     `json:"price"`
	Rid         string    `json:"rid"`
	Name        string    `json:"name"`
	Sale        int64     `json:"sale"`
	Size        string    `json:"size"`
	TotalPrice  int64     `json:"total_price"`
	NmID        int64     `json:"nm_id"`
	Brand       string    `json:"brand"`
	Status      int64     `json:"status"`
}
