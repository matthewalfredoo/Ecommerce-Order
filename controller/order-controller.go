package controller

import (
	"Ecommerce-Order/dto"
	"Ecommerce-Order/helper"
	"Ecommerce-Order/model"
	"Ecommerce-Order/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController interface {
	GetOrders(context *gin.Context)
	GetOrder(context *gin.Context)
	CreateOrder(context *gin.Context)
	UpdateOrder(context *gin.Context)
	CancelOrder(context *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (controller *orderController) GetOrders(context *gin.Context) {
	orders := controller.orderService.GetOrders()
	res := helper.BuildResponse(true, "Orders retrieved successfully", orders)
	context.JSON(http.StatusOK, res)
}

func (controller *orderController) GetOrder(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		res := helper.BuildErrorResponse("Invalid order id", "Error", model.Order{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	order := controller.orderService.GetOrder(idInt)

	if order.ID == 0 {
		res := helper.BuildErrorResponse("Order not found", "Error", model.Order{})
		context.JSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildResponse(true, "Order retrieved successfully", order)
	context.JSON(http.StatusOK, res)
}

func (controller *orderController) CreateOrder(context *gin.Context) {
	var newOrderDTO dto.NewOrderDTO
	err := context.ShouldBind(&newOrderDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid order data", err.Error(), model.Order{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	order := controller.orderService.CreateOrder(newOrderDTO)

	if order.ID == 0 {
		res := helper.BuildErrorResponse("Error creating order", "Error", model.Order{})
		context.JSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Order created successfully", order)
	context.JSON(http.StatusOK, res)
}

func (controller *orderController) UpdateOrder(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)

	var updateOrderDTO dto.UpdateStatusOrderDTO
	err = context.ShouldBind(&updateOrderDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid order data", err.Error(), model.Order{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	updatedOrder := controller.orderService.UpdateOrder(idInt, updateOrderDTO)
	res := helper.BuildResponse(true, "Order updated successfully", updatedOrder)
	context.JSON(http.StatusOK, res)
}

func (controller *orderController) CancelOrder(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid order id", err.Error(), model.Order{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	cancelledOrder := controller.orderService.CancelOrder(idInt)
	res := helper.BuildResponse(true, "Order cancelled successfully", cancelledOrder)
	context.JSON(http.StatusOK, res)
}
