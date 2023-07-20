package db

import (
	"fmt"

	"github.com/athunlal/bookNow-Account-Services/pkg/config"
	"github.com/athunlal/bookNow-Account-Services/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(
		&domain.User{},
		&domain.Address{},
	)
	return db, dbErr
}
