package funcs

// import (
// 	"fmt"

// )

// func JudgeRank()[
// 		// データを取得
// 		totalCommitContributions, totalStarredRepositories, totalIssueContributions, totalPullRequestContributions, totalRepositoryContributions, err := fetchDataInTimeRange(token, username)
// 		rank := "C-"

// 		switch {
// 			case rank < 25:
// 				rank="C-"
// 			case rank < 100:
// 				rank="C"
// 			case rank < 400:
// 				rank="C+"
// 			case rank < 625:
// 				rank="B-"
// 			case rank < 1225:
// 				rank="B"
// 			case rank < 2500:
// 				rank="B+"
// 			case rank < 4900:
// 				rank="A-"
// 			case rank < 8100:
// 				rank="A"
// 			case rank < 10000:
// 				rank="A+"
// 			case 10000 <= rank:
// 				rank="S"
// 			default:
// 				rank="C-"

// 		}

// 		// 言語ごとの割合を持ってくる
// 		languages, err := funcs.CreateLanguageImg(username, token)

// 		// Percent の大きい順にソート
// 		sort.Slice(languages, func(i, j int) bool {
// 			return languages[i].Percent > languages[j].Percent
// 		})

// 		//上位２つの言語を保持
// 		topLanguage := []funcs.LanguageStat{}
// 		// 一時的に保持する
// 		temp := []funcs.LanguageStat{}

// 		// ソートされた languages を使用する
// 		// 一位がHTML,CSS,JavaScript,TypeScriptじゃない場合
// 		if languages[0].Name != "HTML" && languages[0].Name != "CSS"  && languages[0].Name != "JavaScript" && languages[0].Name != "TypeScript"{
// 			topLanguage = append(topLanguage, languages[0])

// 		} else {
// 			temp = append(temp, languages[0])
// 		}

// 		if topLanguage != [] {
// 			topLanguage = append(topLanguage, languages[1])
// 		}
// 		else{
// 			if languages[1].Name != "HTML" && languages[1].Name != "CSS"  && languages[1].Name != "JavaScript" && languages[1].Name != "TypeScript"{
// 				topLanguage = append(topLanguage, languages[1])
// 				topLanguage = append(topLanguage, temp[0])
// 			}

// 			// 上位２つの言語がHTML,CSS,JavaScript,TypeScriptの場合
// 			else{
// 				temp = append(temp, languages[1])
// 				// HTML,CSS,JavaScript,TypeScriptじゃない中の一位の言語を探す
// 				for i, language := range languages[2:] {
// 					if language.Name != "HTML" && language.Name != "CSS"  && language.Name != "JavaScript" && language.Name != "TypeScript"{
// 						if language.Percent >= 15.0{
// 							topLanguage = append(topLanguage, language)
// 							topLanguage = append(topLanguage, temp[0])
// 						}
// 					else{
// 						topLanguage = append(topLanguage, temp[0])
// 						topLanguage = append(topLanguage, temp[1])
// 					}

// 					}
// 				}
// 			}
// 		}

// 		return rank
// ]
