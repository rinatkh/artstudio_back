package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/statistic_students"
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) statisticStudent.StatisticStudentRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetStatisticStudentById(id string) (*core.StatisticStudent, error) {
	var data []core.StatisticStudent
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.statistic_students WHERE student_id='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) CreateStatisticStudent(statisticStudent *core.StatisticStudent) (*core.StatisticStudent, error) {
	res, err := p.db.Query("INSERT INTO public.statistic_students (student_id) VALUES ($1)", statisticStudent.StudentId)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getStatisticStudent(statisticStudent)
}

func (p postgresRepository) UpdateStatisticStudent(statisticStudent *core.StatisticStudent) (*core.StatisticStudent, error) {
	query := fmt.Sprintf("UPDATE statistic_students SET balance_lessons='%d', done_lessons='%d', need_payment_lessons='%d' where Student_id='%s'", statisticStudent.BalanceLessons, statisticStudent.DoneLessons, statisticStudent.NeedPaymentLessons, statisticStudent.StudentId)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getStatisticStudent(statisticStudent)
}

func (p postgresRepository) DeleteStatisticStudent(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM public.statistic_students WHERE student_id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getStatisticStudent(statisticStudent *core.StatisticStudent) (*core.StatisticStudent, error) {
	var data []core.StatisticStudent
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM public.statistic_students WHERE student_id='%s'", statisticStudent.StudentId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
