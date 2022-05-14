package service

import (
	"Ecommerce-Order/dto"
	"Ecommerce-Order/model"
	"Ecommerce-Order/repository"
	"github.com/mashingan/smapping"
	"log"
	"time"
)

type OrderService interface {
	GetOrders() []model.Order
	GetOrder(id int) model.Order
	CreateOrder(order dto.NewOrderDTO) model.Order
	UpdateOrder(id int, order dto.UpdateStatusOrderDTO) model.Order
	CancelOrder(id int) model.Order
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (os *orderService) GetOrders() []model.Order {
	return os.orderRepository.GetOrders()
}

func (os *orderService) GetOrder(id int) model.Order {
	return os.orderRepository.GetOrder(id)
}

func (os *orderService) CreateOrder(dto dto.NewOrderDTO) model.Order {
	orderDTOToModel := model.Order{}
	log.Println(dto)
	err := smapping.FillStruct(&orderDTOToModel, smapping.MapFields(&dto))
	if err != nil {
		log.Println(err)
		return model.Order{}
	}

	// Property that is not set by the user
	orderDTOToModel.Status = "Dipesan"
	orderDTOToModel.TotalHarga = uint64(dto.Harga * dto.JumlahProduct)
	orderDTOToModel.CreatedAt = time.Now()
	orderDTOToModel.UpdatedAt = time.Now()

	return os.orderRepository.CreateOrder(orderDTOToModel)
}

func (os *orderService) UpdateOrder(id int, order dto.UpdateStatusOrderDTO) model.Order {
	productDTOToModel := model.Order{}

	err := smapping.FillStruct(&productDTOToModel, smapping.MapFields(&order))
	if err != nil {
		return model.Order{}
	}

	productDTOToModel.UpdatedAt = time.Now()

	return os.orderRepository.UpdateOrder(id, productDTOToModel)
}

func (os *orderService) CancelOrder(id int) model.Order {
	return os.orderRepository.CancelOrder(id)
}
