package main

import (	
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"funcs" 
    "time"
    "github.com/patrickmn/go-cache"
)


//交互にトークンを取得するための関数
func getToken(tokens []string, currentIndex int) (string,int) {


	key := tokens[currentIndex]
	
	currentIndex = (currentIndex + 1) % len(tokens)
	fmt.Println(currentIndex)
	return key,currentIndex
}

func main() {


		// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// 個人アクセストークンを環境変数から取得
	// トークンをスライスに格納
	tokens := []string{
		os.Getenv("GITHUB_TOKEN1"),
		os.Getenv("GITHUB_TOKEN2"),
	}


	// トークンを取得する
   // キャッシュを作成
   c := cache.New(5*time.Minute, 10*time.Minute) // キャッシュの有効期限やクリーンアップ間隔を設定

	
    key := "kjalfu32la"

	var token string 
	var newCachedData int
	// currentIndex := 0
	



		
    
	// キャッシュにデータが存在しない場合のみデータを保存
	if _, found := c.Get(key); !found {
		
		// トークンを取得する
		token,newCachedData = getToken(tokens,0)

		c.Set(key, newCachedData, cache.DefaultExpiration)
	} else {
		// すでにデータが存在する場合の処理
		// データを取得
		cachedData,_ := c.Get(key)
		// 型アサーションを行い、int型に変換
		cachedIntData, _ := cachedData.(int)
		fmt.Println("cachedIntData")
		fmt.Println(cachedIntData)
		// トークンを取得する
		token,newCachedData = getToken(tokens,cachedIntData)

		c.Set(key, newCachedData, cache.DefaultExpiration)

		fmt.Println(token)
	}

	



    


		
    // ユーザーのリポジトリ情報を取得
    username := "kou7306"
    repos, err := funcs.GetRepositories(username, token)
	fmt.Printf("repos: %v\n", repos)
    if err != nil {
        fmt.Println(err)
        return
    }




	// 言語ごとの全体のファイルサイズを初期化
	totalLanguageSize := make(map[string]int)
	var allSize float64 = 0.0
	// 各リポジトリの言語別のファイルサイズを取得
	for _, repo := range repos {
		repoDetails, totalSize,err := funcs.GetRepositoryLanguage(repo.Name, repo.Owner, token)
		
		if err != nil {
			// エラーハンドリング
			continue
		}

		// 合計ファイルサイズを計算
		// 各言語のファイルサイズをパーセンテージで表示
		// fmt.Printf("Language sizes for repo %s:\n", repo)
		for language, size := range repoDetails {
			totalLanguageSize[language] += size
		}

		allSize += float64(totalSize)
	}

	for language, size := range totalLanguageSize {
		percentage := float64(size) / allSize * 100.0
		fmt.Printf("%s: %.2f%%\n", language, percentage)
	}

}



// func generateSessionKey() string {
//     // ランダムなバイト列を生成
//     randomBytes := make([]byte, 16)
//     _, err := rand.Read(randomBytes)
//     if err != nil {
//         panic(err)
//     }

//     // バイト列を16進数文字列に変換
//     return hex.EncodeToString(randomBytes)
// }