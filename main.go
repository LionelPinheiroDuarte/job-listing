package main

import (
	"fmt"
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request){
}

func main(){
	http.HandleFunc("/", home)
	fmt.Println("Server is running on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}