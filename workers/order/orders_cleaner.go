package order

import (
	"context"
	"github.com/GoGerman/geo-task/module/order/service"
	"log"
	"time"
)

const (
	orderCleanInterval = 5 * time.Second
)

// OrderCleaner воркер, который удаляет старые заказы
// используя метод orderService.RemoveOldOrders()
type OrderCleaner struct {
	orderService service.Orderer
}

func NewOrderCleaner(orderService service.Orderer) *OrderCleaner {
	return &OrderCleaner{orderService: orderService}
}

func (o *OrderCleaner) orderRemover(ctx context.Context) {
	ticker := time.NewTicker(orderCleanInterval)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:

			err := o.orderService.RemoveOldOrders(ctx)

			if err != nil {
				log.Printf("error while removing old orders: %v", err)
			}
		}
	}
}
func (o *OrderCleaner) Run() {
	// исользовать горутину и select
	// внутри горутины нужно использовать time.NewTicker()
	// и вызывать метод orderService.RemoveOldOrders()
	// если при удалении заказов произошла ошибка, то нужно вывести ее в лог

	ctx := context.Background()
	go o.orderRemover(ctx)
}
