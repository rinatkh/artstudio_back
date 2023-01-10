package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/departments"
	"github.com/rinatkh/artstudio_back/internal/departments/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) departments.DepartmentRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetDepartmentById(id int64) (*core.Department, error) {
	var data []core.Department
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM departments WHERE id='%d'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetDepartments(limit, offset int64) (*[]core.Department, int64, error) {
	var data []core.Department
	queryStr := "SELECT * FROM departments WHERE id > 0"
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
	err = p.db.Select(&length, "SELECT count(*) FROM Departments")
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateDepartment(department *core.Department) (*core.Department, error) {
	res, err := p.db.Query("INSERT INTO departments (name, image, description) VALUES ($1, $2, $3)", department.Name, department.Image, department.Description)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getDepartment(department)
}

func (p postgresRepository) UpdateDepartment(department *core.Department) (*core.Department, error) {

	query := fmt.Sprintf("UPDATE departments SET name='%s', image='%s', description='%s' where id='%d'", department.Name, department.Image, department.Description, department.Id)

	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getDepartment(department)
}

func (p postgresRepository) DeleteDepartment(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM departments WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getDepartment(department *core.Department) (*core.Department, error) {
	var data []core.Department
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM departments WHERE name='%s' AND image='%s' and description='%s'", department.Name, department.Image, department.Description))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
