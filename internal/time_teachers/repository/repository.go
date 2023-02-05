package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	timeTeacher "github.com/rinatkh/artstudio_back/internal/time_teachers"
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) timeTeacher.TimeTeacherRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}
func (p postgresRepository) GetTimeTeachers(teacherId string, startTime, finishTime int64) (*[]core.TimeTeacher, error) {
	var data []core.TimeTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM time_teachers WHERE teacher_id='%s' AND start_time >= $1 and finish_time <= $2", teacherId), startTime, finishTime)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.NewCodedError("teacher not free at this time", fiber.StatusConflict)
	}
	return &data, nil
}

func (p postgresRepository) GetTimeTeacherById(id int64) (*core.TimeTeacher, error) {
	var data []core.TimeTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM time_teachers WHERE id='%d'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.NewCodedError("teacher not free at this time", fiber.StatusConflict)
	}
	return &data[0], nil
}

func (p postgresRepository) CreateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error) {
	res, err := p.db.Query("INSERT INTO time_teachers (teacher_id, start_time, finish_time) VALUES ($1, $2, $3)", timeTeacher.TeacherId, timeTeacher.StartTime, timeTeacher.FinishTime)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getTimeTeacher(timeTeacher)
}

func (p postgresRepository) UpdateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error) {
	query := fmt.Sprintf("UPDATE time_teachers SET teacher_id='%s', start_time=$1, finish_time=$2 where id='%d'", timeTeacher.TeacherId, timeTeacher.Id)
	res, err := p.db.Query(query, timeTeacher.StartTime, timeTeacher.FinishTime)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getTimeTeacher(timeTeacher)
}

func (p postgresRepository) DeleteTimeTeacher(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM time_teachers WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteTimeTeacherByUserId(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM time_teachers WHERE teacher_id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error) {
	var data []core.TimeTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM time_teachers WHERE teacher_id='%s' AND start_time=$1 and finish_time=$2", timeTeacher.TeacherId), timeTeacher.StartTime, timeTeacher.FinishTime)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
