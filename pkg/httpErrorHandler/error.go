package httpErrorHandler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/pkg/constants"
)

type HttpErrorHandler struct {
	showUnknownErrors bool
}

func NewErrorHandler(c *config.Config) *HttpErrorHandler {
	return &HttpErrorHandler{
		showUnknownErrors: c.Server.ShowUnknownErrorsInResponse,
	}
}

type responseMsg struct {
	Message string `json:"message"`
}

func (handler *HttpErrorHandler) Handler(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	var response = responseMsg{Message: err.Error()}

	e := err
	for e != nil {
		if ce, ok := e.(*constants.CodedError); ok {
			statusCode = ce.Code()
			response.Message = ce.Error()
			if response.Message == "" && handler.showUnknownErrors {
				if statusCode == fiber.StatusInternalServerError {
					response.Message = "internal server error"
				} else {
					response.Message = fmt.Sprintf("%s", e.Error())
				}
			}
			return c.Status(statusCode).JSON(response)
		} else {
			e = errors.Unwrap(e)
		}
	}
	return c.Status(statusCode).JSON(response)
}
