package config

import (
	"fmt"
	"golang-be/utils/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DbconV2 *gorm.DB
	ErrdbV2 error
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func init() {
	fmt.Println("Init DBV2")
	if DbOpen() != nil {
		fmt.Println("Can't Open db Postgres")
	}

	DbconV2 = GetDbCon()
}

func DbOpen() error {
	fmt.Println("Config: ", DbURL(&DBConfig{}))

	dbConfig := DBConfig{
		User:     helper.GetEnv("DATABASE_USER", "postgres"),
		Password: helper.GetEnv("DATABASE_PASSWORD", "postgres"),
		DBName:   helper.GetEnv("DATABASE_NAME", "local"),
		Host:     helper.GetEnv("DATABASE_HOST", "127.0.0.1"),
		Port:     helper.GetEnv("DATABASE_PORT", "5432"),
	}
	DbconV2, ErrdbV2 = gorm.Open(postgres.Open(DbURL(&dbConfig)), &gorm.Config{})

	if ErrdbV2 != nil {
		fmt.Print("error when connect db", ErrdbV2)
		// logs.Error("open db Err ", ErrdbV2)
		return ErrdbV2
	}
	return nil
}

func GetDbCon() *gorm.DB {
	return DbconV2
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)
}
