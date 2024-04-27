package config

import (
	"fmt"
	"log"

	"github.com/BrondoL/wedding-be/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConn(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		cfg.DB_HOST,
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_NAME,
		cfg.DB_PORT,
		cfg.DB_SSL_MODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Attendance{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
