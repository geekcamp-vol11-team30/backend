package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/geekcamp-vol11-team30/backend/config"
	"go.uber.org/zap"
)

func NewDB(cfg *config.Config, logger *zap.Logger) (*sql.DB, error) {
	password := cfg.MySQL.Password
	ctx := context.Background()
	if cfg.MySQL.IAMAuth {
		awscfg, err := awsconfig.LoadDefaultConfig(ctx)
		if err != nil {
			logger.Error("failed to load aws config", zap.Error(err))
			return nil, err
		}
		endpoint := fmt.Sprintf("%s:%d", cfg.MySQL.Host, cfg.MySQL.Port)
		authenticationToken, err := auth.BuildAuthToken(ctx, endpoint, cfg.AWS.Region, cfg.MySQL.User, awscfg.Credentials)
		if err != nil {
			panic("failed to create authentication token: " + err.Error())
		}
		password = authenticationToken
	}
	// config := mysql.NewConfig()
	// config.Net = "tcp"
	// config.Addr = fmt.Sprintf("%s:%d", cfg.MySQL.Host, cfg.MySQL.Port)
	// config.User = cfg.MySQL.User
	// config.Passwd = cfg.MySQL.Password
	// config.DBName = cfg.MySQL.DBName
	// config.ParseTime = true
	// config.Params = map[string]string{
	// 	"charset": "utf8mb4",
	// }
	// dsn := config.FormatDSN()
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true",
			// "%s:%s@%s:%d/%s?parseTime=true",
			cfg.MySQL.User, password, cfg.MySQL.Host,
			cfg.MySQL.Port, cfg.MySQL.DBName,
		),
	)

	// logger.Info(dsn)
	// logger.Info(fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%d)/%s?parseTime=true",
	// 	cfg.MySQL.Host, cfg.MySQL.User, cfg.MySQL.Host,
	// 	cfg.MySQL.Port, cfg.MySQL.DBName,
	// ))
	if err != nil {
		logger.Error("failed to open db", zap.Error(err))
		return nil, err
	}

	ctx, canncel := context.WithTimeout(context.Background(), 10*time.Second)
	defer canncel()
	if err := db.PingContext(ctx); err != nil {
		logger.Error("failed to ping db", zap.Error(err))
		return nil, err
	}
	return db, nil
}
