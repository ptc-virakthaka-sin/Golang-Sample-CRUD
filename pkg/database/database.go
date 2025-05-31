package database

import (
	"learn-fiber/config"
	"learn-fiber/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() (*gorm.DB, error) {
	log.Println("database is connecting...")

	dsn := config.Cfg.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		TranslateError:         true,
	})
	if err != nil {
		return nil, err
	}

	log.Println("database is connected")

	if config.Cfg.DB.AutoMigrate {
		log.Println("auto migration is running...")
		_ = db.AutoMigrate(
			&model.Department{},
			&model.Student{},
			&model.Token{},
			&model.User{},
		)
	}
	return db, nil
}
