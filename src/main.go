package main

import (
	"fmt"
	"github.com/tomoish/readme/funcs"
	"github.com/tomoish/readme/graphs"
	"log"
	"net/http"
	"strconv"
	"os"
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

	err = graphs.DrawCommitChart(dailyCommits, maxCommits, 1000, 700)
	if err != nil {
		fmt.Println(err)
	}
	http.ServeFile(w, r, "./images/commits_history.png")
}

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

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=3600")

	queryValues := r.URL.Query()
	username := queryValues.Get("username")
	if r.Method == http.MethodGet {
		// GETリクエストの処理
		// 一意の画像ファイル名の生成（例: ユーザー名とタイムスタンプを組み合わせる）
		imageFileName := fmt.Sprintf("./resultimg/result_%s.png", username)

		// 画像ファイルの存在チェック
		if _, err := os.Stat(imageFileName); os.IsNotExist(err) {
			// 画像が存在しない場合は、新たに生成

			// 画像生成の処理...
			// stats取得と画像生成
			stats := funcs.CreateUserStats(username)
			total := stats.TotalStars + stats.ContributedTo + stats.TotalIssues + stats.TotalPRs + stats.TotalCommits
			// 言語画像の生成
			language := funcs.CreateLanguageImg(username)
			//レベル、職業判定
			profession, level := funcs.JudgeRank(language, stats)
			fmt.Println("profession1: ", profession)
			//対象のキャラの画像を取得
			img := funcs.DispatchPictureBasedOnProfession(profession)
			filePath := fmt.Sprintf("characterImages/%s", img)

			// 背景画像の生成
			funcs.DrawBackground(username, "Lv."+strconv.Itoa(level), profession)
			fmt.Println("profession2: ", profession)
			// キャラクター画像の生成
			funcs.CreateCharacterImg( filePath ,"images/gauge.png", total, level)

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

			// 全て合体して画像を保存
			funcs.Merge_all("./images/background.png", "./images/stats.png", "./images/generate_character.png", "./images/language.png", "./images/commits_history.png", imageFileName)
		}

		// // キャッシュ制御ヘッダーを設定
		w.Header().Set("Cache-Control", "public, max-age=3600")

		// 生成済みの画像ファイルをクライアントに返す
		http.ServeFile(w, r, imageFileName)

	} else {
		//ファイル全削除
		// files,err := iotil.ReadDir("./resultimg")
		// if err != nil{
		// 	fmt.Println()
		// }
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	// http.HandleFunc("/test", handler)
	// http.HandleFunc("/streak", getCommitStreakHandler)
	// http.HandleFunc("/language", getLanguageHandler)
	// http.HandleFunc("/character", getCharacterHandler)
	http.HandleFunc("/history", getHistoryHandler)
	// http.HandleFunc("/user", getuserHandler)
	// http.HandleFunc("/merge", mergeAllContents)
	// http.HandleFunc("/background", getBackgroundHandler)
	http.HandleFunc("/create", createHandler)
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

// } else if r.Method == http.MethodPost {
//     // POSTリクエストの処理
// 	totalCommitContributions, totalStarredRepositories, totalIssueContributions, totalPullRequestContributions, totalRepositoryContributions, err := funcs.FetchData(username)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
