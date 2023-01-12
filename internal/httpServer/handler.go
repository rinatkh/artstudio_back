package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
	"github.com/rinatkh/artstudio_back/pkg/storage"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	serverLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	authHTTP "github.com/rinatkh/artstudio_back/internal/auth/delivery/http"
	authRepository "github.com/rinatkh/artstudio_back/internal/auth/repository"
	authUsecase "github.com/rinatkh/artstudio_back/internal/auth/usecase"

	userHTTP "github.com/rinatkh/artstudio_back/internal/users/delivery/http"
	usersRepository "github.com/rinatkh/artstudio_back/internal/users/repository"
	usersUsecase "github.com/rinatkh/artstudio_back/internal/users/usecase"

	subjectHTTP "github.com/rinatkh/artstudio_back/internal/subjects/delivery/http"
	subjectRepository "github.com/rinatkh/artstudio_back/internal/subjects/repository"
	subjectUsecase "github.com/rinatkh/artstudio_back/internal/subjects/usecase"

	departmentHTTP "github.com/rinatkh/artstudio_back/internal/departments/delivery/http"
	departmentRepository "github.com/rinatkh/artstudio_back/internal/departments/repository"
	departmentUsecase "github.com/rinatkh/artstudio_back/internal/departments/usecase"

	timeTeacherHTTP "github.com/rinatkh/artstudio_back/internal/time_teachers/delivery/http"
	timeTeacherRepository "github.com/rinatkh/artstudio_back/internal/time_teachers/repository"
	timeTeacherUsecase "github.com/rinatkh/artstudio_back/internal/time_teachers/usecase"

	statisticTeacherHTTP "github.com/rinatkh/artstudio_back/internal/statistic_teachers/delivery/http"
	statisticTeacherRepository "github.com/rinatkh/artstudio_back/internal/statistic_teachers/repository"
	statisticTeacherUsecase "github.com/rinatkh/artstudio_back/internal/statistic_teachers/usecase"

	departmentSubjectsRepository "github.com/rinatkh/artstudio_back/internal/DepartmentSubjects/repository"
	departmentSubjectsUsecase "github.com/rinatkh/artstudio_back/internal/DepartmentSubjects/usecase"

	oauthHTTP "github.com/rinatkh/artstudio_back/internal/oauth/delivery/http"
	oauthUsecase "github.com/rinatkh/artstudio_back/internal/oauth/usecase"

	staticHTTP "github.com/rinatkh/artstudio_back/internal/static/delivery/http"
	staticUsecase "github.com/rinatkh/artstudio_back/internal/static/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	postgreConnection, err := storage.InitPsqlDB(s.cfg)
	if err != nil {
		return err
	}

	authRepo := authRepository.NewPostgresRepository(postgreConnection, s.log)
	userRepo := usersRepository.NewPostgresRepository(postgreConnection, s.log)
	subjectRepo := subjectRepository.NewPostgresRepository(postgreConnection, s.log)
	departmentRepo := departmentRepository.NewPostgresRepository(postgreConnection, s.log)
	departmentSubjectsRepo := departmentSubjectsRepository.NewPostgresRepository(postgreConnection, s.log)
	timeTeacherRepo := timeTeacherRepository.NewPostgresRepository(postgreConnection, s.log)
	statisticTeacherRepo := statisticTeacherRepository.NewPostgresRepository(postgreConnection, s.log)

	userUC := usersUsecase.NewUserUC(s.cfg, s.log, userRepo, subjectRepo, authRepo)
	authUC := authUsecase.NewAuthUC(s.cfg, s.log, authRepo, userUC)
	oauthUC := oauthUsecase.NewOauthUC(s.cfg, s.log, userRepo)
	timeTeacherUC := timeTeacherUsecase.NewSubjectUC(s.cfg, s.log, timeTeacherRepo, userUC)
	statisticTeacherUC := statisticTeacherUsecase.NewStatisticTeacherUC(s.cfg, s.log, statisticTeacherRepo, userUC)
	staticUC := staticUsecase.NewStaticUC(s.cfg, s.log)
	subjectUC := subjectUsecase.NewSubjectUC(s.cfg, s.log, subjectRepo, userUC)
	departmentSubjectsUC := departmentSubjectsUsecase.NewDepartmentSubjectsUC(s.cfg, s.log, departmentSubjectsRepo, subjectUC, departmentRepo)
	departmentUC := departmentUsecase.NewDepartmentUC(s.cfg, s.log, departmentRepo, subjectUC, departmentSubjectsUC, userUC)

	authHandler := authHTTP.NewAuthHandler(authUC, s.log, s.cfg)
	oauthHandler := oauthHTTP.NewOauthHandler(oauthUC, s.log, s.cfg)
	staticHandler := staticHTTP.NewStaticHandler(staticUC, s.log, s.cfg)
	userHandler := userHTTP.NewUserHandler(userUC, s.log)
	subjectHandler := subjectHTTP.NewSubjectHandler(subjectUC, s.log)
	departmentHandler := departmentHTTP.NewDepartmentHandler(departmentUC, s.log)
	timeTeacherHandler := timeTeacherHTTP.NewTimeTeacherHandler(timeTeacherUC, s.log)
	statisticTeacherHandler := statisticTeacherHTTP.NewStatisticTeacherHandler(statisticTeacherUC, s.log)

	app.Use(serverLogger.New())
	if _, ok := os.LookupEnv("LOCAL"); !ok {
		app.Use(recover.New())
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	mw := middleware.NewMDWManager(s.cfg)

	authHTTP.MapAuthRoutes(app, authHandler)
	oauthHTTP.MapOauthRoutes(app, oauthHandler, mw)
	userHTTP.MapUserRoutes(app, userHandler, mw)
	staticHTTP.MaStaticRoutes(app, staticHandler, mw)
	subjectHTTP.MapSubjectRoutes(app, subjectHandler, mw)
	departmentHTTP.MapDepartmentRoutes(app, departmentHandler, mw)
	timeTeacherHTTP.MapTimeTeacherRoutes(app, timeTeacherHandler, mw)
	statisticTeacherHTTP.MapStatisticTeacherRoutes(app, statisticTeacherHandler, mw)

	return nil
}
