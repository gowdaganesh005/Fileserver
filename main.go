package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./simple_server"))
	http.Handle("/", fileserver)
	http.HandleFunc("/about", abouthandler)
	fmt.Println("Server starting on host 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server could not run properly")
	}

}
func abouthandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about.html")

}
