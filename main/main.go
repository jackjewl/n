package main

import (
	"fmt"

	"github.com/jackjewl/squirrelBlog/db"
)

func main() {
	db.InitDb()
	fmt.Println("db")
	fmt.Println(db.DB)
	fmt.Println("database connect success,")
	fmt.Println(db.DB)

}
