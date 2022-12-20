package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// creates master database file and tables
func createMasterDB(path string) error {
	db, err := sql.Open("sqlite3", path+"/master.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	sqlStmt := `
	create table files (
		file_id integer not null primary key,
		file_name text not null,
		file_hash text not null,
		file_extension text not null,
		creation text not null,
		added text not null,
		last_opened text not null,
		path text not null
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
		return err
	}

	sqlStmt = `
	create table local_tags (
		tag_id integer not null primary key,
		tag_name string not null
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
		return err
	}

	sqlStmt = `
	create table file_tag_map (
		tag_id integer,
		file_id integer,
		foreign key (tag_id)
			references local_tags (tag_id),
		foreign key (file_id)
			references files (file_id)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Successfuly created master.db")
	return nil
}

// verify master DB schema
func verifyMasterDB(path string) error {
	db, err := sql.Open("sqlite3", path+"/master.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	rows, err := db.Query("select * from sqlite_schema")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var table_type string
		var name string
		var tbl_name string
		var rootpage int
		var sql_text string
		rows.Scan(&table_type, &name, &tbl_name, &rootpage, &sql_text)

		// TODO actually verify schema integrity
		switch name {
		case "file_tag_map":
		case "files":
		case "local_tags":
		default:
			log.Fatal("unknown table found")
		}
	}

	return nil
}

func StartDB(path string) {
	_, err := os.Stat(path + "/master.db")
	if os.IsNotExist(err) {
		fmt.Println("master database not found, attempting to create...")
		createMasterDB(path)
	} else {
		fmt.Println("master.db found")
	}
	verifyMasterDB(path)
	fmt.Println("master.db schema verified")
}
