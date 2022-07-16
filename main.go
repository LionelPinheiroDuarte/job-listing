package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"encoding/json"
	"html/template"
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
//Page variables
type PageVariables struct{
	PageTitle string
	PageOffers []Offer
}


func data()(joboffers []Offer){
	offerdata, err := os.Open("./data.json")
	if err != nil{
		fmt.Println("Error opening the offerdata %s", offerdata)
	}

	var offerDecoder *json.Decoder = json.NewDecoder(offerdata)
    if err != nil {
        log.Fatal(err)
    }

	err = offerDecoder.Decode(&joboffers)
    if err != nil {
        log.Fatal(err)
    }

	return joboffers
}

func getOffers(w http.ResponseWriter, r *http.Request){
	
	joboffers := data()
	pageVariables := PageVariables{
		PageTitle: "JobOfferslist",
		PageOffers: joboffers,
	}

	t, err := template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("template parsing error:", err)
	}

	err = t.Execute(w, pageVariables)
}

func home(w http.ResponseWriter, r *http.Request){
	data()
}

func main(){
	http.HandleFunc("/", getOffers)
	fmt.Println("Server is running on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}