package funcs

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func CreateLanguageImg(username string) []LanguageStat {
	// 言語ごとの色をここで決める
	colordict := map[string]string{
		"HTML":        "#ff0000",
		"CSS":         "#ffa500",
		"Python":      "#000080",
		"JavaScript":  "#ffff00",
		"TypeScript":  "#3cb371",
		"R":           "#9932cc",
		"Go":          "#87cefa",
		"Scala":       "##006400",
		"Dart":        "#4169e1",
		"Rust":        "#696969",
		"assembly":    "#ffd700",
		"C":           "#f0e68c",
		"C++":         "#ff69b4",
		"Objective-C": "#a52a2a",
		"Matlab":      "#ff6347",
		"C#":          "#800080",
		"Swift":       "#800000",
		"Kotlin":      "#bdb76b",
		"Ruby":        "#ee82ee",
		"PHP":         "#808000",
		"Java":        "#daa520",
		"others":      "#000000",
	}

	// キャッシュのキーを設定
	key := "kjalfu32la"

	var token string
	var newCachedData int
	// currentIndex := 0
	// キャッシュを作成
	c := cache.New(5*time.Minute, 10*time.Minute) // キャッシュの有効期限やクリーンアップ間隔を設定
	// Githubトークンの交互取得
	// キャッシュにデータが存在しない場合のみデータを保存
	if _, found := c.Get(key); !found {
		// トークンを取得する
		token, newCachedData = GetTokens(0)

		c.Set(key, newCachedData, cache.DefaultExpiration)
	} else {
		// すでにデータが存在する場合の処理
		// データを取得
		cachedData, _ := c.Get(key)
		// 型アサーションを行い、int型に変換
		cachedIntData, _ := cachedData.(int)

		// トークンを取得する
		token, newCachedData = GetTokens(cachedIntData)

		c.Set(key, newCachedData, cache.DefaultExpiration)

		fmt.Println(token)
	}

	// ユーザーのリポジトリ情報を取得

	repos, err := GetRepositories(username, token)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 言語ごとの全体のファイルサイズを初期化
	totalLanguageSize := make(map[string]int)
	var allSize float64 = 0.0

	// 各リポジトリの言語別のファイルサイズを取得
	for _, repo := range repos {
		repoDetails, totalSize, err := GetRepositoryLanguage(repo.Name, repo.Owner, token)

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
	colorCode := ""

	languages := []LanguageStat{}
	// 各言語のファイルサイズをパーセンテージで計算
	for language, size := range totalLanguageSize {
		percentage := float64(size) / allSize * 100.0
		fmt.Printf("%s: %.2f%%\n", language, percentage)

		if _, ok := colordict[language]; ok {
			// language キーが colordict マップに存在する場合の処理
			colorCode = colordict[language]
		} else {
			// language キーが colordict マップに存在しない場合の処理
			colorCode = colordict["others"]
		}

		languages = append(languages, LanguageStat{
			Name:    language,
			Percent: percentage,
			Color:   colorCode,
		})
	}

	ImgBytes, _ := GenerateLanguageUsageGraph(languages, 600, 400)

	// 画像をファイルに保存
	err = SaveImage("./images/language.png", ImgBytes)
	if err != nil {
		fmt.Println(err)

	}

	return languages

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
