package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prasetyodavid/go-stack/models"
	"github.com/prasetyodavid/go-stack/services"
)

type RabbitController struct {
	rabbitService services.RabbitService
}

func NewRabbitController(rabbitService services.RabbitService) RabbitController {
	return RabbitController{rabbitService}
}

func (pc *RabbitController) CreateConsumerRabbit(ctx *gin.Context) {
	var rabbit *models.CreateRabbitRequest

	if err := ctx.ShouldBindJSON(&rabbit); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newRabbit, err := pc.rabbitService.CreateConsumerRabbit(rabbit)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newRabbit})
}

func (pc *RabbitController) CreatePublisherRabbit(ctx *gin.Context) {
	var rabbit *models.CreateRabbitRequest

	if err := ctx.ShouldBindJSON(&rabbit); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newRabbit, err := pc.rabbitService.CreatePublisherRabbit(rabbit)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newRabbit})
}
