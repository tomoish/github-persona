package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v58/github"

	"github.com/tomoish/readme/funcs"
	"github.com/tomoish/readme/graphs"
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
	funcs.CreateLanguageImg()
}

func getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	funcs.CreateCharacterImg()
}

func getCommitStreakHandler(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
	username := queryValues.Get("username")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	streak, dailyCommits, _, err := funcs.GetCommitHistory(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, streak, dailyCommits)

}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
	username := queryValues.Get("username")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	_, dailyCommits, maxCommits, err := funcs.GetCommitHistory(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = graphs.DrawCommitChart(dailyCommits, maxCommits, 1000, 600)
	if err != nil {
		fmt.Println(err)
	}
	http.ServeFile(w, r, "./images/commits_history.png")
}

func main() {
	http.HandleFunc("/test", handler)
	http.HandleFunc("/streak", getCommitStreakHandler)
	http.HandleFunc("/language", getLanguageHandler)
	http.HandleFunc("/character", getCharacterHandler)
	http.HandleFunc("/history", getHistoryHandler)
	fmt.Println("Hello, World!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}

	username := "kou7306"

	// GitHubのアクセストークンを設定
	token, _ := funcs.GetTokens(0)

	// データを取得
	totalCommitContributions, totalStarredRepositories, totalIssueContributions, totalPullRequestContributions, totalRepositoryContributions, err := funcs.FetchDataInTimeRange(token, username)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("totalCommitContributions: ", totalCommitContributions)
	fmt.Println("totalStarredRepositories: ", totalStarredRepositories)
	fmt.Println("totalIssueContributions: ", totalIssueContributions)
	fmt.Println("totalPullRequestContributions: ", totalPullRequestContributions)
	fmt.Println("totalRepositoryContributions: ", totalRepositoryContributions)

}
