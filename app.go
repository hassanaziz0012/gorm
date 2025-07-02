package main

import (
	"fmt"
	"gorm/db"
	"time"
)

type User struct {
	ID        uint   `gorm.constraints:"pk,autoincrement,unique"`
	Username  string `gorm.constraints:"unique"`
	Email     string `gorm.constraints:"unique"`
	Password  string
	CreatedAt time.Time `gorm.default:"now"`
}

func main() {
	dbpool := db.Connect()
	defer dbpool.Close()
	fmt.Println("[gorm] Hello, World!")

	db.CreateTable(User{})

	// user := User{
	// 	Username:  "hassan123",
	// 	Email:     "hassan123@email.com",
	// 	Password:  "pass123",
	// 	CreatedAt: time.Now(),
	// }
	// db.Create(table, &user)

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
