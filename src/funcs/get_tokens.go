package funcs

import (
	"fmt"
	"os"

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
	tokens := []string{
		os.Getenv("GITHUB_TOKEN1"),
		os.Getenv("GITHUB_TOKEN2"),
	}

	key := tokens[currentIndex]

	currentIndex = (currentIndex + 1) % len(tokens)
	fmt.Println(currentIndex)
	return key, currentIndex
}
