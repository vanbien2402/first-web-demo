package rds

import (
	"fmt"
	"github.com/vanbien2402/first-web-demo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SslMode)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	},
	))
	if err != nil {
		return nil, err
	}
	_ = db.AutoMigrate(&models.User{})
	log.Println("DB connected")
	return db, nil
}
