package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func main(){
	log.Print("Starting the server...")
	http.HandleFunc("/",handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
			log.Printf("Listening on port %s", port)
	}

	log.Printf("Listening on port %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s!", name)
}