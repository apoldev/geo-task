package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/GoGerman/geo-task/geo"
	"github.com/GoGerman/geo-task/module/courier/models"
	"github.com/GoGerman/geo-task/module/courier/storage"
	"math"
)

// Направления движения курьера
const (
	DirectionUp    = 0
	DirectionDown  = 1
	DirectionLeft  = 2
	DirectionRight = 3
)

const (
	DefaultCourierLat = 59.9311
	DefaultCourierLng = 30.3609
)

type Courierer interface {
	GetCourier(ctx context.Context) (*models.Courier, error)
	MoveCourier(courier models.Courier, direction, zoom int) error
}

type CourierService struct {
	courierStorage storage.CourierStorager
	allowedZone    geo.PolygonChecker
	disabledZones  []geo.PolygonChecker
}

func NewCourierService(courierStorage storage.CourierStorager, allowedZone geo.PolygonChecker, disbledZones []geo.PolygonChecker) Courierer {
	return &CourierService{courierStorage: courierStorage, allowedZone: allowedZone, disabledZones: disbledZones}
}

func (c *CourierService) checkCoureier() {

}
func (c *CourierService) GetCourier(ctx context.Context) (*models.Courier, error) {
	var courier *models.Courier
	var err error

	// получаем курьера из хранилища используя метод GetOne из storage/courier.go
	courier, err = c.courierStorage.GetOne(ctx)

	if err != nil {
		return nil, err
	}

	if courier == nil {
		courier = &models.Courier{
			Location: models.Point{
				Lat: DefaultCourierLat,
				Lng: DefaultCourierLng,
			},
		}
	}

	// проверяем, что курьер находится в разрешенной зоне
	// если нет, то перемещаем его в случайную точку в разрешенной зоне
	// сохраняем новые координаты курьера
	if !geo.CheckPointIsAllowed(geo.Point{
		courier.Location.Lat,
		courier.Location.Lng,
	}, c.allowedZone, c.disabledZones) {

		fmt.Println("not allowed", courier.Location.Lat, courier.Location.Lng)
		rp := geo.GetRandomAllowedLocation(c.allowedZone, c.disabledZones)

		courier.Location = models.Point{
			Lat: rp.Lat,
			Lng: rp.Lng,
		}

	}

	c.courierStorage.Save(ctx, *courier)

	return courier, nil
}

// MoveCourier : direction - направление движения курьера, zoom - зум карты
func (c *CourierService) MoveCourier(courier models.Courier, direction, zoom int) error {

	var err error

	// точность перемещения зависит от зума карты использовать формулу 0.001 / 2^(zoom - 14)
	// 14 - это максимальный зум карты
	if zoom > 14 {
		zoom = 14
	}

	d := 0.001 / math.Pow(2, float64(zoom-14))

	switch direction {
	case DirectionUp:
		courier.Location.Lat += d
	case DirectionDown:
		courier.Location.Lat -= d
	case DirectionLeft:
		courier.Location.Lng -= d
	case DirectionRight:
		courier.Location.Lng += d
	default:
		return errors.New("incorrect direction")
	}

	// далее нужно проверить, что курьер не вышел за границы зоны
	// если вышел, то нужно переместить его в случайную точку внутри зоны
	if !geo.CheckPointIsAllowed(geo.Point{
		courier.Location.Lat,
		courier.Location.Lng,
	}, c.allowedZone, c.disabledZones) {

		rp := geo.GetRandomAllowedLocation(c.allowedZone, c.disabledZones)

		courier.Location = models.Point{
			Lat: rp.Lat,
			Lng: rp.Lng,
		}
	}

	fmt.Println("move", courier, d)

	// далее сохранить изменения в хранилище
	ctx := context.Background()
	err = c.courierStorage.Save(ctx, courier)
	if err != nil {
		return err
	}

	return nil
}
