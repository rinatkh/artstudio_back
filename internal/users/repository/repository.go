package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/users"
	"github.com/rinatkh/artstudio_back/internal/users/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) users.UserRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetUserById(id string) (*core.User, error) {
	var data []core.User
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.users WHERE uuid='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrUserDBNotFound
	}
	return &data[0], nil
}

func (p postgresRepository) GetUsers(limit, offset int64) (*[]core.User, int64, error) {
	var data []core.User
	queryStr := "SELECT * FROM public.users WHERE uuid <> ''"
	if limit == 0 {
		queryStr += " LIMIT 1"
	} else {
		queryStr += fmt.Sprintf(" LIMIT %d", limit)
	}
	queryStr += fmt.Sprintf(" OFFSET %d", offset)

	err := p.db.Select(
		&data, queryStr)

	if err != nil {
		return nil, 0, err
	}
	if len(data) == 0 {
		return nil, 0, nil
	}
	var length []int64
	err = p.db.Select(&length, "SELECT count(*) FROM public.users")
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateUser(user *core.User) (*core.User, error) {
	res, err := p.db.Query("INSERT INTO public.users (firstname, surname, middlename, sex, birth_date, role, image, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", user.Firstname, user.Surname, user.Middlename, user.Sex, user.BirthDate, user.Role, user.Image, user.CreateAt)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getUser(user)
}

func (p postgresRepository) UpdateUser(user *core.User) (*core.User, error) {
	query := fmt.Sprintf("UPDATE users SET firstname='%s', surname='%s', middlename='%s', sex='%s', birth_date=$1, role = '%s', image = '%s' where uuid='%s'", user.Firstname, user.Surname, user.Middlename, user.Sex, user.Role, user.Image, user.Id)
	res, err := p.db.Query(query, user.BirthDate)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getUser(user)
}

func (p postgresRepository) DeleteUser(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM public.users WHERE uuid='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getUser(user *core.User) (*core.User, error) {
	var data []core.User

	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.users WHERE firstname='%s' AND surname='%s' AND middlename='%s' AND role='%s' AND image='%s' AND sex='%s'", user.Firstname, user.Surname, user.Middlename, user.Role, user.Image, user.Sex))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
