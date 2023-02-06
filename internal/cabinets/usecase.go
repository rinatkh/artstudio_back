package cabinets

import (
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/dto"
)

type UseCase interface {
	GetCabinet(params *dto.GetCabinetRequest) (*dto.GetCabinetResponse, error)
	DeleteCabinet(params *dto.DeleteCabinetRequest) (*dto.DeleteCabinetResponse, error)
	UpdateCabinet(params *dto.UpdateCabinetRequest) (*dto.UpdateCabinetResponse, error)
	CreateCabinet(params *dto.CreateCabinetRequest) (*dto.CreateCabinetResponse, error)
}
