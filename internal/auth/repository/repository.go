package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/auth"
	"github.com/rinatkh/artstudio_back/internal/auth/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) auth.AuthRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetUserByEmail(email string) (*core.Auth, error) {
	var data []core.Auth
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.auth WHERE email='%s'", email))
	if err != nil {
		return nil, constants.ErrUserDBNotFound
	}
	if len(data) == 0 {
		return nil, constants.ErrUserDBNotFound
	}
	return &data[0], nil
}

func (p postgresRepository) GetUserById(id string) (*core.Auth, error) {
	var data []core.Auth
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.auth WHERE uuid='%s'", id))
	if err != nil {
		return nil, constants.ErrUserDBNotFound
	}
	if len(data) == 0 {
		return nil, constants.ErrUserDBNotFound
	}
	return &data[0], nil
}

func (p postgresRepository) UpdateUser(user *core.Auth) (*core.Auth, error) {
	query := fmt.Sprintf("UPDATE users SET email='%s', is_confirmed='%t' where uuid='%s'", user.Email, user.IsConfirmed, user.Id)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getUser(user)
}

func (p postgresRepository) CreateUser(user *core.Auth) (*core.Auth, error) {
	res, err := p.db.Query("INSERT INTO public.auth (uuid, email, password_hash, password_salt) VALUES ($1, $2, $3, $4)", user.Id, user.Email, user.UserPassword.Hash, user.UserPassword.Salt)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getUser(user)
}

func (p postgresRepository) DeleteUser(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM public.auth WHERE uuid='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}
func (p postgresRepository) getUser(user *core.Auth) (*core.Auth, error) {
	var data []core.Auth

	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.auth WHERE uuid='%s' AND email='%s' AND password_hash='%s' AND password_salt='%s'", user.Id, user.Email, user.UserPassword.Hash, user.UserPassword.Salt))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
