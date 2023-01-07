package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	"github.com/rinatkh/artstudio_back/internal/subjects/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) subjects.SubjectRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetSubjectById(id int64) (*core.Subject, error) {
	var data []core.Subject
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM subjects WHERE id='%d'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetSubjects(limit, offset int64) (*[]core.Subject, int64, error) {
	var data []core.Subject
	queryStr := "SELECT * FROM subjects WHERE id <> ''"
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
	err = p.db.Select(&length, "SELECT count(*) FROM subjects")
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateSubject(subject *core.Subject) (*core.Subject, error) {
	if *subject.Description == "" {
		*subject.Description = "NULL"
	}
	res, err := p.db.Query("INSERT INTO subjects (name, image, description, teacher_id) VALUES ($1, $2, $3, $4)", subject.Name, subject.Image, *subject.Description, subject.TeacherId)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getSubject(subject)
}

func (p postgresRepository) UpdateSubject(subject *core.Subject) (*core.Subject, error) {
	var query string
	if *subject.Description == "" {
		query = fmt.Sprintf("UPDATE subjects SET name='%s', image='%s', description=NULL, teacher_id='%s' where id='%d'", subject.Name, subject.Image, subject.TeacherId, subject.Id)
	} else {
		query = fmt.Sprintf("UPDATE subjects SET name='%s', image='%s', description='%s', teacher_id='%s' where id='%d'", subject.Name, subject.Image, *subject.Description, subject.TeacherId, subject.Id)
	}
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getSubject(subject)
}

func (p postgresRepository) DeleteSubject(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM subjects WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getSubject(subject *core.Subject) (*core.Subject, error) {
	var data []core.Subject
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM subjects WHERE name='%s' AND image='%s' and description='%s' and teacher_id='%s'", subject.Name, subject.Image, *subject.Description, subject.TeacherId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
