package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"intern-WB/l0/backend/internal/api/nats"
	"intern-WB/l0/backend/internal/api/router"
	"intern-WB/l0/backend/internal/api/server"

	"intern-WB/l0/backend/internal/app/repository"
	"intern-WB/l0/backend/internal/store"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if envVal, ok := os.LookupEnv(key); ok {
		envInt64, err := strconv.ParseInt(envVal, 10, 64)
		if err == nil {
			return envInt64
		}
	}
	return defaultVal
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	var (
	postgresDSN = flag.String("pg-dsn", getEnv("NATS_POSTGRES_DSN", "postgres://admin:password@localhost/probe-db"), "DSN of the Postgres instance")
	webServerPORT = flag.Int64("srv-port", getEnvInt64("WEB_SERVER_PORT", 8080), "Port for the web server")
	natsClusterID = flag.String("nats-cluster-id", getEnv("NATS_CLUSTER_ID", "test-cluster"), "NATS cluster ID")
	natsClientID = flag.String("nats-client-id", getEnv("NATS_CLIENT_ID", "client-123"), "NATS client ID")
	natsDurableName = flag.String("nats-durable-name", getEnv("NATS_Durable_Name", "my-durable"), "NATS durable name to a subscription when it is created. Doing this causes the NATS Streaming server to track the last acknowledged message for that clientID + durable name, so that only messages since the last acknowledged message will be delivered to the client.")
	natsChanelName = flag.String("nats-ch-name", getEnv("NATS_CHANEL_NAME", "orders"), "The name of the channel to which the nats client subscribes.")
	)
	flag.Parse()
	

	store, err := store.NewStore(ctx, *postgresDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	repo := repository.NewOrdersRepo(store)
	r := router.NewRouter(repo)
	srv := server.NewServer(fmt.Sprintf(":%d", *webServerPORT), r)

	natsConf := nats.NATSconf{
		ClusterID: *natsClusterID,
		ClientID: *natsClientID,
		DurableName: *natsDurableName,
		ChanelName: *natsChanelName,
	}

	nats.ListenNats(ctx, repo, natsConf)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	srv.Start(ctx, wg)

	<-ctx.Done()
	cancel()
	wg.Wait()
	fmt.Println("Server stopped")
}

// func (h *Handler) ListenChannelPurchases(m *stan.Msg) {
// func ListenChannelPurchases(m *stan.Msg) {
// 	var order model.Order
// 	if err := json.Unmarshal(m.Data, &order); err != nil {
// 		// logrus.Error("cannot parse purchase, drop message")
// 		return
// 	}
// 	fmt.Println("Listen channel order", order)

// 	// if purchase.PurchaseUID != purchase.Payment_.Transaction {
// 	// 	// log.Error("purchase.PurchaseUID != purchase.Payment_.Transaction, drop message")
// 	// 	return
// 	// }
// 	// // logrus.Debugf("get new purchase from channel %v", purchase)
// 	// err := h.repository.AddNewPurchase(purchase)
// 	// if err != nil {
// 	// 	// log.Errorf("error %s while adding in db, drop message", err.Error())
// 	// 	return
// 	// }

// 	// err = h.repository.AddNewPurchaseCache(purchase)
// 	// if err != nil {
// 	// 	// log.Error(err.Error())
// 	// } else {
// 	// 	// log.Infof("added new purchase to cache with uid %s", purchase.PurchaseUID)
// 	// }
// }
