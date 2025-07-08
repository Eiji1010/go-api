package main

import (
	"github.com/Eiji1010/go-api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDecimalHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("Server start at port:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
