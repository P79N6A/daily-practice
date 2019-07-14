package services

import (
	"context"
	"errors"
	"log"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/grpc_gen"
)

func convertProfileGRPC(profile *database.Profile) *services.UserResponse_Profile {
	result := services.UserResponse_Profile{
		ID:          uint32(profile.ID),
		Description: profile.Description,
		Avatar:      profile.Avatar,
	}
	return &result
}

func convertUserGRPC(user *database.User) *services.UserResponse {
	result := services.UserResponse{
		ID:       uint32(user.ID),
		Name:     user.Name,
		Password: user.Password,
		Status:   int32(user.Status),
		Profile:  convertProfileGRPC(&user.Profile),
	}
	return &result
}

// impl UserServiceServer
type UserServer struct{}

func (*UserServer) GetUser(ctx context.Context, request *services.UserRequest) (*services.UserResponse, error) {
	username := request.Name
	log.Printf("[GRPC] received a request to /GetUserGRPC(%s)", username)
	success, user := database.RetrieveUserByName(username)

	if !success {
		return nil, errors.New("user not exists")
	}

	return convertUserGRPC(user), nil
}
