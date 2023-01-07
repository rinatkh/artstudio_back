package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/spf13/viper"
)

type MDWManager struct {
	cfg *config.Config
}

func NewMDWManager(cfg *config.Config) *MDWManager {
	return &MDWManager{
		cfg: cfg,
	}
}

func (mw *MDWManager) VerifyTokenMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		cookie := ctx.Cookies(constants.CookieKeyAuthToken)
		if cookie == "" {
			return constants.ErrMissingAuthCookie
		}

		token, err := utils.ParseAuthToken(cookie)
		if err != nil {
			return err
		}

		ctx.Set(constants.CtxKeyUserID, token.UserID)
		return nil
	}
}

func (mw *MDWManager) OAuthTelegramMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var kvs []string
		hash := ""

		f := func(k []byte, v []byte) {
			if string(k) == "hash" {
				hash = string(v[0])
				return
			}
			kvs = append(kvs, string(k)+"="+string(v[0]))
		}

		ctx.Request().URI().QueryArgs().VisitAll(f)
		sort.Strings(kvs)

		var dataCheckString = ""
		for _, s := range kvs {
			if dataCheckString != "" {
				dataCheckString += "\n"
			}
			dataCheckString += s
		}

		sha256hash := sha256.New()

		telegramToken := viper.GetString("service.telegram_token")
		_, _ = io.WriteString(sha256hash, telegramToken)

		hmachash := hmac.New(sha256.New, sha256hash.Sum(nil))
		_, _ = io.WriteString(hmachash, dataCheckString)

		if hash != hex.EncodeToString(hmachash.Sum(nil)) {
			return constants.ErrHashInvalid
		}

		return nil
	}
}
