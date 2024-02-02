package main

import (
	"fmt"
	"strings"
)

func dispatchPictureBasedOnProfession(profession string) string {
	// 特定の職業または状態に応じた画像ファイル名を返す
	switch profession {
	case "赤ちゃん":
		return "c-.png"
	case "少年":
		return "c.png"
	case "冒険者見習い":
		return "c+.png"
	case "魔法使いの見習い":
		return "c+1.png"
	case "アウトローの見習い":
		return "c+2.png"
	case "戦士の見習い":
		return "c+3.png"
	case "魔法使いのラスボス":
		return "a+1.png"
	case "アウトローのラスボス":
		return "a+2.png"
	case "バーカーサ", "闇騎士", "黒魔術師", "ライダー", "魔法戦士", "魔法騎士":
		return "c+.png"
	case "神":
		return "s.png"
	}

	// 職業名からランクと基本の職業を抽出
	rankToPrefix := map[string]string{
		"初級": "b",
		"中級": "b+",
		"上級": "a-",
		"特級": "a",
	}
	professionToCode := map[string]int{
		"攻撃魔法": 1, "ネクロマンサー": 2, "防御魔法": 3,
		"召喚師": 4, "精霊魔法": 5, "回復術師": 6,
		"賞金稼ぎ": 7, "無法者": 8, "盗賊": 9,
		"ファイター": 10, "弓使い": 11, "槍使い": 12,
		"剣士": 13, "騎士": 14, "士官": 15,
	}

	parts := strings.Split(profession, " ")
	rankPart, professionPart := "", ""
	if len(parts) > 1 {
		rankPart = parts[0]
		professionPart = strings.Join(parts[1:], " ")
	} else {
		professionPart = parts[0]
	}

	// 職業コードを取得
	code, exists := professionToCode[professionPart]
	if !exists {
		return "default.png" // 職業がマップに存在しない場合はデフォルト画像を返す
	}

	// ランクに基づいてプレフィックスを決定
	prefix, exists := rankToPrefix[rankPart]
	if !exists {
		// ランク情報がない場合は、職業コードのみで画像ファイル名を決定
		return fmt.Sprintf("a%d.png", code)
	}

	return fmt.Sprintf("%s%d.png", prefix, code)
}

// func main() {
// 	// 職業判定のデモ
// 	profession := judgeProfession("A", []string{"Python", "Java"}, []float64{20, 20})
// 	fmt.Println(profession) // 特級 士官 などの出力を期待

// 	// 画像ファイル名のディスパッチデモ
// 	picture := dispatchPictureBasedOnProfession("初級 攻撃魔法")
// 	fmt.Println(picture) // b1.png などの出力を期待
// }

func main() {
    // デモ
    fmt.Println(dispatchPictureBasedOnProfession("赤ちゃん")) // "c-.png"
    fmt.Println(dispatchPictureBasedOnProfession("少年")) // "c.png"
    fmt.Println(dispatchPictureBasedOnProfession("冒険者見習い")) // "c+.png"
    fmt.Println(dispatchPictureBasedOnProfession("上級 攻撃魔法")) // 他のロジックに基づいて決定される画像
}
