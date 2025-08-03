package db

import (
	"fmt"

	"github.com/sgokul961/GO-PROJECT/pkg/config"
	"github.com/sgokul961/GO-PROJECT/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})
	db.AutoMigrate(&domain.Admin{})

	return db, dbErr
}
