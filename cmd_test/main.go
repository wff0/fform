package main

import (
	"fform"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/ffgorm"

func main() {
	engine, _ := fform.NewEngine("mysql", dsn)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
