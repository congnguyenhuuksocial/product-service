package database

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"product-service/internal/core/entities"
	"product-service/pkg/config"
	"product-service/pkg/constants"
)

type Database struct {
	*gorm.DB
	logger *zap.Logger
}

func NewDatabase(conf *config.Config, logger *zap.Logger) *gorm.DB {
	var err error
	var sqlDB *sql.DB

	logger.Info("Connecting to database...")
	gormDB, err := getDatabaseInstance(conf)
	logger.Info("Database connected")
	db := &Database{gormDB, logger}

	db.RegisterTables()

	if err != nil {
		logger.Fatal("Database connection error", zap.Error(err))
	}
	sqlDB, err = db.DB.DB()
	if err != nil {
		logger.Fatal("sqlDB connection error", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db.DB
}

func getDatabaseInstance(conf *config.Config) (db *gorm.DB, err error) {
	switch conf.Database.Driver {
	case constants.Mysql:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.Database.Username,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Database,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}
	case constants.Postgres:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			conf.Database.Host, conf.Database.Username, conf.Database.Password, conf.Database.Database, conf.Database.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

		if err != nil {
			return nil, fmt.Errorf("failed to connect database: %w", err)
		}
	}
	return db, nil
}

func (d *Database) RegisterTables() {
	fmt.Println(d.DB)
	err := d.DB.AutoMigrate(entities.Product{})

	if err != nil {
		d.logger.Fatal("Database migration error", zap.Error(err))
		os.Exit(0)
	}
	d.logger.Info("Database migration success")
}
