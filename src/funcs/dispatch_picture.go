package funcs

import (
	"fmt"
	"strings"
)

func DispatchPictureBasedOnProfession(profession string) string {
	// 特定の職業または状態に応じた画像ファイル名を返す
	switch profession {
	case "赤ちゃん":
		return "c-.png"
	case "少年":
		return "c.png"
	case "冒険者見習い":
		return "c+.png"
	case "魔法使いの見習い":
		return "b-1.png"
	case "不良":
		return "b-2.png"
	case "駆け出し冒険者":
		return "b-3.png"
	case "裏社会のボス":
		return "a+1.png"
	case "魔王":
		return "a+2.png"
	case "魔法帝":
		return "a+3.png"
	case "勇者":
		return "a+4.png"
	case "騎士団長":
		return "a+5.png"
	case "賢者":
		return "a+6.png"
	case "サイコパス":
		return "a+7.png"
	case "バーカーサ":
		return "a-16.png"
	case "闇騎士":
		return "a-17.png"
	case "黒魔術師":
		return "a-18.png"
	case "ライダー":
		return "a-19.png"
	case "魔法戦士":
		return "a-20.png"	
	case "魔法騎士":
		return "a-21.png"	
	case "特級 バーカーサ","特級 闇騎士":
		return "a-16.png"
	case "特級 黒魔術師","特級 ライダー":
		return "a15.png"
	case "特級 魔法戦士","特級 魔法騎士":
		return "a14.png"	
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
		"攻撃魔術師": 1, "ネクロマンサー": 2, "防御魔術師": 3,
		"召喚師": 4, "精霊魔法": 5, "回復術師": 6,
		"賞金稼ぎ": 7, "犯罪者": 8, "盗賊": 9,
		"武闘家": 10, "弓使い": 11, "槍使い": 12,
		"剣士": 13, "騎士": 14, "士官": 15,
	}

	
	// ここを時間があったら細かく
	parts := strings.Split(profession, " ")
	rankPart, professionPart := "", ""
	if len(parts) > 1 {
		rankPart = parts[0]
		professionPart = strings.Join(parts[1:], " ")
	} else {
		professionPart = parts[0]
	}

	if rankPart == "a" {
		switch professionPart{
			case "犯罪者","賞金稼ぎ":
				return "a2.png"
			case "盗賊":
				return "a1.png"
			case "ネクロマンサー":
				return "a4.png"
			case "攻撃魔術師":
				return "a6.png"
			case "防御魔術師","回復術師":
				return "a7.png"
			case "召喚士","精霊魔術師":
				return "a5.png"
			case "弓使い":
				return "a8.png"
			case "槍使い":
				return "a9.png"
			case "剣士":
				return "a10.png"
			case "武闘家":
				return "a11.png"
			case "士官":
				return "a12.png"
			case "騎士":
				return "a13.png"
		}
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
// 	picture := dispatchPictureBasedOnProfession("初級 攻撃魔術師")
// 	fmt.Println(picture) // b1.png などの出力を期待
// }

// func main() {
//     // デモ
//     fmt.Println(dispatchPictureBasedOnProfession("赤ちゃん")) // "c-.png"
//     fmt.Println(dispatchPictureBasedOnProfession("少年")) // "c.png"
//     fmt.Println(dispatchPictureBasedOnProfession("冒険者見習い")) // "c+.png"
//     fmt.Println(dispatchPictureBasedOnProfession("上級 攻撃魔術師")) // 他のロジックに基づいて決定される画像
// }
