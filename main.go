package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"encoding/json"
)

type Offer struct {
	ID        int           `json:"id"`
	Company   string        `json:"company"`
	Logo      string        `json:"logo"`
	New       bool          `json:"new"`
	Featured  bool          `json:"featured"`
	Position  string        `json:"position"`
	Role      string        `json:"role"`
	Level     string        `json:"level"`
	PostedAt  string        `json:"postedAt"`
	Contract  string        `json:"contract"`
	Location  string        `json:"location"`
	Languages []string      `json:"languages"`
	Tools     []string `json:"tools"`
}


func data(){
	offerdata, err := os.Open("./data.json")
	if err != nil{
		fmt.Println("Error opening the offerdata %s", offerdata)
	}

	var offerDecoder *json.Decoder = json.NewDecoder(offerdata)
    if err != nil {
        log.Fatal(err)
    }

	var joboffer []Offer

	err = offerDecoder.Decode(&joboffer)
    if err != nil {
        log.Fatal(err)
    }

	for _, offer := range joboffer {
        fmt.Println("offer name:", offer.Company)
    }
}

func home(w http.ResponseWriter, r *http.Request){
	data()
}

func main(){
	http.HandleFunc("/", home)
	fmt.Println("Server is running on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}