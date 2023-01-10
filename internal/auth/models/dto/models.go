package dto

import "github.com/rinatkh/artstudio_back/internal/users/models/dto"

type SignupUserRequest struct {
	dto.CreateUserRequest
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignupUserResponse struct {
	dto.User
	AuthToken string `json:"auth_token"`
}
type LoginUserResponse struct {
	dto.User
	AuthToken string `json:"auth_token"`
}
