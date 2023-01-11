package constants

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type CodedError struct {
	err  error
	code int
}

func (ce *CodedError) Error() string {
	return ce.err.Error()
}

func (ce *CodedError) Code() int {
	return ce.code
}

func NewCodedError(errMessage string, code int) *CodedError {
	return &CodedError{errors.New(errMessage), code}
}

var (
	// Unathorized
	InputError                 = &CodedError{errors.New("bad json request"), fiber.StatusBadRequest}
	ErrUserDBNotFound          = &CodedError{errors.New("user not found in the database"), fiber.StatusBadRequest}
	ErrDepartmentDBNotFound    = &CodedError{errors.New("department not found in the database"), fiber.StatusBadRequest}
	ErrSubjectDBNotFound       = &CodedError{errors.New("subject not found in the database"), fiber.StatusBadRequest}
	ErrTelegramDBNotFound      = &CodedError{errors.New("telegram not found in the database"), fiber.StatusBadRequest}
	ErrOrderDBNotFound         = &CodedError{errors.New("order not found in the database"), fiber.StatusBadRequest}
	AuthError                  = &CodedError{errors.New("Invalid public api key"), fiber.StatusUnauthorized}
	ErrConvertData             = &CodedError{errors.New("failed to convert"), fiber.StatusInternalServerError}
	ErrDB                      = &CodedError{errors.New("failed working with db"), fiber.StatusInternalServerError}
	ErrGenerateUUID            = &CodedError{errors.New("failed to generate UUID"), fiber.StatusInternalServerError}
	ErrSignToken               = &CodedError{errors.New("failed to sign token"), fiber.StatusInternalServerError}
	ErrParseAuthToken          = &CodedError{errors.New("failed to parse authorization token"), fiber.StatusInternalServerError}
	ErrAuthTokenExpired        = &CodedError{errors.New("authorization token is expired"), fiber.StatusForbidden}
	ErrAuthTokenInvalid        = &CodedError{errors.New("authorization token is invalid"), fiber.StatusUnauthorized}
	ErrUnexpectedSigningMethod = &CodedError{errors.New("unexpected signing method"), fiber.StatusUnauthorized}
	ErrUnauthorized            = &CodedError{errors.New("unauthorized"), fiber.StatusUnauthorized}
	ErrMissingAuthCookie       = &CodedError{errors.New("missing authorization cookie"), fiber.StatusUnauthorized}
	ErrHashInvalid             = &CodedError{errors.New("hash is invalid"), fiber.StatusUnauthorized}

	ErrNoPrivileges      = &CodedError{errors.New("you have no privileges"), fiber.StatusForbidden}
	ErrEmailAlreadyTaken = &CodedError{errors.New("email is taken already by other user"), fiber.StatusConflict}
	ErrPasswordMismatch  = &CodedError{errors.New("password mismatch"), fiber.StatusUnauthorized}
)
