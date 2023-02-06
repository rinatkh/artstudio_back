package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/oauth"
	"github.com/rinatkh/artstudio_back/internal/oauth/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	"github.com/rinatkh/artstudio_back/internal/users/models/core"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type OauthUseCase struct {
	cfg      *config.Config
	log      *logrus.Entry
	userRepo users.UserRepository
}

func NewOauthUC(cfg *config.Config, log *logrus.Entry, userRepo users.UserRepository) oauth.UseCase {
	return &OauthUseCase{
		cfg:      cfg,
		log:      log,
		userRepo: userRepo,
	}
}

func (u OauthUseCase) AuthenticateThroughTelergam(params *dto.AuthenticateThroughTelergamRequest) error {
	exist, err := u.userRepo.GetUserById(params.ID)
	if err == nil && exist != nil {
		user := &core.User{
			Id:        params.ID,
			Firstname: params.FirstName,
			Surname:   params.LastName,
			Sex:       consts.NOTHING,
			BirthDate: 0,
			Role:      consts.Student,
			Image:     params.Image,
		}
		_, err := u.userRepo.CreateUser(user)
		if err != nil {
			return err
		}
	}
	return constants.NewCodedError("User exist", fiber.StatusConflict)
}
