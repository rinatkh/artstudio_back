package usecase

import (
	"crypto/tls"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/auth"
	"github.com/rinatkh/artstudio_back/internal/auth/models/core"
	"github.com/rinatkh/artstudio_back/internal/auth/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
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
	if user.IsConfirmed == false {
		return nil, constants.NewCodedError("Почта не была успешно подтверждена", fiber.StatusConflict)
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

func (u AuthUseCase) SignupUser(id string) (*dto.SignupUserResponse, error) {
	user, err := u.repoAuth.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if user.IsConfirmed == true {
		return &dto.SignupUserResponse{Message: "Почта уже была успешна подверждена"}, err
	}
	user.IsConfirmed = true
	user, err = u.repoAuth.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &dto.SignupUserResponse{Message: "Почта успешно подверждена"}, err
}

func (u AuthUseCase) sendMail(email, password, id string) error {
	d := gomail.NewDialer(u.cfg.Email.Host, u.cfg.Email.Port, u.cfg.Email.Email, u.cfg.Email.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	message := "Email: " + email +
		"\nПароль: " + password +
		"\nПодтверждение почты: http://192.0.0.1/user/" + id

	m := gomail.NewMessage()
	m.SetHeader("From", u.cfg.Email.Email)
	m.SetHeader("To", email)
	m.SetBody("text/plain", message)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}

func (u AuthUseCase) SignupUserPre(params *dto.SignupUserRequest) (*dto.SignupPreResponse, error) {
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

	err = u.sendMail(params.Email, params.Password, user.Id)
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

	return &dto.SignupPreResponse{Message: "Поздравляем, вы успешно зарегистрировались. Подтвердите регистрацию по ссылке, отправленной на почту."}, nil
}
