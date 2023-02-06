package repository

import (
	"errors"
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

func (p postgresRepository) checkTimeTeachers(teacherId string, startTime, finishTime int64) (*[]core.TimeTeacher, error) {
	var data []core.TimeTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM time_teachers WHERE teacher_id='%s' AND (start_time >= $1 and start_time <= $2) or (finish_time >= $1 and finish_time <= $2)", teacherId), startTime, finishTime)
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
		return nil, constants.NewCodedError("teacher time not found", fiber.StatusConflict)
	}
	return &data[0], nil
}

func (p postgresRepository) CreateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error) {
	if _, err := p.checkTimeTeachers(timeTeacher.TeacherId, timeTeacher.StartTime, timeTeacher.FinishTime); !errors.Is(err, constants.NewCodedError("teacher not free at this time", fiber.StatusConflict)) {
		return nil, constants.NewCodedError("teacher already have this time for schedules", fiber.StatusConflict)
	}
	res, err := p.db.Query("INSERT INTO time_teachers (teacher_id, start_time, finish_time) VALUES ($1, $2, $3)", timeTeacher.TeacherId, timeTeacher.StartTime, timeTeacher.FinishTime)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getTimeTeacher(timeTeacher)
}

func (p postgresRepository) checkTimeTeachersForUpdate(teacherId string, startTime, finishTime, id int64) (*[]core.TimeTeacher, error) {
	var data []core.TimeTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM time_teachers WHERE teacher_id='%s' AND id <> '%d' AND (start_time >= $1 and start_time <= $2) or (finish_time >= $1 and finish_time <= $2)", teacherId, id), startTime, finishTime)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.NewCodedError("teacher not free at this time", fiber.StatusConflict)
	}
	return &data, nil
}

func (p postgresRepository) UpdateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error) {
	if _, err := p.checkTimeTeachersForUpdate(timeTeacher.TeacherId, timeTeacher.StartTime, timeTeacher.FinishTime, timeTeacher.Id); !errors.Is(err, constants.NewCodedError("teacher not free at this time", fiber.StatusConflict)) {
		return nil, constants.NewCodedError("teacher already have this time for schedules", fiber.StatusConflict)
	}
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
