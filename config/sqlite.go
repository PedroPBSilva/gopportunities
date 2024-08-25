package config

import (
	"os"

	"github.com/PedroPBSilva/gopportunities.git/schemas"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InicializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	pathDB := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(pathDB)
	if os.IsNotExist(err){
		logger.Info("Database SQLite not found, creating...")
		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Created directory error: %v", err)
			return nil, err
		}
		file, err := os.Create(pathDB)
		if err != nil {
			logger.Errorf("Created database file error: %v", err)
			return nil, err
		}
		file.Close()
	}
	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(pathDB), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}
	// Migrate DB the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}