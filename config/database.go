package config

import (
	"fmt"
	"golang-be/utils/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseOpen ..
func DatabaseOpen() (*gorm.DB, error) {
	dbuser := helper.GetEnv("DATABASE_USER", "postgres")
	dbpass := helper.GetEnv("DATABASE_PASSWORD", "")
	dbname := helper.GetEnv("DATABASE_NAME", "local")
	dbaddres := helper.GetEnv("DATABASE_HOST", "127.0.0.1")
	dbport := helper.GetEnv("DATABASE_PORT", "5432")
	sslmode := helper.GetEnv("DATABASE_SSLMODE", "disable")
	dbtimeout := helper.GetEnv("DATABASE_TIMEOUT", "60")

	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", dbaddres, dbport, dbuser, dbpass, dbname, sslmode, dbtimeout)

	fmt.Println("DB CONNECTION :", args)

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		fmt.Println("Error Connecting Database : ", err)
		return db, err
	}

	sqldb, _ := db.DB()
	sqldb.SetConnMaxLifetime(10)
	sqldb.SetMaxIdleConns(25)
	sqldb.SetMaxOpenConns(10)

	return db, nil
}
