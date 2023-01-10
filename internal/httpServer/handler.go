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

	subjectHTTP "github.com/rinatkh/artstudio_back/internal/subjects/delivery/http"
	subjectRepository "github.com/rinatkh/artstudio_back/internal/subjects/repository"
	subjectUsecase "github.com/rinatkh/artstudio_back/internal/subjects/usecase"

	departmentHTTP "github.com/rinatkh/artstudio_back/internal/departments/delivery/http"
	departmentRepository "github.com/rinatkh/artstudio_back/internal/departments/repository"
	departmentUsecase "github.com/rinatkh/artstudio_back/internal/departments/usecase"

	departmentSubjectsRepository "github.com/rinatkh/artstudio_back/internal/DepartmentSubjects/repository"
	departmentSubjectsUsecase "github.com/rinatkh/artstudio_back/internal/DepartmentSubjects/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	postgreConnection, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	userRepo := usersRepository.NewPostgresRepository(postgreConnection, s.log)
	subjectRepo := subjectRepository.NewPostgresRepository(postgreConnection, s.log)
	departmentRepo := departmentRepository.NewPostgresRepository(postgreConnection, s.log)
	departmentSubjectsRepo := departmentSubjectsRepository.NewPostgresRepository(postgreConnection, s.log)

	userUC := usersUsecase.NewUserUC(s.cfg, s.log, userRepo, subjectRepo)
	subjectUC := subjectUsecase.NewSubjectUC(s.cfg, s.log, subjectRepo, userUC)
	departmentSubjectsUC := departmentSubjectsUsecase.NewDepartmentSubjectsUC(s.cfg, s.log, departmentSubjectsRepo, subjectUC, departmentRepo)
	departmentUC := departmentUsecase.NewDepartmentUC(s.cfg, s.log, departmentRepo, subjectUC, departmentSubjectsUC, userUC)

	userHandler := userHTTP.NewUserHandler(userUC, s.log)
	subjectHandler := subjectHTTP.NewSubjectHandler(subjectUC, s.log)
	departmentHandler := departmentHTTP.NewDepartmentHandler(departmentUC, s.log)

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
	subjectHTTP.MapSubjectRoutes(app, subjectHandler, mw)
	departmentHTTP.MapDepartmentRoutes(app, departmentHandler, mw)

	return nil
}
