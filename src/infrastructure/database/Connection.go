package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Erro ao migrar", err)
	}
	d.DB = db
}
