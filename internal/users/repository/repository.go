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
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM users WHERE uuid='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetUsers(limit, offset int64) (*[]core.User, int64, error) {
	var data []core.User
	queryStr := "SELECT * FROM Users WHERE uuid <> ''"
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
	err = p.db.Select(&length, "SELECT count(*) FROM Users")
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateUser(user *core.User) (*core.User, error) {
	if *user.Middlename == "" {
		*user.Middlename = "NULL"
	}
	res, err := p.db.Query("INSERT INTO users (firstname, surname, middlename, sex, birth_date, role, image) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.Firstname, user.Surname, *user.Middlename, user.Sex, user.BirthDate, user.Role, user.Image)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getUser(user)
}

func (p postgresRepository) UpdateUser(user *core.User) (*core.User, error) {
	var query string
	if *user.Middlename == "" {
		query = fmt.Sprintf("UPDATE Users SET firstname='%s', surname='%s', middlename=NULL, sex='%s', fio = '%s', birth_date=$1, role = '%s', image = '%s' where uuid='%s'", user.Firstname, user.Surname, user.Sex, user.Firstname+" "+user.Surname, user.Id, user.Role, user.Image)
	} else {
		query = fmt.Sprintf("UPDATE Users SET firstname='%s', surname='%s', middlename='%s', sex='%s', fio = '%s', birth_date=$1, role = '%s', image = '%s' where uuid='%s'", user.Firstname, user.Surname, *user.Middlename, user.Sex, user.Firstname+" "+user.Surname+" "+*user.Middlename, user.Id, user.Role, user.Image)
	}
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
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM users WHERE uuid='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getUser(user *core.User) (*core.User, error) {
	var data []core.User
	var fio string
	if *user.Middlename == "NULL" {
		fio = user.Firstname + " " + user.Surname
	} else {
		fio = user.Firstname + " " + user.Surname + " " + *user.Middlename
	}
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM users WHERE fio='%s' AND sex='%s'", fio, user.Sex))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
