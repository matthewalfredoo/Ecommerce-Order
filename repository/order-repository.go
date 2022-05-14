package repository

import (
	"Ecommerce-Order/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrders() []model.Order
	GetOrder(id int) model.Order
	CreateOrder(order model.Order) model.Order
	UpdateOrder(id int, order model.Order) model.Order
	CancelOrder(id int) model.Order
}

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) OrderRepository {
	return &orderRepository{
		connection: conn,
	}
}

func (repository *orderRepository) GetOrders() []model.Order {
	var orders []model.Order
	repository.connection.Find(&orders)

	return orders
}

func (repository *orderRepository) GetOrder(id int) model.Order {
	var order model.Order
	repository.connection.First(&order, id)
	return order
}

func (repository *orderRepository) CreateOrder(order model.Order) model.Order {
	repository.connection.Create(&order)
	return order
}

func (repository *orderRepository) UpdateOrder(id int, order model.Order) model.Order {
	var orderUpdate model.Order // data of order that will be updated
	repository.connection.First(&orderUpdate, id)

	if orderUpdate.Status != order.Status {
		orderUpdate.Status = order.Status
		repository.connection.Save(&orderUpdate)
	}

	repository.connection.Save(&orderUpdate)
	return orderUpdate
}

func (repository *orderRepository) CancelOrder(id int) model.Order {
	var order model.Order
	repository.connection.First(&order, id)
	order.Status = "Dibatalkan"
	repository.connection.Save(&order)

	return order
}
