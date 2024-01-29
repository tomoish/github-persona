package main

import (
	"fmt"
	"os"

	"github.com/tomoish/readme/funcs"

	"github.com/joho/godotenv"
)

func main2() {

	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// 個人アクセストークンを環境変数から取得
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("GITHUB_TOKEN is not set")
		return
	}

	// ユーザーのリポジトリ情報を取得
	username := "kou7306"
	repos, err := funcs.GetRepositories(username, token)
	fmt.Printf("repos: %v\n", repos)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 各リポジトリの言語別のファイルサイズを取得
	for _, repo := range repos {
		repoDetails, totalSize, err := funcs.GetRepositoryLanguage(repo.Name, repo.Owner, token)

		if err != nil {
			// エラーハンドリング
			continue
		}

		// 合計ファイルサイズを計算
		// 各言語のファイルサイズをパーセンテージで表示
		fmt.Printf("Language sizes for repo %s:\n", repo.Name)
		for language, size := range repoDetails {
			percentage := float64(size) / float64(totalSize) * 100.0
			fmt.Printf("%s: %.2f%%\n", language, percentage)
		}
	}

}
