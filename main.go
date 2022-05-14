package main

import (
	"Ecommerce-Order/conn"
	"Ecommerce-Order/controller"
	"Ecommerce-Order/repository"
	"Ecommerce-Order/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)

var (
	db              *gorm.DB                   = conn.SetupDatabaseConnection()
	orderRepository repository.OrderRepository = repository.NewOrderRepository(db)
	orderService    service.OrderService       = service.NewOrderService(orderRepository)
	orderController controller.OrderController = controller.NewOrderController(orderService)
)

func main() {
	defer conn.CloseDatabaseConnection(db)
	router := gin.Default()

	routes := router.Group("/api/order")
	{
		routes.GET("/", orderController.GetOrders)
		routes.GET("/:id", orderController.GetOrder)
		routes.POST("/", orderController.CreateOrder)
		routes.PUT("/:id", orderController.UpdateOrder)
		routes.DELETE("/cancel/:id", orderController.CancelOrder)
	}

	err := router.Run(os.Getenv("BASE_URL_ORDER"))
	if err != nil {
		panic(err)
	}
}
