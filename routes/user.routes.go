package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prasetyodavid/go-stack/controllers"
	"github.com/prasetyodavid/go-stack/middleware"
	"github.com/prasetyodavid/go-stack/services"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
