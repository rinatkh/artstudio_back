package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/DepartmentSubjects"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) departmentSubjects.DepartmentSubjectsRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetDepartmentSubjects(departmentId, limit, offset int64) (*[]departmentSubjects.DepartmentSubjects, int64, error) {
	var data []departmentSubjects.DepartmentSubjects
	queryStr := fmt.Sprintf("SELECT * FROM DepartmentSubjects WHERE department_id='%d'", departmentId)
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
	err = p.db.Select(&length, fmt.Sprintf("SELECT count(*) FROM DepartmentSubjects WHERE department_id='%d'", departmentId))
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) AddDepartmentSubjects(departmentId, subjectId int64) error {
	query := fmt.Sprintf("INSERT INTO DepartmentSubjects (department_id, subject_id) VALUES ('%d', '%d')", departmentId, subjectId)

	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteDepartmentSubjects(departmentId, subjectId int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM DepartmentSubjects WHERE department_id='%d' AND subject_id='%d'", departmentId, subjectId))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteAll(departmentId int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM DepartmentSubjects WHERE department_id='%d'", departmentId))

	if res != nil {
		_ = res.Close()
	}
	return err
}
