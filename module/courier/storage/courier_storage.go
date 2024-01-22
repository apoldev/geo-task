package storage

import (
	"context"
	"github.com/GoGerman/geo-task/module/courier/models"
	"github.com/go-redis/redis"
)

type CourierStorager interface {
	Save(ctx context.Context, courier models.Courier) error // сохранить курьера по ключу courier
	GetOne(ctx context.Context) (*models.Courier, error)    // получить курьера по ключу courier
}

type CourierStorage struct {
	storage *redis.Client
}

func NewCourierStorage(storage *redis.Client) CourierStorager {
	return &CourierStorage{storage: storage}
}
