package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/schedules"
	"github.com/rinatkh/artstudio_back/internal/schedules/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
	"time"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) schedules.ScheduleRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetScheduleById(id int64) (*core.Schedule, error) {
	var data []core.Schedule
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM schedules WHERE id='%d'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrScheduleDBNotFound
	}
	return &data[0], nil
}

func (p postgresRepository) GetTeacherSchedules(teacherId string, startTime, finishTime int64) (*[]core.Schedule, error) {
	var data []core.Schedule
	queryStr := fmt.Sprintf("SELECT * FROM schedules WHERE subject_id in (SELECT id FROM subjects WHERE teacher_id = %s) AND FROM_UNIXTIME(when_time) >= $1 AND  FROM_UNIXTIME(when_time) <= $2", teacherId)
	err := p.db.Select(
		&data, queryStr, time.Unix(startTime, 0), time.Unix(finishTime, 0).Add(consts.LessonTime+consts.Duration))

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrScheduleDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) GetStudentSchedules(studentId string, startTime, finishTime int64) (*[]core.Schedule, error) {
	var data []core.Schedule
	queryStr := fmt.Sprintf("SELECT * FROM schedules WHERE student_id = %s AND FROM_UNIXTIME(when_time) >= $1 AND  FROM_UNIXTIME(when_time) <= $2", studentId)
	err := p.db.Select(
		&data, queryStr, time.Unix(startTime, 0), time.Unix(finishTime, 0).Add(consts.LessonTime+consts.Duration))

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrScheduleDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) GetSchedules(startTime, finishTime int64) (*[]core.Schedule, error) {
	var data []core.Schedule
	queryStr := fmt.Sprintf("SELECT * FROM schedules WHERE FROM_UNIXTIME(when_time) >= $1 AND  FROM_UNIXTIME(when_time) <= $2")
	err := p.db.Select(
		&data, queryStr, time.Unix(startTime, 0), time.Unix(finishTime, 0).Add(consts.LessonTime+consts.Duration))

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrScheduleDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) GetSubjectSchedules(subjectId int64, startTime, finishTime int64) (*[]core.Schedule, error) {
	var data []core.Schedule
	queryStr := fmt.Sprintf("SELECT * FROM schedules WHERE subject_id  = %d AND FROM_UNIXTIME(when_time) >= $1 AND  FROM_UNIXTIME(when_time) <= $2", subjectId)
	err := p.db.Select(
		&data, queryStr, time.Unix(startTime, 0), time.Unix(finishTime, 0).Add(consts.LessonTime+consts.Duration))

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrScheduleDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) CreateSchedule(schedule *core.Schedule) (*core.Schedule, error) {
	res, err := p.db.Query("INSERT INTO schedules (subject_id, student_id, cabinet_id, when_time, is_paid, is_finished, is_confirmed, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", schedule.SubjectId, schedule.StudentId, schedule.CabinetId, schedule.WhenTime, schedule.IsPaid, schedule.IsFinished, schedule.IsConfirmed, schedule.Description)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getSchedule(schedule)
}

func (p postgresRepository) UpdateSchedule(schedule *core.Schedule) (*core.Schedule, error) {
	query := fmt.Sprintf("UPDATE schedules SET subject_id='%d', student_id='%s', cabinet_id='%d', when_time='%d', is_paid='%t', is_finished='%t', is_confirmed='%t', description='%s' where id = %d", schedule.SubjectId, schedule.StudentId, schedule.CabinetId, schedule.WhenTime, schedule.IsPaid, schedule.IsFinished, schedule.IsConfirmed, schedule.Description, schedule.Id)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getSchedule(schedule)
}

func (p postgresRepository) DeleteSchedule(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM schedules WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteScheduleBySubjectId(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM schedules WHERE subject_id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteScheduleByStudentId(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM schedules WHERE student_id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getSchedule(schedule *core.Schedule) (*core.Schedule, error) {
	var data []core.Schedule

	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM schedules WHERE subject_id='%d' AND student_id='%s' AND cabinet_id='%d' AND when_time='%d' AND is_paid='%t' AND is_finished='%t' AND is_confirmed='%t' AND description='%s'", schedule.SubjectId, schedule.StudentId, schedule.CabinetId, schedule.WhenTime, schedule.IsPaid, schedule.IsFinished, schedule.IsConfirmed, schedule.Description))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
