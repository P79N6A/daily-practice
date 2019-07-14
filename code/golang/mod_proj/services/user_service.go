package services

import (
	"context"
	"errors"
	"log"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/gen"
)

func convertProfile(profile *database.Profile) *services.UserResponse_Profile {
	result := services.UserResponse_Profile{
		ID:          uint32(profile.ID),
		Description: profile.Description,
		Avatar:      profile.Avatar,
	}
	return &result
}

func convertUser(user *database.User) *services.UserResponse {
	result := services.UserResponse{
		ID:       uint32(user.ID),
		Name:     user.Name,
		Password: user.Password,
		Status:   int32(user.Status),
		Profile:  convertProfile(&user.Profile),
	}
	return &result
}

// impl UserServiceServer
type UserServer struct{}

func (*UserServer) GetUser(ctx context.Context, request *services.UserRequest) (*services.UserResponse, error) {
	username := request.Name
	log.Printf("[GRPC] received a request to /GetUser(%s)", username)
	success, user := database.RetrieveUserByName(username)

	if !success {
		return nil, errors.New("user not exists")
	}

	return convertUser(user), nil
}
