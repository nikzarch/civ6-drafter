package main

import (
	"fmt"
	"net/http"
	"project/handlers"
	"project/middleware"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))

	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	}) // images sharing (images itself tbd)

	mux.HandleFunc("/api/sample", handlers.HandleSample)
	mux.HandleFunc("/api/nations", handlers.HandleNations)
	mux.HandleFunc("/api/drafts", handlers.HandleDrafts)
	handler := middleware.CORSFilter(mux)
	fmt.Println("Server is listening")
	http.ListenAndServe(":8080", handler)

}
