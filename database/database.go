package database

import (
	"fmt"

	"github.com/mskydream/micro/config"
	"github.com/mskydream/micro/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Init(conf *config.Configuration) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Database.Host, conf.Database.Username, conf.Database.Password, conf.Database.Name, conf.Database.Address)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Todo{})
	Database = DbInstance{
		Db: db,
	}
}
