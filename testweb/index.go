package main 

import (
	// "fmt"
	"net/http"
	"github.com/gorilla/mux" 
)

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html");
	})

	http.ListenAndServe(":8080", router);   
}