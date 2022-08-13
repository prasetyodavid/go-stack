package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prasetyodavid/go-stack/controllers"
	"github.com/prasetyodavid/go-stack/middleware"
	"github.com/prasetyodavid/go-stack/services"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), rc.authController.LogoutUser)
	router.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail)
	router.POST("/forgotPassword", rc.authController.ForgotPassword)
}
