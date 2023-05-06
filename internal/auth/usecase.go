package auth

import "github.com/rinatkh/artstudio_back/internal/auth/models/dto"

type UseCase interface {
	SignupUser(id string) (*dto.SignupUserResponse, error)
	SignupUserPre(params *dto.SignupUserRequest) (*dto.SignupPreResponse, error)
	LoginUser(params *dto.LoginUserRequest) (*dto.LoginUserResponse, error)
}
