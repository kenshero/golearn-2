package routes

import (
	"encoding/json"
	"fakenews-core/news/model"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type news_struct struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetNewsRoutes(router *mux.Router) {
	router.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		news := newsmodel.GetNews()
		// fmt.Println(news)
		fmt.Fprintf(w, "news 3333333333: ", news)
	})
}

func AddNewsRoutes(router *mux.Router) {
	router.HandleFunc("/add_news", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var t news_struct
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&t)

		if err != nil {
			fmt.Println("Error")
			panic(err)
		}

		json.NewEncoder(w).Encode(t)
		// fmt.Println(t)
		// fmt.Fprintf(w, t)
		// newsmodel.AddNews()
	}).Methods("POST")
}

func DeleteNewsRoutes(router *mux.Router) {
	router.HandleFunc("/delete_news", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DeleteNewsRoutes 55555 zzz")
	})
}

func UpdateNewsRoutes(router *mux.Router) {
	router.HandleFunc("/update_news", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UpdateNewsRoutes")
	})
}

func MainNewsRoutes(router *mux.Router) {
	GetNewsRoutes(router)
	AddNewsRoutes(router)
	DeleteNewsRoutes(router)
	UpdateNewsRoutes(router)
}
