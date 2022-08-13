package gapi

import (
	"github.com/prasetyodavid/go-stack/config"
	"github.com/prasetyodavid/go-stack/pb"
	"github.com/prasetyodavid/go-stack/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config         config.Config
	userService    services.UserService
	userCollection *mongo.Collection
}

func NewGrpcUserServer(config config.Config, userService services.UserService, userCollection *mongo.Collection) (*UserServer, error) {
	userServer := &UserServer{
		config:         config,
		userService:    userService,
		userCollection: userCollection,
	}

	return userServer, nil
}
