package auth

import "github.com/rinatkh/artstudio_back/internal/auth/models/dto"

type UseCase interface {
	SignupUser(params *dto.SignupUserRequest) (*dto.SignupUserResponse, error)
	LoginUser(params *dto.LoginUserRequest) (*dto.LoginUserResponse, error)
}
