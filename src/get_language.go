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
	user(login: "kou7306") {
		repositories(first: 100, privacy: PUBLIC, orderBy: {field: NAME, direction: ASC}) {
			nodes {
				name
				languages(first: 10, orderBy: {field: SIZE, direction: DESC}) {
					edges {
						node {
							name
						}
						size
					}
				}
			}
		}
	}
}
`

// レスポンスの型を定義
type ApiResponse struct {
	Data struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name string
					Languages struct {
						Edges []struct {
							Node struct {
								Name string
							}
							Size int
						}
					}
				}
			}
		}
	}
}
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
	queryBody, err := json.Marshal(graphQLQuery) //構造体をJSONに変換
	if err != nil {
		fmt.Printf("Could not marshal query: %v\n", err)
		return
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(queryBody)) //バッファに書き込むためにbytes.NewBufferを使用してデータ変換
	if err != nil {
		fmt.Printf("Could not create request: %v\n", err)
		return
	}

	// ヘッダーにトークンを追加
	req.Header.Set("Authorization", "bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req) //Doメソッドでリクエストを送信
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		return
	}
	defer resp.Body.Close() //リソースの解放

	// レスポンスを読み取る
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %v\n", err)
		return
	}

	var response ApiResponse //あらかじめ定義した構造体にレスポンスを格納
	err = json.Unmarshal(body, &response) //JSONを構造体に変換	
	if err != nil {
		fmt.Printf("Could not unmarshal response: %v\n", err)
		return
	}
	// 新しいレスポンスの解析と集計ロジック
	languageSizes := make(map[string]int)
	var totalSize int

	for _, repo := range response.Data.User.Repositories.Nodes {
		for _, languageEdge := range repo.Languages.Edges {
			languageSizes[languageEdge.Node.Name] += languageEdge.Size //言語ごとのサイズを集計
			totalSize += languageEdge.Size //全体のサイズを集計
		}
	}

	fmt.Println("\nLanguage Usage:")
	for lang, size := range languageSizes {
		percentage := float64(size) / float64(totalSize) * 100 //言語ごとのサイズを全体のサイズで割ってパーセンテージを計算
		fmt.Printf("%s: %.2f%%\n", lang, percentage)
	}
	// languageCounts := make(map[string]int)
	// var totalCount int
	
	// for _, node := range response.Data.User.Repositories.Nodes {
	// 	if node.PrimaryLanguage.Name != "" {
	// 		languageCounts[node.PrimaryLanguage.Name]++
	// 		totalCount++
	// 	}
	// }
	
	// fmt.Println("\nLanguage Usage:")
	// for lang, count := range languageCounts {
	// 	percentage := float64(count) / float64(totalCount) * 100
	// 	fmt.Printf("%s: %.2f%%\n", lang, percentage)
	// }
}