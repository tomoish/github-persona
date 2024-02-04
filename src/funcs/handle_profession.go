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

	fmt.Println("final: ", finalProfession)
	return finalProfession
}
