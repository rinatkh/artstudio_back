package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
	"github.com/rinatkh/artstudio_back/pkg/storage"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	userHTTP "github.com/rinatkh/artstudio_back/internal/users/delivery/http"
	usersRepository "github.com/rinatkh/artstudio_back/internal/users/repository"
	usersUsecase "github.com/rinatkh/artstudio_back/internal/users/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	postgreConnection, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	userRepo := usersRepository.NewPostgresRepository(postgreConnection, s.log)

	userUC := usersUsecase.NewUserUC(s.cfg, s.log, userRepo)

	userHandler := userHTTP.NewUserHandler(userUC, s.log)

	app.Use(serverLogger.New())
	if _, ok := os.LookupEnv("LOCAL"); !ok {
		app.Use(recover.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	mw := middleware.NewMDWManager(s.cfg)

	userHTTP.MapUserRoutes(app, userHandler, mw)

	return nil
}
