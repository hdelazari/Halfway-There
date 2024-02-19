package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static"+r.URL.Path)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := r.FormValue("user")
		pass := r.FormValue("pass")

		if user == "admin" && pass == "admin" {
			http.Redirect(w, r, "/welcome.html", http.StatusSeeOther)
		}
		//Check for user and password validity
	} else if r.Method == http.MethodGet {
		r.URL.Path = "/login.html"
		viewHandler(w, r)
	}

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/view/", viewHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} else {
		log.Flags()
	}
}
