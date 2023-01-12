package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) statisticTeacher.StatisticTeacherRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetStatisticTeacherById(id string) (*core.StatisticTeacher, error) {
	var data []core.StatisticTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM statistic_teachers WHERE teacher_id='%s'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return &data[0], nil
}

func (p postgresRepository) CreateStatisticTeacher(StatisticTeacher *core.StatisticTeacher) (*core.StatisticTeacher, error) {
	res, err := p.db.Query("INSERT INTO statistic_teachers (teacher_id) VALUES ($1)", StatisticTeacher.TeacherId)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getStatisticTeacher(StatisticTeacher)
}

func (p postgresRepository) UpdateStatisticTeacher(statisticTeacher *core.StatisticTeacher) (*core.StatisticTeacher, error) {
	query := fmt.Sprintf("UPDATE statistic_teachers SET future_lessons='%d', need_dz_lessons='%d', past_lessons='%d' where teacher_id='%s'", statisticTeacher.FutureLessons, statisticTeacher.NeedDzLessons, statisticTeacher.PastLessons, statisticTeacher.TeacherId)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getStatisticTeacher(statisticTeacher)
}

func (p postgresRepository) DeleteStatisticTeacher(id string) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM statistic_teachers WHERE teacher_id='%s'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getStatisticTeacher(statisticTeacher *core.StatisticTeacher) (*core.StatisticTeacher, error) {
	var data []core.StatisticTeacher
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM statistic_teachers WHERE teacher_id='%s'", statisticTeacher.TeacherId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
