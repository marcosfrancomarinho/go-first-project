package database

type User struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
}


type Product struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Price    float64
	Quantity int
}


