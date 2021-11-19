package db

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Con    *gorm.DB
	DBName string
}

var ConPool = &DB{}

func NewDBConPool(ctx context.Context) error {
	ConPool.DBName = os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ConPool.Con = db.WithContext(ctx)
	mysqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetConnMaxIdleTime(24 * time.Hour)
	maxOpenConnection, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTION"))
	if err != nil {
		return fmt.Errorf("failed to convert string to int: %v", err)
	}
	maxIdleConnection, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTION"))
	if err != nil {
		return fmt.Errorf("failed to convert string to int: %v", err)
	}
	mysqlDB.SetMaxOpenConns(maxOpenConnection)
	mysqlDB.SetMaxIdleConns(maxIdleConnection)
	return nil
}
