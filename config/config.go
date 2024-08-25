package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	//return errors.New("fake Error")
	return nil
}

func GetLogger(prefix string) *Logger {
	//Inicialize Logger
	logger = NewLogger(prefix)
	return logger
}