package routers

import (
	"net/http"

	"github.com/an-0305/study-go-api/controllers"
	"github.com/gorilla/mux"
)


func NewRouter(con *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	return r
}
