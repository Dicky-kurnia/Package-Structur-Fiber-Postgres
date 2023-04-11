package config

import (
	"database/sql"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewFMCGPostgresDB() *gorm.DB {
	dsn := os.Getenv("POSTGRES_DSN")

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	setMaxIdleConns, err := strconv.Atoi(os.Getenv("POSTGRES_SET_MAX_IDLE_CONNS"))
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(setMaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	setMaxOpenConns, err := strconv.Atoi(os.Getenv("POSTGRES_SET_MAX_OPEN_CONNS"))
	if err != nil {
		panic(err)
	}
	conn.SetMaxOpenConns(setMaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	setConnMaxLifetimeMinute, err := strconv.Atoi(os.Getenv("POSTGRES_SET_CONN_MAX_LIFETIME_MINUTE"))
	if err != nil {
		panic(err)
	}
	conn.SetConnMaxLifetime(time.Duration(setConnMaxLifetimeMinute) * time.Minute)

	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	setConnMaxIdleTimeMinute, err := strconv.Atoi(os.Getenv("POSTGRES_SET_CONN_MAX_IDLE_TIME_MINUTE"))
	if err != nil {
		panic(err)
	}
	conn.SetConnMaxIdleTime(time.Duration(setConnMaxIdleTimeMinute) * time.Minute)

	postgres, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
		// PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	// Debug
	// postgres = postgres.Debug()

	// err = postgres.AutoMigrate(
	// 	&model.BranchOffice{},
	// 	&model.Driver{},
	// 	&model.Sales{},
	// 	&model.Orders{},
	// 	&model.Product{},
	// 	&model.OrderDetails{},
	// 	&model.ProofOfPayment{},
	// 	&model.ProofOfPaymentDetails{},
	// 	&model.OrderInfoDetails{},
	// )

	// if err != nil {
	// 	panic(err)
	// }

	return postgres
}
