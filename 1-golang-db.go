package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// tabelの作成
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// データの投入
	/*
		cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
		_, err = DbConnection.Exec(cmd, "Nancy", 20)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// データの更新
	/*
		cmd = "UPDATE person Set age = ? WHERE name =?"
		_, err = DbConnection.Exec(cmd, 25, "Mike")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// マルチプルセレクト
	/*
		cmd = "SELECT * FROM person"
		rows, _ := DbConnection.Query(cmd)
		defer rows.Close()
		var pp []Person   //スライスを用意してここにDBの中身を入れていく
		for rows.Next() { //繰り返し処理
			var p Person
			err := rows.Scan(&p.Name, &p.Age)
			if err != nil {
				log.Fatalln(err)
			}
			pp = append(pp, p)
			fmt.Println(p)
		}
		fmt.Println(pp)
		for _, p := range pp {
			fmt.Println(p.Name, p.Age)
		}
	*/

	//シングルセレクト
	/*
		cmd = "SELECT * FROM person where age = ?"
		row := DbConnection.QueryRow(cmd, 20)
		var ppp Person
		err = row.Scan(&ppp.Name, &ppp.Age)
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
		fmt.Println(ppp.Name, ppp.Age)
	*/

	// 削除
	/*
		cmd = "DELETE FROM person WHERE name = ?"
		_, err = DbConnection.Exec(cmd, "Nancy")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person   //スライスを用意してここにDBの中身を入れていく
	for rows.Next() { //繰り返し処理
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
