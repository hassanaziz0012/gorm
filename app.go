package main

import (
	"fmt"
	"gorm/builder/tables"
	"gorm/db"
	"time"
)

type User struct {
	ID        uint      `gorm.constraints:"pk,autoincrement,unique"`
	Username  string    `gorm.constraints:"unique" gorm.validators:"min(6),max(18)"`
	Email     string    `gorm.constraints:"unique" gorm.validators:"email"`
	Age       int       `gorm.default:"19"`
	CreatedAt time.Time `gorm.default:"now"`
	Website   string    `gorm.validators:"url" gorm.constraints:"nullable" gorm.default:"https://example.com"`
	Approved  bool      `gorm.default:"true"`
	Password  string
}

type Task struct {
	ID     uint `gorm.constraints:"pk,autoincrement,unique"`
	Detail string
	User
}

type Project struct {
	ID    uint `gorm.constraints:"pk,autoincrement,unique"`
	Name  string
	Users []User
}

func main() {
	dbpool := db.Connect()
	defer dbpool.Close()
	fmt.Println("[gorm] Hello, World!")

	table := tables.Table(Project{}).BuildQuery().Execute()
	fmt.Println(table.Name)
}
