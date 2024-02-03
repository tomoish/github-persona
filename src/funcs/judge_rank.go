package funcs

import (
	// 他の import ステートメント
	"fmt"
	"math"
	"sort"
)

func JudgeRank(languages []LanguageStat, stats UserStats,star int) (string, int) {
	// データを取得
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
	total := star + stats.ContributedTo + stats.TotalIssues + stats.TotalPRs + stats.TotalCommits

	// レベルを計算コントリビューション5000でレベル100
	level := int(math.Sqrt(float64(total))*2.24)
	if level > 100 {
		level = 100
	}

	rank := ""

	switch {
	case total < 25:
		rank = "C-"
	case total < 75:
		rank = "C"
	case total < 175:
		rank = "C+"
	case total < 320:
		rank = "B-"
	case total < 525:
		rank = "B"
	case total < 700:
		rank = "B+"
	case total < 900:
		rank = "A-"
	case total < 1200:
		rank = "A"
	case total < 1550:
		rank = "A+"
	case 2000 <= total:
		rank = "S"
	default:
		rank = "C-"

	}

	// Percent の大きい順にソート
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Percent > languages[j].Percent
	})

	//上位２つの言語を保持
	topLanguage := []string{}
	// パーセントを保持
	percentages := []float64{}
	// 一時的に保持する
	temp := []LanguageStat{}

	// ソートされた languages を使用する
	// 一位がHTML,CSS,JavaScript,TypeScriptじゃない場合
	if languages[0].Name != "HTML" && languages[0].Name != "CSS" && languages[0].Name != "JavaScript" && languages[0].Name != "TypeScript" {
		topLanguage = append(topLanguage, languages[0].Name)
		percentages = append(percentages, languages[0].Percent)

	} else {
		temp = append(temp, languages[0])
	}

	if len(topLanguage) != 0 {
		topLanguage = append(topLanguage, languages[1].Name)
		percentages = append(percentages, languages[1].Percent)
	} else {

		if languages[1].Name != "HTML" && languages[1].Name != "CSS" && languages[1].Name != "JavaScript" && languages[1].Name != "TypeScript" {
			topLanguage = append(topLanguage, languages[1].Name)
			percentages = append(percentages, languages[1].Percent)
			topLanguage = append(topLanguage, temp[0].Name)
			percentages = append(percentages, temp[0].Percent)
			topLanguage = append(topLanguage, temp[0].Name)
			percentages = append(percentages, temp[0].Percent)
		} else {
			// 上位２つの言語がHTML,CSS,JavaScript,TypeScriptの場合

			temp = append(temp, languages[1])
			// HTML,CSS,JavaScript,TypeScriptじゃない中の一位の言語を探す
			for _, language := range languages[2:] {
				if language.Name != "HTML" && language.Name != "CSS" && language.Name != "JavaScript" && language.Name != "TypeScript" {
					if language.Percent >= 15.0 {
						if _, exists := colordict[language.Name]; exists {
							topLanguage = append(topLanguage, language.Name)
							percentages = append(percentages, language.Percent)
							topLanguage = append(topLanguage, temp[0].Name)
							percentages = append(percentages, temp[0].Percent)
						}
					} else {
						topLanguage = append(topLanguage, temp[0].Name)
						topLanguage = append(topLanguage, temp[1].Name)
						percentages = append(percentages, temp[0].Percent)
						percentages = append(percentages, temp[1].Percent)
					}
				}
			}
		}
	}

	fmt.Println("Top Language: ", topLanguage)
	fmt.Println("Percentages: ", percentages)
	fmt.Println("Rank: ", rank)
	//ここまでいけてる　未定義のものがよくない
	return JudgeProfession(rank, topLanguage, percentages), level
}
