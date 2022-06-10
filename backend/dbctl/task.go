package dbctl

import (
	"fmt"
	"log"
	"os"
	"time"

	"chat/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func gormConnect() *gorm.DB {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("key")))
	if err != nil {
		fmt.Println("Error loading environment")
		log.Fatal(err)
	}

	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	hostname := os.Getenv("db_hostname")
	name := os.Getenv("db_name")

	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+hostname+")/"+name)
	if err != nil {
		log.Fatal(err)
	}
	return db
} //DBに接続
func InsertNewTask(task structs.JsonReturn) error {
	db := gormConnect()
	defer db.Close()

	var dbData structs.ChatLog

	dbData.Name = task.Name
	dbData.Text = task.Text
	dbData.Time = time.Now().Local().String()

	if err := db.Create(&dbData).Error; err != nil {
		return err
	}
	return nil
}
