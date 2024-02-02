package funcs

import (
	"fmt"
	"log"
)

func CreateUserStats(username string) UserStats {
	// GitHubのアクセストークンを設定
	token, _ := GetTokens(0)
	stats := GetUserInfo(username, token)
	fmt.Println("stats: ", stats)
	ImgBytes, _ := GenerateGitHubStatsImage(stats, 600, 400)

	err := SaveImage("./images/stats.png", ImgBytes)
	if err != nil {
		// エラーが発生した場合の処理
		log.Fatal(err) // または他のエラーハンドリング方法を選択してください
	}

	return stats

}
