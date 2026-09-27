package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(".env file not found")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	/*dsn := "root:pass123@tcp(localhost:3306)/cloudtool?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db*/
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
	)

	for i := 0; i < 10; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			return
		}
		fmt.Println("Waiting for MySQL...")
		time.Sleep(3 * time.Second)
	}

	panic("Could not connect to database")
}
