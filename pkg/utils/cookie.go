package utils

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreateHttpOnlyCookie(name, value string, ttl int64) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(time.Second * time.Duration(ttl)),
		Path:     "/",
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
	}
}
