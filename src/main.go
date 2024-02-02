package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomoish/readme/funcs"
	"github.com/tomoish/readme/graphs"
)

// func handler(w http.ResponseWriter, r *http.Request) {

// 	ctx := context.Background()

// 	path := r.URL.Path
// 	segments := strings.Split(path, "/")
// 	username := segments[1]

// 	client := github.NewClient(nil)
// 	repos, _, _ := client.Repositories.ListByUser(ctx, username, nil)
// 	for _, repo := range repos {
// 		repoName := *repo.Name
// 		stars := *repo.StargazersCount
// 		forks := *repo.ForksCount

// 		fmt.Fprintf(w, "Repo: %v, Stars: %d, Forks: %d\n", repoName, stars, forks)
// 	}
// 	fmt.Fprint(w, repos)
// }

// // 言語画像生成
// func getLanguageHandler(w http.ResponseWriter, r *http.Request) {
// 	funcs.CreateLanguageImg()
// }

// //キャラ画像生成

// func getCharacterHandler(w http.ResponseWriter, r *http.Request) {
// 	funcs.CreateCharacterImg()
// }

// // 全て合体
// func mergeAllContents(w http.ResponseWriter, r *http.Request) {
// 	funcs.Merge_all("./images/background.png", "./images/stats.png", "./images/generate_character.png", "./images/language.png", "./images/commits_history.png")
// }

// // 背景生成
// func getBackgroundHandler(w http.ResponseWriter, r *http.Request) {
// 	funcs.DrawBackground("Lv.30", "神")
// }

// func getCommitStreakHandler(w http.ResponseWriter, r *http.Request) {

// 	queryValues := r.URL.Query()
// 	username := queryValues.Get("username")

// 	if username == "" {
// 		http.Error(w, "username is required", http.StatusBadRequest)
// 		return
// 	}

// 	streak, dailyCommits, _, err := funcs.GetCommitHistory(username)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprint(w, streak, dailyCommits)

// }

// func getHistoryHandler(w http.ResponseWriter, r *http.Request) {

// 	// queryValues := r.URL.Query()
// 	// username := queryValues.Get("username")

// 	// if username == "" {
// 	// 	http.Error(w, "username is required", http.StatusBadRequest)
// 	// 	return
// 	// }

// 	username := "kou7306"

// 	_, dailyCommits, maxCommits, err := funcs.GetCommitHistory(username)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = graphs.DrawCommitChart(dailyCommits, maxCommits, 1000, 700)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	http.ServeFile(w, r, "./images/commits_history.png")
// }

// func getuserHandler(w http.ResponseWriter, r *http.Request) {

// 	username := "kou7306"
// 	// GitHubのアクセストークンを設定
// 	token, _ := funcs.GetTokens(0)
// 	stats := funcs.GetUserInfo(username, token)
// 	fmt.Println("stats: ", stats)
// 	ImgBytes, _ := funcs.GenerateGitHubStatsImage(stats, 600, 400)
// 	fmt.Println("ImgBytes: ", ImgBytes)

// 	err := funcs.SaveImage("images/stats.png", ImgBytes)
// 	if err != nil {
// 		// エラーが発生した場合の処理
// 		log.Fatal(err) // または他のエラーハンドリング方法を選択してください
// 	}

// }

// 画像生成エンドポイント
func createhandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	username := queryValues.Get("username")
	// stats取得と画像生成
	stats := funcs.CreateUserStats(username)

	//レベル、職業判定

	// 背景画像の生成
	funcs.DrawBackground(username, "Lv.30", "神")

	// キャラクター画像の生成
	funcs.CreateCharacterImg("images/character.png", "images/gauge.png", stats.TotalCommits, 30)

	// 言語画像の生成
	funcs.CreateLanguageImg(username)

	// コミットカレンダー画像の生成

	_, dailyCommits, maxCommits, err := funcs.GetCommitHistory(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = graphs.DrawCommitChart(dailyCommits, maxCommits, 1000, 700)
	if err != nil {
		fmt.Println(err)
	}

	// 全て合体
	funcs.Merge_all("./images/background.png", "./images/stats.png", "./images/generate_character.png", "./images/language.png", "./images/commits_history.png")

	http.ServeFile(w, r, "./result.png")
}

func main() {
	// http.HandleFunc("/test", handler)
	// http.HandleFunc("/streak", getCommitStreakHandler)
	// http.HandleFunc("/language", getLanguageHandler)
	// http.HandleFunc("/character", getCharacterHandler)
	// http.HandleFunc("/history", getHistoryHandler)
	// http.HandleFunc("/user", getuserHandler)
	// http.HandleFunc("/merge", mergeAllContents)
	// http.HandleFunc("/background", getBackgroundHandler)
	http.HandleFunc("/create", createhandler)
	fmt.Println("Hello, World!")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
	// ImgBytes, _ := funcs.GenerateGitHubStatsImage(stats,700,500)
	// fmt.Println("ImgBytes: ", ImgBytes)

	// 画像をファイルに保存
	// err = funcs.SaveImage("images/language.png", ImgBytes)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// データを取得
	// totalCommitContributions, totalStarredRepositories, totalIssueContributions, totalPullRequestContributions, totalRepositoryContributions, err := funcs.FetchDataInTimeRange(token, username)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("totalCommitContributions: ", totalCommitContributions)
	// fmt.Println("totalStarredRepositories: ", totalStarredRepositories)
	// fmt.Println("totalIssueContributions: ", totalIssueContributions)
	// fmt.Println("totalPullRequestContributions: ", totalPullRequestContributions)
	// fmt.Println("totalRepositoryContributions: ", totalRepositoryContributions)
	fmt.Println(funcs.JudgeProfession("C+", []string{"Go"}, []float64{100}))
}
