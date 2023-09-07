package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Slack_name      string `json:"slack_name"`
	Current_day     string `json:"current_day"`
	Utc_time        string `json:"utc_time"`
	Track           string `json:"track"`
	Github_file_url string `json:"github_file_url"`
	Github_repo_url string `json:"github_repo_url"`
	Status_code     int64  `json:"status_code"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// validate request
	if !r.URL.Query().Has("slack_name") {
		w.WriteHeader(400)
		io.WriteString(w, "{\"error\": \"include 'slack_name' in the query string\"}")
	} else if !r.URL.Query().Has("track") {
		w.WriteHeader(400)
		io.WriteString(w, "{\"error\": \"include 'track' in the query string\"}")
	} else {

		now := time.Now()
		m := Message{
			Slack_name:      r.URL.Query().Get("slack_name"),
			Current_day:     now.Weekday().String(),
			Utc_time:        now.UTC().Format(time.RFC3339),
			Track:           r.URL.Query().Get("track"),
			Github_file_url: "https://github.com/jbrit/goone/main.go",
			Github_repo_url: "https://github.com/jbrit/goone",
			Status_code:     200,
		}
		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			fmt.Printf("error marshalling json: %s\n", err)
		}
		fmt.Printf("got /api request\n")
		io.WriteString(w, string(b))
	}
}

func main() {
	http.HandleFunc("/api", getRoot)

	fmt.Printf("listening for requests on port 80\n")
	err := http.ListenAndServe(":80", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
