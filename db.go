package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB() {
	db, _ = sql.Open("sqlite3", "./bogo.db")

	statement, _ := db.Prepare("DROP TABLE IF EXISTS thing")
	statement.Exec()

	var err error
	statement, err = db.Prepare("CREATE TABLE thing (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name VARCHAR(200), type VARCHAR(200))")
	if err != nil {
		log.Fatalf("CREATE TABLE thing Prepare: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("CREATE TABLE thing Exec: %v", err)
	}
}

type Thing struct {
	Id   int64
	Name string
	Type string
}

func GetThing(id int64) (Thing, error) {
 	var t Thing
 	row := db.QueryRow("SELECT id, name, type FROM thing WHERE id = ?", id)
 	err := row.Scan(&t.Id, &t.Name, &t.Type)
 	return t, err
}

func InsertThing(t Thing) int64 {
	statement, err := db.Prepare("INSERT INTO thing (name, type) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("InsertThing db.Prepare: %v", err)
	}
	res, err := statement.Exec(t.Name, t.Type)
	if err != nil {
		log.Fatalf("InsertThing Exec: %v", err)
	}

	lastInsertId, _ := res.LastInsertId()
	rowsAffected,_ := res.RowsAffected()
	log.Printf("res: %v, %v", lastInsertId, rowsAffected)
	if err != nil {
		log.Fatalf("InsertThing statement.Exec: %v", err)
	}

	return lastInsertId
}

func ListThings() (ts []Thing) {
 	rows, err := db.Query("SELECT id, name, type FROM thing ORDER BY id")
 	defer rows.Close()
	if err != nil {
		log.Fatalf("GetAllThings db.Query: %v", err)
	}
 	for rows.Next() {
 		var t Thing
 		err = rows.Scan(&t.Id, &t.Name, &t.Type)
 		if err != nil {
 			log.Fatalf("GetAllThings rows.Scan: %v", err)
 		}
 		ts = append(ts, t)
 	}

 	return ts
}
