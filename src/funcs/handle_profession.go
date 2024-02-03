package funcs

import (
	"fmt"
)

// 職業を判定する関数
func JudgeProfession(rank string, topLanguages []string, percentages []float64) string {
	// 職業のルートと言語を定義
	magicRoute := map[string]string{
		"TypeScript": "攻撃魔術師", "R": "ネクロマンサー", "Dart": "防御魔術師",
		"Go": "召喚士", "Scala": "精霊魔法", "Rust": "回復術師",
	}
	outlawRoute := map[string]string{
		"Assembly": "賞金稼ぎ", "C": "犯罪者", "C++": "犯罪者",
		"ObjectiveC": "盗賊", "Matlab": "盗賊",
	}
	warriorRoute := map[string]string{
		"C#": "武闘家", "Swift": "弓使い", "Kotlin": "弓使い",
		"Ruby": "槍使い", "PHP": "槍使い", "HTML": "剣士", "CSS": "剣士",
		"JavaScript": "剣士", "Java": "騎士", "Python": "士官",
	}

	profession := "未定義の職業"
	var route string

	for _, language := range topLanguages {
		if p, ok := magicRoute[language]; ok {
			profession = p
			route = "魔法"
			break
		} else if p, ok := outlawRoute[language]; ok {
			profession = p
			route = "アウトロー"
			break
		} else if p, ok := warriorRoute[language]; ok {
			profession = p
			route = "戦士"
			break
		}
	}

	if rank == "A-" || rank == "A" || rank == "A+" {
		if len(topLanguages) >= 2 && len(percentages) >= 2 {
			firstProfession := ""
			secondProfession := ""
			for i, language := range topLanguages[:2] {
				if _, ok := magicRoute[language]; ok {
					if i == 0 {
						firstProfession = "魔法"
					} else {
						secondProfession = "魔法"
					}
				} else if _, ok := outlawRoute[language]; ok {
					if i == 0 {
						firstProfession = "アウトロー"
					} else {
						secondProfession = "アウトロー"
					}
				} else if _, ok := warriorRoute[language]; ok {
					if i == 0 {
						firstProfession = "戦士"
					} else {
						secondProfession = "戦士"
					}
				}
			}

			if firstProfession != secondProfession && percentages[0] >= 15 && (percentages[1] >= 15 && rank == "A-" || rank == "A" || rank == "A+") {
				if firstProfession == "アウトロー" && secondProfession == "戦士" {
					profession = "バーカーサ"
				} else if firstProfession == "戦士" && secondProfession == "アウトロー" {
					profession = "闇騎士"
				} else if firstProfession == "魔法" && secondProfession == "アウトロー" {
					profession = "黒魔術師"
				} else if firstProfession == "アウトロー" && secondProfession == "魔法" {
					profession = "ライダー"
				} else if firstProfession == "戦士" && secondProfession == "魔法" {
					profession = "魔法戦士"

				} else if firstProfession == "魔法" && secondProfession == "戦士" {
					profession = "魔法騎士"
				}
			}
		}
	}

	fmt.Println("pro: ", profession)
	fmt.Println("ro: ", route)
	fmt.Println("ranks ", rank)

	return getFinalProfession(profession, rank, route) // 最終的な職業をランクとルートに基づいて修正
}

// 最終的な職業名をランクとルートに基づいて修正する関数
// 最終的な職業名をランクとルートに基づいて修正する関数
func getFinalProfession(profession string, rank string, route string) string {
	finalProfession := ""
	switch rank {
	case "C-":
		finalProfession = "少年"
	case "C":
		finalProfession = "少年"
	case "C+":
		finalProfession = "冒険者見習い"
	case "B-":
		if route == "魔法" {
			finalProfession = "魔術師の見習い"
		} else if route == "アウトロー" {
			finalProfession = "不良"
		} else {
			finalProfession = "駆け出し冒険者"
		}
	case "B":
		finalProfession = "初級 " + profession
	case "B+":
		finalProfession = "中級 " + profession
	case "A-":
		finalProfession = "上級 " + profession
	case "A":
		finalProfession = "特級 " + profession
	case "A+":
		switch profession {
		case "賞金稼ぎ", "犯罪者", "盗賊":
			finalProfession = "裏社会のボス"
		case "攻撃魔術師", "防御魔術師", "召喚士", "精霊魔法", "回復術師":
			finalProfession = "魔法帝"
		case "武闘家", "弓使い", "槍使い", "剣士":
			finalProfession = "勇者"
		case "騎士", "士官":
			finalProfession = "騎士団長"
		case "魔法戦士", "魔法騎士":
			finalProfession = "賢者"
		case "バーカーサ", "闇騎士":
			finalProfession = "サイコパス"
		default:
			finalProfession = "魔王"
		}
	case "S":
		finalProfession = "神"
	default:
		finalProfession = profession // ランクが指定外の場合は基本職業をそのまま使用
	}
	return finalProfession
}

// func main() {
// 	fmt.Println(judgeProfession("C+", []string{"Go"}, []float64{100}))               // 召喚士
// 	fmt.Println(judgeProfession("A", []string{"Python", "Java"}, []float64{20, 20})) // 特級 士官
// 	fmt.Println(judgeProfession("C", []string{}, []float64{}))                       // 少年
// 	fmt.Println(judgeProfession("B-", []string{"TypeScript"}, []float64{100}))       // 攻撃魔術師の見習い
// 	fmt.Println(judgeProfession("B", []string{"C"}, []float64{100}))                 // 初級 犯罪者
// 	fmt.Println(judgeProfession("B+", []string{"Java"}, []float64{100}))             // 中級 騎士
// 	fmt.Println(judgeProfession("A-", []string{"Rust"}, []float64{100}))             // 上級 回復術師
// 	fmt.Println(judgeProfession("A", []string{"Go"}, []float64{100}))                // 特級 召喚士
// 	fmt.Println(judgeProfession("A+", []string{"Assembly"}, []float64{100}))         // 賞金稼ぎのラスボス
// 	fmt.Println(judgeProfession("S", []string{"Python"}, []float64{100}))            // 神
// }

// func dispatchPictureBasedOnProfession(profession string) string {
//     rankToPrefix := map[string]string{
//         "初級":   "b",
//         "中級":   "b+",
//         "上級":   "a-",
//         "特級":   "a",
//     }
//     professionToCode := map[string]int{
//         "攻撃魔術師": 1, "ネクロマンサー": 2, "防御魔術師": 3,
//         "召喚士": 4, "精霊魔法": 5, "回復術師": 6,
//         "賞金稼ぎ": 7, "犯罪者": 8, "盗賊": 9,
//         "武闘家": 10, "弓使い": 11, "槍使い": 12,
//         "剣士": 13, "騎士": 14, "士官": 15,
//         "魔法使い": 16, // 魔法ルート
//         // アウトローと戦士ルートは混合職を考慮してデフォルトで扱う
//     }
//     routeToCode := map[string]int{
//         "魔法": 1,
//         "アウトロー": 2,
//         "戦士": 3,
//     }

//     // 職業名からランクと基本の職業を抽出
//     var baseProfession, route string
//     profession = strings.Replace(profession, "の魔法使い", "", -1) // "の魔法使い"を削除
//     for r, prefix := range rankToPrefix {
//         if strings.Contains(profession, r) {
//             rank = prefix
//             baseProfession = strings.Replace(profession, r+" ", "", 1)
//             break
//         }
//     }

//     // ランクが見習いまたはラスボスで、かつルートが特定できない場合の処理
//     if rank == "b-" || rank == "a+" {
//         if strings.Contains(profession, "魔法") {
//             route = "魔法"
//         } else if strings.Contains(profession, "アウトロー") {
//             route = "アウトロー"
//         } else if strings.Contains(profession, "戦士") {
//             route = "戦士"
//         }
//     }

//     // 職業コードを取得
//     code, exists := professionToCode[strings.TrimSpace(baseProfession)]
//     if !exists {
//         // ルートに基づいたコードを取得
//         code, exists = routeToCode[route]
//         if !exists {
//             return "default.png" // 職業がマップにない場合はデフォルト画像を返す
//         }
//     }

//     // 画像ファイル名を生成
//     filename := fmt.Sprintf("%s%d.png", rank, code)
//     return filename
// }
