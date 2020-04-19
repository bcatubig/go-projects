package main

import (
	"github.com/russross/blackfriday"
	"net/http"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	w.Write(markdown)
}
