package main

import (
	"./db"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"net/http"
	"os"
	"text/template"
)

type IndexTemplate struct {
	Message []db.Message
	Cnt     int
}

func index(c web.C, w http.ResponseWriter, r *http.Request) {
	Repo, err := db.Open(mysqlDsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot connect MySQL server")
	}

	count, err := Repo.CountMessage()
	count += 1

	if err != nil {
	}

	messages, err := Repo.RecentMessages()

	if err != nil {
	}
	for _, data := range messages {
		pp.Print(data)
	}
	t := template.Must(template.ParseFiles("./templates/top.html"))

	args := &IndexTemplate{Message: messages, Cnt: count}
	t.Execute(w, args)

}

func getPost(c web.C, w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/post.html"))
	t.Execute(w, nil)
}
func postPost(c web.C, w http.ResponseWriter, r *http.Request) {
	Repo, err := db.Open(mysqlDsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot connect MySQL server")
	}
	name := r.FormValue("name")
	body := r.FormValue("body")
	if len(name) == 0 || len(body) == 0 {
		fmt.Fprintln(w, "エラーが発生しました<br />")
		fmt.Fprintln(w, "<a href='/'>トップページへ戻る</a>")
		return
	} else {
		err = Repo.PostMessage(name, body)
		if err != nil {
			fmt.Fprintln(w, "書き込みに失敗しました")
			fmt.Fprintln(w, "<a href='/'>トップページへ戻る</a>")
		} else {
			fmt.Fprintln(w, "書き込みに成功しました<br />")
			fmt.Fprintln(w, "<a href='/'>トップページへ戻る</a>")
		}

	}
}
