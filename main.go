package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fahmiaz411/go-microservices/details"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Checking application health")
	response := map[string]string{
		"status": "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request){
	hostname, _ := os.Hostname()
	ip, _ := details.GetIP()
	response := map[string]string{
		"ip": ip.String(),
		"hostname": hostname,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}