package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v58/github"

	"github.com/tomoish/readme/funcs"
)

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	path := r.URL.Path
	segments := strings.Split(path, "/")
	username := segments[1]

	client := github.NewClient(nil)
	repos, _, _ := client.Repositories.ListByUser(ctx, username, nil)
	for _, repo := range repos {
		repoName := *repo.Name
		stars := *repo.StargazersCount
		forks := *repo.ForksCount

		fmt.Fprintf(w, "Repo: %v, Stars: %d, Forks: %d\n", repoName, stars, forks)
	}
	fmt.Fprint(w, repos)
}

func getLanguageHandler(w http.ResponseWriter, r *http.Request) {
	CreateLanguageImg()
}

func getCommitStreakHandler(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
	username := queryValues.Get("username")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	streak, err := funcs.GetLongestStreak(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, streak)

}

func main() {
	http.HandleFunc("/test", handler)
	http.HandleFunc("/streak", getCommitStreakHandler)
	http.HandleFunc("/language", getLanguageHandler)
	fmt.Println("Hello, World!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}

}
