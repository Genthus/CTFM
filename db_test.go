package db_test

import (
	"fmt"
	"os"
	"testing"
)

func testDatabaseCreation(t *testing.T) {
	_, err := os.Stat("./master.db")
	if err != nil {
		t.Error("Failed to create master database file")
	}
}

func main() {
	fmt.Println("vim-go")
}
