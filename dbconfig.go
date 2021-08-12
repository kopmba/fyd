package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

//db in params to return db to use in controller
func dbConnect() {
	db, err := sql.Open("mysql", "user:password@tcp()127.0.0.1:3306)/fyd")
	check(err)
	defer db.Close()
}

func fetch(S *Struct, query string) {
	/*var (
		id string
		name string 
		address string
	)*/
	fields := []Value
	s := S{}
	st := reflect.TypeOf(s)
	num := st.NumField()
	for i := 0; i < num; i++ {
		fields = append(fields, st.Field[i])
	}
	rows, err := db.Query(query)
	check(err)
	defer rows.Close()
	for rows.Next() {
		//Iterate field of struct
		err := rows.Scan(&fields)
		check(err)
		fmt.Println(id, name, address)
	}
	err = rows.Err()
	check(err)
}

func fetchById(obj *Struct,  id string, query string) {
	
    stmt, err := db.Prepare(query)
	check(err)
	defer stmt.Close()
	//Iterate field of struct
	err = stmt.QueryRow(1).Scan(&id, &name, &address)
	check(err)
	fmt.Println(id, name, address)
}

func insert(query string) {
	stmt, err := db.Prepare(query)
	check(err)
	defer stmt.Close()//danger
	_, err = stmt.Exec("Dolly")
	check(err)
	lastId, err := res.LastInsertId()
	check(err)
	rowCtn, err := res.RowsAffected()
	check(err)
	fmt.Println("ID = %d, affected = %d\n", lastId, rowCtn)
	
	//using preparestmt
	/*tx, err := db.Begin()
	check(err)
	defer tx.Rollback()
	stmt, err := db.Prepare("INSERT INTO fyd(id, name, address) VALUES(?)")
	check(err)
	defer stmt.Close()//danger
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i)
		check(err)
	}
	err := tx.Commit()
	check(err)*/
}
func update(obj *Struct,  id string, query string) {
	stmt, err := db.Prepare(query)
	check(err)
	defer stmt.Close()//danger
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i)
		check(err)
	}
	check(err)
	//or use db.Exec(query)
}

func delete(obj *Struct,  id string, query string) {
	stmt, err := db.Prepare(query)
	check(err)
	defer stmt.Close()//danger
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i)
		check(err)
	}
	check(err)
	//or use db.Exec(query)
}
