package usecase

import (
	"errors"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/auth"
	"github.com/rinatkh/artstudio_back/internal/auth/models/core"
	"github.com/rinatkh/artstudio_back/internal/auth/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
)

type AuthUseCase struct {
	cfg      *config.Config
	log      *logrus.Entry
	repoAuth auth.AuthRepository
	userUC   users.UseCase
}

func NewAuthUC(cfg *config.Config, log *logrus.Entry, repoAuth auth.AuthRepository, userUC users.UseCase) auth.UseCase {
	return &AuthUseCase{
		cfg:      cfg,
		log:      log,
		repoAuth: repoAuth,
		userUC:   userUC,
	}
}

func (u AuthUseCase) LoginUser(params *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	user, err := u.repoAuth.GetUserByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	if err := user.UserPassword.Validate(params.Password); err != nil {
		return nil, err
	}

	authToken, err := utils.GenerateAuthToken(&utils.AuthTokenWrapper{UserID: user.Id}, u.cfg)
	if err != nil {
		return nil, err
	}
	author, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: user.Id})
	if err != nil {
		return nil, err
	}
	return &dto.LoginUserResponse{User: author.User, AuthToken: authToken}, nil
}

func (u AuthUseCase) SignupUser(params *dto.SignupUserRequest) (*dto.SignupUserResponse, error) {
	if _, err := u.repoAuth.GetUserByEmail(params.Email); !errors.Is(err, constants.ErrUserDBNotFound) {
		if err == nil {
			return nil, constants.ErrEmailAlreadyTaken
		}
		return nil, err
	}

	user, err := u.userUC.CreateUser(&dtoUser.CreateUserRequest{
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: params.Middlename,
		BirthDate:  params.BirthDate,
		Sex:        params.Sex,
		Image:      params.Image,
		Role:       params.Role,
	})
	if err != nil {
		return nil, err
	}

	authUser := &core.Auth{
		Id:    user.Id,
		Email: params.Email,
	}
	if err := authUser.UserPassword.Init(params.Password); err != nil {
		return nil, err
	}

	_, err = u.repoAuth.CreateUser(authUser)
	if err != nil {
		return nil, err
	}
	authToken, err := utils.GenerateAuthToken(&utils.AuthTokenWrapper{UserID: user.Id}, u.cfg)
	if err != nil {
		return nil, err
	}

	return &dto.SignupUserResponse{User: user.User, AuthToken: authToken}, nil
}
