package main

import (
	"fmt"

	"github.com/jackjewl/squirrelBlog/db"
)

func main() {
	db.InitDb()
	fmt.Println(db.DB)
}
