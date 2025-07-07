package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Database struct {
	DB *gorm.DB
}

type User struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func (d *Database) Connection() {
	if d.DB != nil {
		return
	}

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Erro ao migrar", err)
	}
	
	d.DB = db
}
