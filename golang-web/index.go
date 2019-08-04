package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"time"
)

type Product struct {
	Name  string
	Price int
}

type Cooke struct {
	Name    string
	Value   string
	Expires time.Time
}

func main() {
	// var templates = template.Must(template.ParseFiles("index.html"))
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// expiration := time.Now().Add(time.Hour * 24 * 365)
		// cookie := http.Cookie{Name: "user", Value: "kenshero", Expires: expiration}
		// http.SetCookie(w, &cookie)
		// fmt.Fprintf(w, "Create Cookie")
		// myProduct := Product{"Milk", 500}
		// templates.ExecuteTemplate(w, "index.html", myProduct)
	})

	router.HandleFunc("/login", login)
	router.HandleFunc("/upload", upload)
	router.HandleFunc("/uploading", uploadHandler)

	// router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {

	// 	vars := mux.Vars(r)
	// 	name := vars["name"]
	// 	age := userDB[name]
	// 	fmt.Fprintf(w, "My Name is %s %d year old", name, age)

	// }).Methods("GET")
	fmt.Println("server run")
	http.ListenAndServe(":8080", router)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method", r.Method)
	r.ParseForm()
	fmt.Println("Username", r.Form["username"])
	fmt.Println("Password", r.Form["password"])
	http.ServeFile(w, r, "login.html")
}

func upload(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "upload.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v", handle.Header)
	f, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w, "Upload Complete")
}
