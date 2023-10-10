package app

import (
	"fmt"
	"log"
	"time"

	"github.com/dexxp/L0/config"
	myNats "github.com/dexxp/L0/internal/nats"
	"github.com/dexxp/L0/internal/order/controller"
	"github.com/dexxp/L0/internal/order/generator"
	"github.com/dexxp/L0/internal/order/usecase"
	"github.com/dexxp/L0/internal/repository"
	server "github.com/dexxp/L0/pkg/httpserver"
	"github.com/dexxp/L0/pkg/postgres"
)



func Run(cfg *config.Config) {
	nts := myNats.NewNats(&cfg.Nats)

	// ------------

	pool, err := postgres.Connect(&cfg.PG)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// ------------
	
	postgresRepository := repository.NewPostgresRepository(pool)

	// ------------

	orderUC := usecase.NewOrderUseCase(postgresRepository)
	orderUC.LoadCache()

	// ------------

	go func() {
		for {
			fmt.Println("Send order")
			order := generator.OrderGenerator()

			err := nts.Publish(cfg.Nats.Topic, *order)
			if err != nil {
				fmt.Println("Cannot publish message: ", err)
				return
			}

			time.Sleep(time.Second * 7)
		}
	}()

	go func() {
		for {
			fmt.Println("Get order")
			order, err := nts.Subscribe(cfg.Nats.Topic)
			if err != nil {
				fmt.Println(err)
			}
			orderUC.CreateOrder(*order)
			fmt.Println("Save order to bd!")
			time.Sleep(time.Second * 6)
		}
	}()
	

	app := server.NewServer(&cfg.HTTP)

	orderHandler := controller.NewOrderHandler(orderUC)

	controller.GetOrderRoute(app.Fiber, orderHandler)	

	app.Run()
}