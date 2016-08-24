package main

import (
	"./db"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
)

type IndexTemplate struct {
	Message []db.Message
}

type MessageTemplate struct {
	Message string
}

func index(c web.C, w http.ResponseWriter, r *http.Request) {
	Repo, err := db.Open(mysqlDsn)
	if err != nil {
		renderMessage("DBへの接続に失敗しました", w)
		return
	}

	messages, err := Repo.RecentMessages()

	if err != nil {
		renderMessage("エラーが発生しました", w)
		return
	}
	t := template.Must(template.ParseFiles("./templates/top.html"))

	args := &IndexTemplate{Message: messages}
	t.Execute(w, args)

}

func getPost(c web.C, w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/post.html"))
	t.Execute(w, nil)
}
func postPost(c web.C, w http.ResponseWriter, r *http.Request) {
	Repo, err := db.Open(mysqlDsn)
	if err != nil {
		renderMessage("DBへの接続に失敗しました", w)
		return
	}
	name := r.FormValue("name")
	body := r.FormValue("body")
	if len(name) == 0 || len(body) == 0 {
		renderMessage("名前欄か，内容欄が空白です", w)
		return
	}
	err = Repo.PostMessage(name, body)
	if err != nil {
		renderMessage("エラーが発生しました", w)
		return
	}
	renderMessage("書き込みに成功しました", w)
	return

}

func renderMessage(message string, w http.ResponseWriter) {
	t := template.Must(template.ParseFiles("./templates/message.html"))
	args := &MessageTemplate{Message: message}
	t.Execute(w, args)
}
