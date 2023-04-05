package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/slack-go/slack"
)

type RateResponse struct {
	Currency  string  `json:"currency"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
	Date      string  `json:"date"`
	Time      string  `json:"time"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/raterocket", rateRocketHandler).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":" + port, router))
}

func rateRocketHandler(w http.ResponseWriter, r *http.Request) {
	slackVerificationToken := os.Getenv("SLACK_VERIFICATION_TOKEN")
	token := r.FormValue("token")
	if token != slackVerificationToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	currency := r.FormValue("text")
	response, err := fetchRate(currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slackResponse := fmt.Sprintf("Rate for %s: %.2f (%s %s)", response.Currency, response.Price, response.Date, response.Time)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&slack.Msg{Text: slackResponse})
}

func fetchRate(currency string) (*RateResponse, error) {
	url := fmt.Sprintf("https://RateRocket.io/api/%s", currency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rateResponse RateResponse
	err = json.NewDecoder(resp.Body).Decode(&rateResponse)
	if err != nil {
		return nil, err
	}

	return &rateResponse, nil
}