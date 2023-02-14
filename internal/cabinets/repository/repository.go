package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rinatkh/artstudio_back/internal/cabinets"
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type postgresRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewPostgresRepository(db *sqlx.DB, log *logrus.Entry) cabinets.CabinetRepository {
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
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE subject_id='%d' ORDER BY start_time", subjectId))
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
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE start_time >= $1 AND  finish_time <= $2 ORDER BY start_time"), startTime, finishTime)
	if err != nil {
		return nil, err
	}
	return p.transformFromTakenTimeToFreeTime(startTime, finishTime, &data)
}

func (p postgresRepository) transformFromTakenTimeToFreeTime(startTime, finishTime int64, times *[]core.CabinetTime) (*[]core.CabinetTime, error) {
	var result []core.CabinetTime
	if times == nil {
		result = []core.CabinetTime{{
			StartTime:  startTime,
			FinishTime: finishTime,
		}}
		return &result, nil
	}
	count, err := p.GetCountCabinets()
	if err != nil {
		return nil, err
	}
	var buildTimeTemp []core.CabinetTime
	var point int
	var isCabinetChange bool
	for i := int64(1); i <= count; i++ {
		isCabinetChange = true
		for j := 0; j < len(*times)-1; j++ {
			if (*times)[j].StartTime < startTime && (*times)[j].FinishTime > startTime && (*times)[j].CabinetId == i {
				(*times)[j].StartTime = startTime
			}
			if (*times)[j].StartTime < finishTime && (*times)[j].FinishTime > finishTime && (*times)[j].CabinetId == i {
				(*times)[j].FinishTime = finishTime
			}
			if (*times)[j].StartTime >= startTime && (*times)[j].FinishTime <= finishTime && (*times)[j].CabinetId == i {
				if isCabinetChange {
					isCabinetChange = false
					buildTimeTemp = append(buildTimeTemp, (*times)[j])
				}
				if (*times)[j].FinishTime == (*times)[j+1].StartTime {
					buildTimeTemp[point].FinishTime = (*times)[j+1].FinishTime
				} else {
					point += 1
					buildTimeTemp = append(buildTimeTemp, (*times)[j+1])
				}
			}
		}
	}
	for i := int64(1); i <= count; i++ {
		isCabinetChange = true
		for j := 0; j < len(buildTimeTemp)-1; j++ {
			if buildTimeTemp[j].CabinetId == i {
				if isCabinetChange {
					isCabinetChange = false
					if buildTimeTemp[j].StartTime != startTime {
						result = append(result, core.CabinetTime{
							StartTime:  startTime,
							FinishTime: buildTimeTemp[j].StartTime,
						})
					}
				}
				if j == len(buildTimeTemp)-2 || buildTimeTemp[j+1].CabinetId != buildTimeTemp[j+2].CabinetId {
					if buildTimeTemp[j+1].FinishTime != finishTime {
						result = append(result, core.CabinetTime{
							StartTime:  buildTimeTemp[j+1].StartTime,
							FinishTime: finishTime,
						})
					}
					continue
				}
				result = append(result, core.CabinetTime{
					StartTime:  buildTimeTemp[j].FinishTime,
					FinishTime: buildTimeTemp[j+1].StartTime,
				})

			}
		}
	}
	if len(result) == 0 {
		return nil, constants.ErrCabinetTimeDBNotFound
	}
	return &result, nil
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
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE cabinet_id=$1 AND start_time >= $2 AND finish_time <= $3"), cabinetId, startTime, finishTime)
	if err != nil {
		return nil, err
	}
	return p.transformFromTakenTimeToFreeTime(startTime, finishTime, &data)
}
func (p postgresRepository) checkCabinetTimes(cabinetId, startTime, finishTime int64) (*[]core.CabinetTime, error) {
	var data []core.CabinetTime
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinetTimes WHERE cabinet_id=$1 AND (start_time >= $2 AND start_time <= $3) OR (finish_time >= $2 AND finish_time <= $3) "), cabinetId, startTime, finishTime)
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
	if _, err := p.checkCabinetTimes(cabinetTime.CabinetId, cabinetTime.StartTime, cabinetTime.FinishTime); !errors.Is(err, constants.ErrCabinetTimeDBNotFound) {
		return nil, constants.NewCodedError("Cabinet isn't free at this time, choose another one", fiber.StatusConflict)
	}
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
	err := p.db.Select(&data, fmt.Sprintf("SELECT * FROM cabinets WHERE name='%s'", cabinet.Name))
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
