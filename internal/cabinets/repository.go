package cabinets

import (
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/core"
)

type CabinetRepository interface {
	GetCabinetById(id int64) (*core.Cabinet, error)
	GetCountCabinets() (int64, error)
	DeleteCabinet(id int64) error
	CreateCabinet(cabinet *core.Cabinet) (*core.Cabinet, error)
	UpdateCabinet(cabinet *core.Cabinet) (*core.Cabinet, error)

	GetCabinetTimeById(cabinetId int64) (*[]core.CabinetTime, error)
	DeleteCabinetTime(id int64) error
	CreateCabinetTime(cabinetTime *core.CabinetTime) (*core.CabinetTime, error)
	UpdateCabinetTime(cabinetTime *core.CabinetTime) (*core.CabinetTime, error)

	GetCabinetTimesBySubjectId(subjectId int64) (*[]core.CabinetTime, error)
	GetCabinetTimesByCabinetId(cabinetId int64) (*[]core.CabinetTime, error)

	GetAllCabinetTimes(startTime, finishTime int64) (*[]core.CabinetTime, error)
	GetCabinetTimes(cabinetId, startTime, finishTime int64) (*[]core.CabinetTime, error)
}
