package controller

import (
	"context"
	"encoding/json"
	"github.com/GoGerman/geo-task/module/courierfacade/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type CourierController struct {
	courierService service.CourierFacer
}

func NewCourierController(courierService service.CourierFacer) *CourierController {
	return &CourierController{courierService: courierService}
}

func (c *CourierController) GetStatus(ctx *gin.Context) {
	// установить задержку в 50 миллисекунд
	time.Sleep(time.Millisecond * 50)

	// получить статус курьера из сервиса courierService используя метод GetStatus
	// отправить статус курьера в ответ
	ctx.JSON(200, c.courierService.GetStatus(ctx))
}

func (c *CourierController) MoveCourier(m webSocketMessage) {
	var cm CourierMove
	var err error
	// получить данные из m.Data и десериализовать их в структуру CourierMove

	if v, ok := m.Data.([]byte); ok {

		err = json.Unmarshal(v, &cm)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if cm.Direction == 0 && cm.Zoom == 0 {
		return
	}

	// вызвать метод MoveCourier у courierService
	ctx := context.Background()
	c.courierService.MoveCourier(ctx, cm.Direction, cm.Zoom)
}
