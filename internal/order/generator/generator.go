package generator

import (
	"math/rand"
	"time"

	"github.com/dexxp/L0/internal/models"
)

func OrderGenerator() (*models.Order) {
	order := models.Order{
		OrderUid:          GenerateString(18),
		TrackNumber:       GenerateString(10),
		Entry:             GenerateString(10),
		Delivery:          DeliveryGenerator(),
		Payment:           PaymentGenerator(),
		Items:             ItemsGenerator(),
		Locale:            GenerateString(3),
		DateCreated:       time.Now(),
		OofShard:          GenerateString(10),
	}
	

	return &order
}

func DeliveryGenerator() (models.Delivery) {
	delivery := models.Delivery{
		Name    : GenerateString(10),
    Phone   : GenerateString(10),
    Zip     : GenerateString(10),
    City    : GenerateString(10),
    Address : GenerateString(10),
    Region  : GenerateString(10),
    Email   : GenerateString(10),
	}

	return delivery
}

func ItemsGenerator() ([]models.Item) {
	var (
		count = GenerateInt(5)()
		items = make([]models.Item, 0, count)
	)
	
	for i := 0; i < 5; i++ {
		item := ItemGenerator()
		items = append(items, item)
	}
	return items
}

func ItemGenerator() (models.Item) {
	generator10000 := GenerateInt(10000)

	item := models.Item{
		ChrtId     	: generator10000(),
    TrackNumber : GenerateString(10),
    Price      	: generator10000(),
    Rid        	: GenerateString(10),
    Name       	: GenerateString(10),
    Sale       	: generator10000(),
    Size       	: GenerateString(10),
    TotalPrice 	: generator10000(),
    NmId       	: generator10000(),
    Brand      	: GenerateString(10),
    Status     	: generator10000(),
	}

	return item
}

func PaymentGenerator() (models.Payment) {
	generate10000 := GenerateInt(10000)

	payment := models.Payment{
		Transaction   : GenerateString(10),
    RequestId     : GenerateString(10),
    Currency      : GenerateString(10),
    Provider      : GenerateString(10),
    Amount        : generate10000(),
    PaymentDt     : int64(generate10000()),
    Bank          : GenerateString(10),
    DeliveryCost  : generate10000(),
    GoodsTotal    : generate10000(),
    CustomFee     : generate10000(),
	}

	return payment
}

func GenerateString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateInt(max int) func() int {
	return func() int {
		return rand.Intn(max)
	}
}


