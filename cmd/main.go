package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	activityreporter "github.com/zavolokas/oura-data/activity-reporter"
)

const OuraToken = ""

// MetricsResponse is the response that will be sent back to Telegraf
type MetricsResponse struct {
	TrackerName string `json:"tracker_name"`
	Steps       int    `json:"steps"`
}

func getSteps(w http.ResponseWriter, r *http.Request) {

	metrics := []MetricsResponse{}
	if r.Method == "GET" {

		activityReporter := activityreporter.New(
			activityreporter.WithOuraRingData(OuraToken),
		)
		steps, err := activityReporter.GetSteps(time.Now().Add(-24*time.Hour), time.Now())
		if err != nil {
			fmt.Printf("Error getting steps - %v\n", err)
		} else {
			fmt.Printf("Amount of steps: %d\n", steps)
		}

		metrics = append(metrics, MetricsResponse{
			TrackerName: "oura",
			Steps:       steps,
		})

		resp, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		log.Printf("200 GET /steps")
		w.Write(resp)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/steps", getSteps)
	log.Fatal("Err:", http.ListenAndServe(":8010", nil))

}
