package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"encoding/json"
	"html/template"
	"time"

	"github.com/prometheus/client_golang/prometheus"
    	"github.com/prometheus/client_golang/prometheus/promhttp"
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

var (
    
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "job_listing_http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "job_listing_http_request_duration_seconds",
            Help: "Duration of HTTP requests in seconds",
        },
        []string{"method", "endpoint"},
    )

    jobsDisplayed = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "job_listing_jobs_displayed_total",
            Help: "Total number of jobs displayed to users",
        },
    )
)

func init() {
    // Enregistrer les métriques
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)
    prometheus.MustRegister(jobsDisplayed)

}
func healthcheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusOK)
    
   	 response := map[string]string{
        	"status": "healthy",
        	"service": "job-listing-app",
        	"timestamp": time.Now().Format(time.RFC3339),
    	}
    
    	json.NewEncoder(w).Encode(response)
}

func loadJobOffers()(joboffers []Offer){
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
	start := time.Now()
	httpRequestsTotal.WithLabelValues(r.Method, "/").Inc()
	
	joboffers := loadJobOffers()

	jobsDisplayed.Add(float64(len(joboffers)))

	pageVariables := PageVariables{
		PageTitle: "JobOfferslist",
		PageOffers: joboffers,
	}

	t, err := template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("template parsing error:", err)
		return
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
        	log.Print("template execution error:", err)
    	}
	
	duration := time.Since(start).Seconds()
    	httpRequestDuration.WithLabelValues(r.Method, "/").Observe(duration)
}

func main(){
	http.HandleFunc("/", getOffers)
	http.HandleFunc("/health", healthcheck)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server is running on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
