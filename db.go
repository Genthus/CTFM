package db

import (
	"database/sql"
	"fmt"
)

// creates master database file and tables
func CreateMasterDB(path string) error {
	db, err := sql.Open("sqlite3", "./master.db")
	if err != nil {
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
		return err
	}

	sqlStmt = `
	create table local_tags (
		tag_id integer not null primary key,
		tag_name string not null
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	sqlStmt = `
	create table file_tag_map (
		tag_id integer,
		foreign key (tag_id)
			references local_tags (tag_id)
		file_id integer,
		foreign key (file_id)
			references files (file_id)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

// verify master DB schema
func VerifyMasterDB(path string) error {
	db, err := sql.Open("sqlite3", "./master.db")
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec(".schema")
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
