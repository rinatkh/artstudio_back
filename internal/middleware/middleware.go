package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/users"
	"io"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
)

type MDWManager struct {
	cfg      *config.Config
	userRepo users.UserRepository
}

func NewMDWManager(cfg *config.Config, userRepo users.UserRepository) *MDWManager {
	return &MDWManager{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (mw *MDWManager) VerifyTokenMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		cookie := ctx.Cookies(constants.CookieKeyAuthToken)
		if cookie == "" {
			return constants.ErrMissingAuthCookie
		}

		token, err := utils.ParseAuthToken(cookie, mw.cfg)
		if err != nil {
			return err
		}

		ctx.Request().Header.Add(constants.CtxKeyUserID, token.UserID)
		return ctx.Next()
	}
}

func (mw *MDWManager) VerifyAdminMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, err := mw.userRepo.GetUserById(ctx.Get(constants.CtxKeyUserID))
		if err != nil {
			return err
		}
		if user.Role != consts.Admin {
			return constants.ErrNoPrivileges
		}
		return ctx.Next()
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

		_, _ = io.WriteString(sha256hash, mw.cfg.Service.TelegramToken)

		hmachash := hmac.New(sha256.New, sha256hash.Sum(nil))
		_, _ = io.WriteString(hmachash, dataCheckString)

		if hash != hex.EncodeToString(hmachash.Sum(nil)) {
			return constants.ErrHashInvalid
		}

		return ctx.Next()
	}
}
