package services

import "github.com/prasetyodavid/go-stack/models"

type RabbitService interface {
	CreatePublisherRabbit(*models.CreateRabbitRequest) (*models.DBRabbit, error)
	CreateConsumerRabbit(*models.CreateRabbitRequest) (*models.DBRabbit, error)
}
