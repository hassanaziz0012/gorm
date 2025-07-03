package main

import (
	"fmt"
	"gorm/db"
	"time"
)

type User struct {
	ID        uint   `gorm.constraints:"pk,autoincrement,unique"`
	Username  string `gorm.constraints:"unique" gorm.validators:"min(6),max(18)"`
	Email     string `gorm.constraints:"unique" gorm.validators:"email"`
	Password  string
	CreatedAt time.Time `gorm.default:"now"`
	Website   string    `gorm.validators:"url"`
}

func main() {
	dbpool := db.Connect()
	defer dbpool.Close()
	fmt.Println("[gorm] Hello, World!")

	table := db.CreateTable(User{})

	user := User{
		Username:  "hassanaziz123456789",
		Email:     "hassan123@email.com",
		Password:  "pass123",
		CreatedAt: time.Now(),
		Website:   "https://www.hassandev.me",
	}
	db.Create(table, &user)

	// filters := []db.ColumnValue{
	// 	{
	// 		Colname: "username",
	// 		Value:   "hassan",
	// 	},
	// }

	// var users []User
	// db.Filter(table, &users, filters)
	// fmt.Println(users)
}
