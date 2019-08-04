package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:1234@/keeplearning")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//insert
	// stmt, err := db.Prepare("insert into products (code, price) values(?, ?)")
	// stmt.Exec("Mouse", 500)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// delete
	// stmt, err := db.Prepare("delete from products where id=?")
	// stmt.Exec(4)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// update
	stmt, err := db.Prepare("update products set code=? where id=?")
	stmt.Exec("Keyboard", 2)
	if err != nil {
		fmt.Println(err)
	}

	//query
	rows, err := db.Query("select id, code, price from products order by price DESC")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var code string
		var price string
		err = rows.Scan(&id, &code, &price)
		fmt.Printf("id: %d code: %s price: %s \n", id, code, price)
	}
}
