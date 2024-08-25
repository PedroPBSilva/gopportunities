package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InicializeSQLite()

	if err != nil {
		return fmt.Errorf("Error initializing sqlite: %v", err)
	}
	//return errors.New("fake Error")
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	//Inicialize Logger
	logger = NewLogger(prefix)
	return logger
}