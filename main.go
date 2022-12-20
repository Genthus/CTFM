package main

import (
	"fmt"

	db "CTFM/db"
)

func main() {

	fmt.Println("Hello")
	db.StartDB("./db")
}
