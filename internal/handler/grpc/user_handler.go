package grpc

import (
	"context"

	authv1 "github.com/xreyc/grpc-auth/internal/gen/go/auth/v1"
)

type UserHandler struct {
	authv1.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUserDetails(ctx context.Context, req *authv1.GetUserRequest) (*authv1.GetUserResponse, error) {
	// Hardcoded response
	return &authv1.GetUserResponse{
		Username: req.GetUsername(),
		Email:    "xreyc@example.com",
		FullName: "Reyco Seguma",
	}, nil
}
