package auth

import "github.com/rinatkh/artstudio_back/internal/auth/models/core"

type AuthRepository interface {
	GetUserByEmail(email string) (*core.Auth, error)
	GetUserById(id string) (*core.Auth, error)
	DeleteUser(id string) error
	CreateUser(user *core.Auth) (*core.Auth, error)
}
