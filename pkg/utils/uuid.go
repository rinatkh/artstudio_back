package utils

import (
	"fmt"
	"github.com/rinatkh/artstudio_back/pkg/constants"

	"github.com/gofrs/uuid"
)

func GenUUID() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("%w: %v", constants.ErrGenerateUUID, err)
	}
	return uuid.String(), nil
}
