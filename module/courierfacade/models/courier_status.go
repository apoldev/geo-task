package models

import (
	cm "github.com/GoGerman/geo-task/module/courier/models"
	om "github.com/GoGerman/geo-task/module/order/models"
)

type CourierStatus struct {
	Courier cm.Courier `json:"courier"`
	Orders  []om.Order `json:"orders"`
}
