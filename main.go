package main

import (
	"./db"
	"fmt"
	"github.com/k0kubun/pp"
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

	messages, err := repo.RecentMessages()
	if err != nil {
		panic(err)
	}
	pp.Print(messages)

	//	var message db.Message
	//err = repo.MessageById(&message, "1")
	//if err != nil {
	//	panic(err)
	//}
	//pp.Print(message)

	goji.Serve()
}
