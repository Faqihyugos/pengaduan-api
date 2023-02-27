package configdb

import (
	"fmt"
	"os"

	entityuser "github.com/faqihyugos/pengaduan-api/entities/entityUser"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := newConfig().String()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entityuser.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

type pgConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

func newConfig() *pgConfig {
	dbConfig := pgConfig{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func (dbConfig *pgConfig) String() string {

	mode := os.Getenv("MODE")
	var dsn string

	if mode == "production" {
		// map config to dsn
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Pass,
			dbConfig.DbName,
			dbConfig.Port,
		)
	} else {
		// map config to dsn
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Pass,
			dbConfig.DbName,
			dbConfig.Port,
		)
	}
	return dsn
}
