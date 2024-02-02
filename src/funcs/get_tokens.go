package funcs

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"github.com/joho/godotenv"
)

// 交互にトークンを取得するための関数
func GetTokens(currentIndex int) (string, int) {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", currentIndex
	}
	// 個人アクセストークンを環境変数から取得
	// トークンをスライスに格納
	// tokens := []string{
	// 	os.Getenv("GITHUB_TOKEN1"),
	// 	os.Getenv("GITHUB_TOKEN2"),
	// }
	tokens :=os.Getenv("GITHUB_TOKEN"),

	// ランダムシードの初期化
	rand.Seed(time.Now().UnixNano())


	// 元の数字をランダムな数字（0~3）で置き換え
	// currentIndex= rand.Intn(2) // 0から3のランダムな数



	// key := tokens[currentIndex]

	fmt.Println("key: ", currentIndex)

	return key, currentIndex
}
