package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/cabinet"
	"github.com/rinatkh/artstudio_back/internal/cabinet/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
	"time"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) cabinet.CabinetRepository {
	return &postgresRepository{
		db:  db,
		log: log,
	}
}

func (p postgresRepository) GetCabinetById(id int64) (*core.Cabinet, error) {
	var data []core.Cabinet
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinets WHERE id='%d'", id))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrCabinetDBNotFound
	}
	return &data[0], nil
}

func (p postgresRepository) GetCabinetTimesBySubjectId(subjectId int64) (*[]core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE subject_id='%d'", subjectId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrCabinetTimeDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) GetCabinetTimeById(cabinetId int64) (*[]core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE cabinet_id='%d'", cabinetId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrCabinetTimeDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) GetCabinetTimesByCabinetId(cabinetId int64) (*[]core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE cabinet_id='%d' ORDER BY start_time", cabinetId))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrCabinetTimeDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) GetAllCabinetTimes(startTime, finishTime int64) (*[]core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE FROM_UNIXTIME(start_time) >= $1 AND  FROM_UNIXTIME(finish_time) <= $2 ORDER BY start_time"), time.Unix(startTime, 0), time.Unix(finishTime, 0))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrCabinetTimeDBNotFound
	}
	return &data, nil
}
func (p postgresRepository) GetCountCabinets() (int64, error) {
	var data []int64
	err := p.db.Select(&data, "SELECT count(*) FROM cabinetTimes")
	if err != nil {
		return 0, err
	}
	return data[0], nil
}

func (p postgresRepository) GetCabinetTimes(cabinetId, startTime, finishTime int64) (*[]core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE cabinet_id=$1 AND FROM_UNIXTIME(start_time) >= $2 AND  FROM_UNIXTIME(finish_time) <= $3"), cabinetId, time.Unix(startTime, 0), time.Unix(finishTime, 0))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrCabinetTimeDBNotFound
	}
	return &data, nil
}

func (p postgresRepository) CreateCabinet(cabinet *core.Cabinet) (*core.Cabinet, error) {
	res, err := p.db.Query("INSERT INTO cabinets (name) VALUES ($1)", cabinet.Name)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getCabinet(cabinet)
}

func (p postgresRepository) CreateCabinetTime(cabinetTime *core.CabinetTime) (*core.CabinetTime, error) {
	res, err := p.db.Query("INSERT INTO cabinetTimes (subject_id, cabinet_id, start_time, finish_time) VALUES ($1, $2, $3, $4)", cabinetTime.SubjectId, cabinetTime.CabinetId, cabinetTime.StartTime, cabinetTime.FinishTime)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}

	return p.getCabinetTime(cabinetTime)
}

func (p postgresRepository) UpdateCabinet(cabinet *core.Cabinet) (*core.Cabinet, error) {
	query := fmt.Sprintf("UPDATE cabinets SET name='%s' where id='%d'", cabinet.Name, cabinet.Id)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getCabinet(cabinet)
}

func (p postgresRepository) UpdateCabinetTime(cabinetTime *core.CabinetTime) (*core.CabinetTime, error) {
	query := fmt.Sprintf("UPDATE cabinetTimes SET subject_id='%d', cabinet_id='%d', start_time='%d', finish_time='%d' where id='%d'", cabinetTime.SubjectId, cabinetTime.CabinetId, cabinetTime.StartTime, cabinetTime.FinishTime, cabinetTime.Id)
	res, err := p.db.Query(query)
	if res != nil {
		_ = res.Close()
	}
	if err != nil {
		return nil, err
	}
	return p.getCabinetTime(cabinetTime)
}

func (p postgresRepository) DeleteCabinet(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM cabinets WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) DeleteCabinetTime(id int64) error {
	res, err := p.db.Query(fmt.Sprintf("DELETE FROM cabinetTimes WHERE id='%d'", id))

	if res != nil {
		_ = res.Close()
	}
	return err
}

func (p postgresRepository) getCabinet(cabinet *core.Cabinet) (*core.Cabinet, error) {
	var data []core.Cabinet
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinets WHERE name='%d'", cabinet.Name))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}

func (p postgresRepository) getCabinetTime(cabinetTime *core.CabinetTime) (*core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE subject_id='%d' AND cabinet_id='%d' AND start_time='%d' AND finish_time='%d'", cabinetTime.SubjectId, cabinetTime.CabinetId, cabinetTime.StartTime, cabinetTime.FinishTime))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, constants.ErrDB
	}
	return &data[0], nil
}
