package oauth

import "github.com/rinatkh/artstudio_back/internal/oauth/models/dto"

type UseCase interface {
	AuthenticateThroughTelergam(params *dto.AuthenticateThroughTelergamRequest) error
}
