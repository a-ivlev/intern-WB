package nats

import (
	"context"
	"encoding/json"
	"intern-WB/l0/backend/internal/app/model"
	"intern-WB/l0/backend/internal/app/repository"
	"log"

	"github.com/nats-io/stan.go"
)

type NATSconf struct {
	ClusterID string
	ClientID string
	DurableName string
	ChanelName string
}

func ListenNats(ctx context.Context, repo *repository.OrdersRepo, conf NATSconf) {
	sc, err := stan.Connect(conf.ClusterID, conf.ClientID)
	if err != nil {
		log.Printf("%s: %s", ErrConnNats, err.Error())
	}

	// Subscribe with durable name
	sc.Subscribe(conf.ChanelName, func(m *stan.Msg) {
		order := model.Order{}
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Printf("%s: %s", ErrIncorDataNats, err.Error())
		}

		_, err = repo.CreateOrder(ctx, order)
		if err != nil {
			log.Println(err)
		}
	}, stan.DeliverAllAvailable(), stan.DurableName(conf.DurableName))
}
