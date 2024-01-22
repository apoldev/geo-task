package controller

import (
	"github.com/gin-gonic/gin"
)

type CourierController struct {
	courierService service.CourierFacer
}

func NewCourierController(courierService service.CourierFacer) *CourierController {
	return &CourierController{courierService: courierService}
}

func (c *CourierController) GetStatus(ctx *gin.Context) {
	// установить задержку в 50 миллисекунд

	// получить статус курьера из сервиса courierService используя метод GetStatus
	// отправить статус курьера в ответ
}

func (c *CourierController) MoveCourier(m webSocketMessage) {
	var cm CourierMove
	var err error
	// получить данные из m.Data и десериализовать их в структуру CourierMove

	// вызвать метод MoveCourier у courierService
}
