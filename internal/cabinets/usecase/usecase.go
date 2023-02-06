package usecase

import (
	"errors"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/cabinets"
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/convert"
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/core"
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
)

type CabinetUseCase struct {
	cfg         *config.Config
	log         *logrus.Entry
	repoCabinet cabinets.CabinetRepository
	userUC      users.UseCase
}

func NewCabinetUC(cfg *config.Config, log *logrus.Entry, repoCabinet cabinets.CabinetRepository, userUC users.UseCase) cabinets.UseCase {
	return &CabinetUseCase{
		cfg:         cfg,
		log:         log,
		repoCabinet: repoCabinet,
		userUC:      userUC,
	}
}

func (u CabinetUseCase) GetCabinet(params *dto.GetCabinetRequest) (*dto.GetCabinetResponse, error) {
	res, err := u.repoCabinet.GetCabinetById(params.CabinetId)
	if err != nil {
		return nil, err
	}
	resTime, err := u.repoCabinet.GetCabinetTimeById(params.CabinetId)
	if errors.Is(err, constants.ErrCabinetTimeDBNotFound) {
		return &dto.GetCabinetResponse{Cabinet: convert.ConvertCabinet2DTO(res, nil)}, nil
	}
	if err != nil {
		return nil, err
	}
	return &dto.GetCabinetResponse{Cabinet: convert.ConvertCabinet2DTO(res, resTime)}, nil
}

func (u CabinetUseCase) DeleteCabinet(params *dto.DeleteCabinetRequest) (*dto.DeleteCabinetResponse, error) {
	_, err := u.repoCabinet.GetCabinetById(params.CabinetId)
	if err != nil {
		return nil, err
	}
	err = u.repoCabinet.DeleteCabinet(params.CabinetId)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteCabinetResponse{}, nil
}

func (u CabinetUseCase) UpdateCabinet(params *dto.UpdateCabinetRequest) (*dto.UpdateCabinetResponse, error) {
	_, err := u.repoCabinet.GetCabinetById(params.CabinetId)
	if err != nil {
		return nil, err
	}
	res, err := u.repoCabinet.UpdateCabinet(&core.Cabinet{Id: params.CabinetId, Name: params.Name})
	if err != nil {
		return nil, err
	}
	resTime, err := u.repoCabinet.GetCabinetTimeById(params.CabinetId)
	if errors.Is(err, constants.ErrCabinetTimeDBNotFound) {
		return &dto.UpdateCabinetResponse{Cabinet: convert.ConvertCabinet2DTO(res, nil)}, nil
	}
	if err != nil {
		return nil, err
	}
	return &dto.UpdateCabinetResponse{Cabinet: convert.ConvertCabinet2DTO(res, resTime)}, nil
}

func (u CabinetUseCase) CreateCabinet(params *dto.CreateCabinetRequest) (*dto.CreateCabinetResponse, error) {
	for _, i := range params.Time {
		if err := utils.IsTime15MinDuration(i.StartTime, i.FinishTime); err != nil {
			return nil, err
		}
	}
	res, err := u.repoCabinet.CreateCabinet(&core.Cabinet{Name: params.Name})
	if err != nil {
		return nil, err
	}
	return &dto.CreateCabinetResponse{Cabinet: convert.ConvertCabinet2DTO(res, nil)}, nil
}
