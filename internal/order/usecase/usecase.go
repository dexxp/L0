package usecase

import (
	"fmt"
	"sync"

	"github.com/dexxp/L0/internal/models"
	"github.com/dexxp/L0/internal/repository"
)

type OrderUseCase struct {
	cacheOrder map[string]*models.Order
	repo *repository.PostgresRepository
	mu sync.Mutex //
}

func NewOrderUseCase(repo *repository.PostgresRepository) *OrderUseCase {
	return &OrderUseCase{
		cacheOrder: map[string]*models.Order{},
		repo: repo,
		mu: sync.Mutex{},
	}
}

func (o *OrderUseCase) CreateOrder(order models.Order) {
	o.repo.InsertOrder(order)

	o.mu.Lock()
	o.cacheOrder[order.OrderUid] = &order
	o.mu.Unlock()
}

func (o *OrderUseCase) LoadCache()  {
	orders, err := o.repo.GetAllOrders()

	if err != nil {
		fmt.Println("Error getting orders: ", err)
	}

	for _, order := range orders {
		o.cacheOrder[order.OrderUid] = &order
	}
}

func (o *OrderUseCase) GetOrderFromCacheByID(id string) *models.Order {
	return o.cacheOrder[id]
}

func (o *OrderUseCase) GetAllOrdersFromCache() map[string] *models.Order {
	return o.cacheOrder
}