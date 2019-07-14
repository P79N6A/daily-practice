package services

import (
	"context"
	"errors"
	"log"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services/thrift_gen/services"
)

func convertProfileThrift(profile *database.Profile) *services.Profile {
	result := services.Profile{
		ID:          int32(profile.ID),
		Description: profile.Description,
		Avatar:      profile.Avatar,
	}
	return &result
}

func convertUserThrift(user *database.User) *services.UserResponse {
	result := services.UserResponse{
		ID:       int32(user.ID),
		Name:     user.Name,
		Password: user.Password,
		Status:   int32(user.Status),
		Profile:  convertProfileThrift(&user.Profile),
	}
	return &result
}

// impl UserServiceServer
type UserServiceHandler struct{}

func (*UserServiceHandler) GetUser(ctx context.Context, req *services.UserRequest) (r *services.UserResponse, err error) {
	username := req.GetName()
	log.Printf("[THRIFT] received a request to /GetUserGRPC(%s)", username)
	success, user := database.RetrieveUserByName(username)

	if !success {
		return nil, errors.New("user not exists")
	}

	return convertUserThrift(user), nil
}
