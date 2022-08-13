package services

import "github.com/prasetyodavid/go-stack/models"

type AuthService interface {
	SignUpUser(*models.SignUpInput) (*models.UserDBResponse, error)
	SignInUser(*models.SignInInput) (*models.UserDBResponse, error)
}
