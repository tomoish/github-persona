package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"github.com/google/go-github/v58/github"
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

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Hello, World!!")
	http.ListenAndServe(":8080", nil)
}
