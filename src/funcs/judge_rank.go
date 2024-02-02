package funcs

import (
	// 他の import ステートメント
	"math"
	"sort"
)

func JudgeRank(languages []LanguageStat, stats UserStats) (string, int) {
	// データを取得

	total := stats.TotalStars + stats.ContributedTo + stats.TotalIssues + stats.TotalPRs + stats.TotalCommits

	level := int(math.Sqrt(float64(total)))

	rank := ""

	switch {
	case total < 25:
		rank = "C-"
	case total < 100:
		rank = "C"
	case total < 400:
		rank = "C+"
	case total < 625:
		rank = "B-"
	case total < 1225:
		rank = "B"
	case total < 2500:
		rank = "B+"
	case total < 4900:
		rank = "A-"
	case total < 8100:
		rank = "A"
	case total < 10000:
		rank = "A+"
	case 10000 <= total:
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
						topLanguage = append(topLanguage, language.Name)
						percentages = append(percentages, language.Percent)
						topLanguage = append(topLanguage, temp[0].Name)
						percentages = append(percentages, temp[0].Percent)
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

	return JudgeProfession(rank, topLanguage, percentages), level
}
