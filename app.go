package main

import (
	"fmt"
	"gorm/builder"
	"gorm/db"
	"gorm/tables"
	"log"
	"time"
)

type User struct {
	ID        uint      `gorm.constraints:"pk,autoincrement,unique"`
	Username  string    `gorm.constraints:"unique" gorm.validators:"min(6),max(18)"`
	Email     string    `gorm.constraints:"unique" gorm.validators:"email"`
	CreatedAt time.Time `gorm.default:"now"`
	Website   string    `gorm.validators:"url"`
	Approved  bool      `gorm.default:"true"`
	Password  string
}

func main() {
	dbpool := db.Connect()
	defer dbpool.Close()
	fmt.Println("[gorm] Hello, World!")

	table := tables.CreateTable(User{})

	filters := map[string]any{
		"password": "reused123",
	}
	items, err := builder.
		Query(table.Model).
		Select().
		Where(filters).
		Build().
		Execute()

	for _, item := range items {
		fmt.Println(item.Username, item.Password)
	}

	if err != nil {
		log.Fatal("query error: ", err)
	}

}
