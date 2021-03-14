package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id    int64  `gorm:"PRIMARY_KET;AUTO_INCREMENT"`
	Name  string `gorm:"size:32;not null"`
	Email string `gorm:"size:128;not null"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	var user User
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:  "小明",
		Email: "xiaoming@gmail.com",
	})
	db.Where("name = ?", "小明").First(&user)
	fmt.Println("小明的邮箱为" + user.Email)

	var users []User
	db.Find(&users)
	fmt.Println("总用户数：", len(users))
	//user.Email = "xiaoming2@gmail.com"
	//db.Save(user)

	defer db.Close()
}
