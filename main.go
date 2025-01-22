package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NameResponse represents the response structure from the names mcquay API
type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// JokeResponse represents the response structure from the joke API
type JokeResponse struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func getName() (*NameResponse, error) {
	resp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch name: %w", err)
	}
	defer resp.Body.Close()

	var nameResp NameResponse
	if err := json.NewDecoder(resp.Body).Decode(&nameResp); err != nil {
		return nil, fmt.Errorf("failed to decode name response: %w", err)
	}
	return &nameResp, nil
}

func getJoke(firstName, lastName string) (*JokeResponse, error) {
	url := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", firstName, lastName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch joke: %w", err)
	}
	defer resp.Body.Close()

	var jokeResp JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&jokeResp); err != nil {
		return nil, fmt.Errorf("failed to decode joke response: %w", err)
	}
	return &jokeResp, nil
}

// rootHandler is the entrypoint for the incoming request on /
func rootHandler(w http.ResponseWriter, r *http.Request) {
	var nameResp *NameResponse
	var jokeResp *JokeResponse
	var nameErr, jokeErr error

	nameResp, nameErr = getName()
	jokeResp, jokeErr = getJoke(nameResp.FirstName, nameResp.LastName)

	if nameErr != nil {
		http.Error(w, fmt.Sprintf("Error fetching name: %s", nameErr), http.StatusInternalServerError)
		return
	}
	if jokeErr != nil {
		http.Error(w, fmt.Sprintf("Error fetching joke: %s", jokeErr), http.StatusInternalServerError)
		return
	}

	response := jokeResp.Value.Joke

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	port := 5000
	log.Printf("Server is running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
