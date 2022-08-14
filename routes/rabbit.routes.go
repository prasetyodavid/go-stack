package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prasetyodavid/go-stack/controllers"
)

type RabbitRouteController struct {
	rabbitController controllers.RabbitController
}

func NewRabbitControllerRoute(rabbitController controllers.RabbitController) RabbitRouteController {
	return RabbitRouteController{rabbitController}
}

func (r *RabbitRouteController) RabbitRoute(rg *gin.RouterGroup) {
	router := rg.Group("/rabbits")
	router.POST("/publisher", r.rabbitController.CreatePublisherRabbit)
	router.POST("/consumer", r.rabbitController.CreateConsumerRabbit)
}
