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
	now := time.Now()
	m := Message{
		Slack_name:      "jibs",
		Current_day:     now.Weekday().String(),
		Utc_time:        now.UTC().String(),
		Track:           "backend",
		Github_file_url: "https://github.com/jbrit/goone/main.go",
		Github_repo_url: "https://github.com/jbrit/goone",
		Status_code:     200,
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Printf("error marshalling json: %s\n", err)
	}
	fmt.Printf("got / request\n")
	io.WriteString(w, string(b))
}

func main() {
	http.HandleFunc("/", getRoot)

	fmt.Printf("listening for requests on port 3333\n")
	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
