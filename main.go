package main

import (
	"./db"
	"fmt"
	"github.com/zenazn/goji"
	"os"
)

var mysqlDsn string

func main() {

	mysqlDsn = os.Getenv("BOARD_MYSQLDSN")

	if len(mysqlDsn) == 0 {
		fmt.Fprintln(os.Stderr, "require enviroment variable 'BOARD_MYSQLDSN'")
		os.Exit(2)
	}
	// fmt.Printf("BOARD_MYSQLDSN : %v \n", mysqlDsn)

	_, err := db.Open(mysqlDsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot connect MySQL server")
	}
	//	var message db.Message
	//err = repo.MessageById(&message, "1")
	//if err != nil {
	//	panic(err)
	//}
	//pp.Print(message)

	goji.Get("/", index)
	goji.Get("/post", getPost)
	goji.Post("/post", postPost)
	goji.Serve()
}
