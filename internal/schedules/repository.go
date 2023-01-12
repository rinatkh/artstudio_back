package schedules

import "github.com/rinatkh/artstudio_back/internal/users/models/core"

type UserRepository interface {
	GetUserById(id string) (*core.User, error)
	GetUsers(limit, offset int64) (*[]core.User, int64, error)
	DeleteUser(id string) error
	CreateUser(user *core.User) (*core.User, error)
	UpdateUser(user *core.User) (*core.User, error)
}
