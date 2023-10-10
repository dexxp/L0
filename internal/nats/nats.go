package nats

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dexxp/L0/config"
	"github.com/dexxp/L0/internal/models"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Nats struct {
	cfg *config.Nats
	sc stan.Conn	
	ns *nats.Conn
}

func NewNats(cfg *config.Nats) *Nats {
	url := fmt.Sprintf("nast://%s:%s", cfg.Host, cfg.Port)
	ns, err := nats.Connect(url)

	if err != nil {
		fmt.Println("Nats connect failed")
		return nil
	}

	sc, err := stan.Connect(cfg.Cluster, cfg.Client)
	if err != nil {
		fmt.Println("Stan connect failed")
		return nil
	}

	return &Nats{cfg: cfg, sc: sc, ns: ns}
}

func (nts *Nats) Publish(topic string, order models.Order) error {
	orderJSON, err := json.Marshal(order)
	
	if err != nil {
		return err
	}

	return nts.sc.Publish(topic, orderJSON)
}

func (nts *Nats) Subscribe(topic string) (*models.Order, error) {
	var receivedMessage models.Order

	ch := make(chan *models.Order)
	_, err := nts.sc.Subscribe(topic, func(msg *stan.Msg) {
		err := json.Unmarshal(msg.Data, &receivedMessage)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			return
		}
		ch <- &receivedMessage
  })

	if err != nil {
		return nil, err
	}

	select {
		case receivedMessage := <-ch:
			return receivedMessage, nil
		case <-time.After(60 * time.Second):
			return nil, stan.ErrTimeout
	}
}