package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/GoGerman/geo-task/module/courier/models"
	"github.com/redis/go-redis/v9"
)

const CourierKey = "courier"

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

func (s CourierStorage) GetOne(ctx context.Context) (*models.Courier, error) {
	var courier models.Courier
	var data []byte
	var err error

	data, err = s.storage.Get(ctx, CourierKey).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &courier)
	if err != nil {
		return nil, err
	}

	return &courier, nil
}

func (s CourierStorage) Save(ctx context.Context, courier models.Courier) error {

	_, err := s.storage.Set(ctx, CourierKey, courier, 0).Result()

	if err != nil {
		return err
	}

	return nil
}
