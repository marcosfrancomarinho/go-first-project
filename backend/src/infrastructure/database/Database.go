package database

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
}

var (
	instance *gorm.DB
	once     sync.Once
)

func NewDataBase() *gorm.DB {
	once.Do(func() {
		var err error
		instance, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatal("Erro ao conectar com o banco:", err)
		}

		if err = instance.AutoMigrate(&User{}, &Product{}); err != nil {
			log.Fatal("Erro ao migrar", err)
		}
	})
	return instance

}
