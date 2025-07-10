package handlers

import (
	"encoding/json"
	"github.com/Eiji1010/go-api/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle

	err := json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	reqArticles := []models.Article{models.Article1, models.Article2}

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	log.Println(page)

	articles := reqArticles
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

}

func ArticleDecimalHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1

	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	log.Println(articleId)
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqNice models.Nice

	if err := json.NewDecoder(req.Body).Decode(&reqNice); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(&reqNice)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(reqComment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}
