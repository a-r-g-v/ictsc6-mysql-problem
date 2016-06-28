package main

import (
	"./db"
	"fmt"
	"github.com/zenazn/goji"
	"os"
)

func main() {
	mysqlDsn := os.Getenv("BOARD_MYSQLDSN")

	if mysqlDsn == "" {
		fmt.Fprintln(os.Stderr, "require enviroment variable 'BOARD_MYSQLDSN'")
		os.Exit(2)
	}
	fmt.Printf("BOARD_MYSQLDSN : %v \n", mysqlDsn)

	repo, err := db.Open(mysqlDsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot connect MySQL server")
		panic(err)
	}

	var message interface{}
	err = repo.MessageRecent(message)
	if err != nil {
		panic(err)
	}

	goji.Serve()
}
