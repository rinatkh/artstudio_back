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
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.subjects WHERE id='%d'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) GetSubjects(limit, offset int64, teacherId string) (*[]core.Subject, int64, error) {
	var data []core.Subject
	queryStr := "SELECT * FROM public.subjects WHERE id > 0"

	if teacherId != "" {
		queryStr += fmt.Sprintf(" AND teacher_id='%s'", teacherId)
	}
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
	q := "SELECT count(*) FROM public.subjects"
	if teacherId != "" {
		q += fmt.Sprintf(" WHERE teacher_id='%s'", teacherId)
	}
	err = p.db.Select(&length, q)
	if err != nil {
		return nil, 0, err
	}
	return &data, length[0], nil
}

func (p postgresRepository) CreateSubject(subject *core.Subject) (*core.Subject, error) {
	res, err := p.db.Query("INSERT INTO public.subjects (name, cabinet_id, image, description, teacher_id) VALUES ($1, $2, $3, $4, $5)", subject.Name, subject.CabinetId, subject.Image, subject.Description, subject.TeacherId)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getSubject(subject)
}

func (p postgresRepository) UpdateSubject(subject *core.Subject) (*core.Subject, error) {
	query := fmt.Sprintf("UPDATE subjects SET name='%s', cabinet_id='%d', image='%s', description='%s', teacher_id='%s' where id='%d'", subject.Name, subject.CabinetId, subject.Image, subject.Description, subject.TeacherId, subject.Id)
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
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM public.subjects WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getSubject(subject *core.Subject) (*core.Subject, error) {
	var data []core.Subject
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.subjects WHERE name='%s' AND cabinet_id='%d' AND image='%s' and description='%s' and teacher_id='%s'", subject.Name, subject.CabinetId, subject.Image, subject.Description, subject.TeacherId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
