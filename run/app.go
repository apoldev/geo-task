package run

import (
	"context"
	"github.com/GoGerman/geo-task/cache"
	"github.com/GoGerman/geo-task/geo"
	cservice "github.com/GoGerman/geo-task/module/courier/service"
	storage2 "github.com/GoGerman/geo-task/module/courier/storage"
	"github.com/GoGerman/geo-task/module/courierfacade/controller"
	"github.com/GoGerman/geo-task/module/courierfacade/service"
	oservice "github.com/GoGerman/geo-task/module/order/service"
	"github.com/GoGerman/geo-task/module/order/storage"
	"github.com/GoGerman/geo-task/router"
	"github.com/GoGerman/geo-task/server"
	"github.com/GoGerman/geo-task/workers/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	// получение хоста и порта redis
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	// инициализация клиента redis
	rclient := cache.NewRedisClient(host, port)

	// инициализация контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// проверка доступности redis
	_, err := rclient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	// инициализация разрешенной зоны
	allowedZone := geo.NewAllowedZone()
	// инициализация запрещенных зон
	disAllowedZones := []geo.PolygonChecker{geo.NewDisAllowedZone1(), geo.NewDisAllowedZone2()}

	// инициализация хранилища заказов
	orderStorage := storage.NewOrderStorage(rclient)
	// инициализация сервиса заказов
	orderService := oservice.NewOrderService(orderStorage, allowedZone, disAllowedZones)

	orderGenerator := order.NewOrderGenerator(orderService)
	orderGenerator.Run()

	oldOrderCleaner := order.NewOrderCleaner(orderService)
	oldOrderCleaner.Run()

	// инициализация хранилища курьеров
	courierStorage := storage2.NewCourierStorage(rclient)
	// инициализация сервиса курьеров
	courierSevice := cservice.NewCourierService(courierStorage, allowedZone, disAllowedZones)

	// инициализация фасада сервиса курьеров
	courierFacade := service.NewCourierFacade(courierSevice, orderService)

	// инициализация контроллера курьеров
	courierController := controller.NewCourierController(courierFacade)

	// инициализация роутера
	routes := router.NewRouter(courierController)
	// инициализация сервера
	r := server.NewHTTPServer()
	// инициализация группы роутов
	api := r.Group("/api")
	// инициализация роутов
	routes.CourierAPI(api)

	mainRoute := r.Group("/")

	routes.Swagger(mainRoute)
	// инициализация статических файлов
	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("public"))))

	// запуск сервера
	//serverPort := os.Getenv("SERVER_PORT")

	if os.Getenv("ENV") == "prod" {
		certFile := "/app/certs/cert.pem"
		keyFile := "/app/certs/private.pem"
		return r.RunTLS(":443", certFile, keyFile)
	}

	return r.Run()
}
