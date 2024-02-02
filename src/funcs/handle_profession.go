package funcs

// 職業を判定する関数
func JudgeProfession(rank string, topLanguages []string, percentages []float64) string {
	// 職業のルートと言語を定義
	magicRoute := map[string]string{
		"TypeScript": "攻撃魔法", "R": "ネクロマンサー", "Flutter": "防御魔法",
		"Go": "召喚士", "Scala": "精霊魔法", "Rust": "回復術師",
	}
	outlawRoute := map[string]string{
		"Assembly": "賞金稼ぎ", "C": "無法者", "C++": "無法者",
		"ObjectiveC": "盗賊", "Matlab": "盗賊",
	}
	warriorRoute := map[string]string{
		"C#": "ファイター", "Swift": "弓使い", "Kotlin": "弓使い",
		"Ruby": "槍使い", "PHP": "槍使い", "HTML": "剣士", "CSS": "剣士",
		"JavaScript": "剣士", "Java": "騎士", "Python": "士官",
	}

	// ランクに応じた共通の職業を判定
	switch rank {
	case "C+":
		return "冒険者見習い"
	case "C":
		return "少年"
	case "C-":
		return "赤ちゃん"
	}

	profession := "未定義の職業" // 初期値を設定

	// ランクがC+以下の場合は、基本職業をそのまま使用
	for _, language := range topLanguages {
		if p, ok := magicRoute[language]; ok {
			profession = p
			break // 最初に見つかった職業を使用
		} else if p, ok := outlawRoute[language]; ok {
			profession = p
			break // 最初に見つかった職業を使用
		} else if p, ok := warriorRoute[language]; ok {
			profession = p
			break // 最初に見つかった職業を使用
		}
	}

	// 混合職のロジックを実装
	if rank == "A-" || rank == "A" || rank == "A+" {
		if len(topLanguages) >= 2 && len(percentages) >= 2 {
			firstProfession := ""
			secondProfession := ""
			for i, language := range topLanguages[:2] { // 上位2つの言語のみを考慮
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

			// 混合職種の判定ロジック
			if firstProfession != secondProfession {
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

	return getFinalProfession(profession, rank) // 最終的な職業をランクに基づいて修正
}

func getFinalProfession(profession string, rank string) string {
	finalProfession := ""
	switch rank {
	case "B-":
		finalProfession = profession + "の見習い"
	case "B":
		finalProfession = "初級 " + profession
	case "B+":
		finalProfession = "中級 " + profession
	case "A-":
		finalProfession = "上級 " + profession
	case "A":
		finalProfession = "特級 " + profession
	case "A+":
		finalProfession = profession + "のラスボス"
	case "S":
		finalProfession = "神"
	default:
		finalProfession = profession // ランクが指定外の場合は基本職業をそのまま使用
	}
	return finalProfession
}

// func main() {
// 	fmt.Println(judgeProfession("C+", []string{"Go"}, []float64{100}))               // 召喚士
// 	fmt.Println(judgeProfession("A", []string{"Python", "Java"}, []float64{20, 20})) // 特級 士官と騎士
// 	fmt.Println(judgeProfession("C", []string{}, []float64{}))                       // 少年
// 	fmt.Println(judgeProfession("B-", []string{"TypeScript"}, []float64{100}))       // 攻撃魔法の見習い
// 	fmt.Println(judgeProfession("B", []string{"C"}, []float64{100}))                 // 初級 無法者
// 	fmt.Println(judgeProfession("B+", []string{"Java"}, []float64{100}))             // 中級 騎士
// 	fmt.Println(judgeProfession("A-", []string{"Rust"}, []float64{100}))             // 上級 回復術師
// 	fmt.Println(judgeProfession("A", []string{"Go"}, []float64{100}))                // 特級 召喚士
// 	fmt.Println(judgeProfession("A+", []string{"Assembly"}, []float64{100}))         // 賞金稼ぎのラスボス
// 	fmt.Println(judgeProfession("S", []string{"Python"}, []float64{100}))            // 神
// }
