package service

import (
	"context"
	"github.com/GoGerman/geo-task/module/courier/models"
	cservice "github.com/GoGerman/geo-task/module/courier/service"
	cfm "github.com/GoGerman/geo-task/module/courierfacade/models"
	om "github.com/GoGerman/geo-task/module/order/models"
	oservice "github.com/GoGerman/geo-task/module/order/service"
)

const (
	CourierVisibilityRadius = 2500 // 2500m
)

type CourierFacer interface {
	MoveCourier(ctx context.Context, direction, zoom int) // отвечает за движение курьера по карте direction - направление движения, zoom - уровень зума
	GetStatus(ctx context.Context) cfm.CourierStatus      // отвечает за получение статуса курьера и заказов вокруг него
}

// CourierFacade фасад для курьера и заказов вокруг него (для фронта)
type CourierFacade struct {
	courierService cservice.Courierer
	orderService   oservice.Orderer
}

func NewCourierFacade(courierService cservice.Courierer, orderService oservice.Orderer) CourierFacer {
	return &CourierFacade{courierService: courierService, orderService: orderService}
}

func (c *CourierFacade) MoveCourier(ctx context.Context, direction, zoom int) {
	var courier *models.Courier
	var err error

	courier, err = c.courierService.GetCourier(ctx)
	if err != nil {
		return
	}

	c.courierService.MoveCourier(*courier, direction, zoom)

}

func (c *CourierFacade) GetStatus(ctx context.Context) (res cfm.CourierStatus) {
	var courier *models.Courier
	var orders []om.Order
	var err error

	courier, err = c.courierService.GetCourier(ctx)

	if err != nil {
		return
	}

	orders, err = c.orderService.GetByRadius(
		ctx,
		courier.Location.Lng,
		courier.Location.Lat,
		CourierVisibilityRadius,
		"m",
	)

	return cfm.CourierStatus{
		Courier: *courier,
		Orders:  orders,
	}
}
