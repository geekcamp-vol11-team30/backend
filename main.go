package main

import (
	"context"
	"fmt"
	"net"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/db"
	applogger "github.com/geekcamp-vol11-team30/backend/logger"
	"github.com/geekcamp-vol11-team30/backend/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := run(context.Background(), logger); err != nil {
		sugar.Fatal(err)
	}
}

func run(ctx context.Context, logger *zap.Logger) error {
	logger.Info("magische starting...")

	cfg, err := config.New()
	if err != nil {
		return err
	}
	db, err := db.NewDB(cfg, logger)
	if err != nil {
		return err
	}
	boil.SetDB(db)
	boil.DebugMode = cfg.SqlLog

	// ur := repository.NewUserRepository(db)
	// ar := repository.NewAuthRepository(db)
	// er := repository.NewEventRepository(db)
	// oar := repository.NewOauthRepository(db)
	// gs := service.NewGoogleService(cfg, oar, ur)
	// ms := service.NewMicrosoftService(cfg, oar, ur)
	// uv := validator.NewUserValidator()
	// uu := usecase.NewUserUsecase(ur, oar, er, uv, gs, ms)
	// au := usecase.NewAuthUsecase(cfg, logger, ar)
	// eu := usecase.NewEventUsecase(cfg, er)
	// oau := usecase.NewOauthUsecase(cfg, oar, ur, gs, ms, uu)

	// em := middleware.NewErrorMiddleware(logger, uu)
	// atm := middleware.NewAccessTimeMiddleware()
	// am := middleware.NewAuthMiddleware(cfg, logger, au, uu)

	// uc := controller.NewUserController(uu)
	// ac := controller.NewAuthController(cfg, uu, au)
	// ec := controller.NewEventController(eu)
	// oc := controller.NewOauthController(cfg, oau, uu, au)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		logger.Fatal("failed to listen port", zap.Error(err))
	}

	// err = util.SendMail(*cfg, "tak848.0428771@gmail.com", "konnitiha", "hello")
	// fmt.Println(err)

	// e := router.NewRouter(cfg, logger, em, atm, am, uc, ac, ec, oc)
	e := echo.New()
	// enable log
	e.GET("/health", func(c echo.Context) error {
		return util.JSONResponse(c, 200, "OK")
	})
	applogger.SetRequestLoggerToEcho(e, logger)
	s := NewServer(e, l, logger)
	return s.Run(ctx)
}
