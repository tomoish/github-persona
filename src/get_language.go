package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

// GraphQLQuery はGraphQLのクエリを格納するための構造体です
type GraphQLQuery struct {
	Query string `json:"query"`
}

// 変更するクエリをここに入力してください
const query = `
{
  viewer {
    login
    name
  }
}
`

func main() {
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

	// GraphQLクエリを設定
	graphQLQuery := GraphQLQuery{
		Query: query,
	}
	queryBody, err := json.Marshal(graphQLQuery)
	if err != nil {
		fmt.Printf("Could not marshal query: %v\n", err)
		return
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(queryBody))
	if err != nil {
		fmt.Printf("Could not create request: %v\n", err)
		return
	}

	// ヘッダーにトークンを追加
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// レスポンスを読み取る
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %v\n", err)
		return
	}

	// レスポンスを出力
	fmt.Println("Response:")
	fmt.Println(string(body))
}