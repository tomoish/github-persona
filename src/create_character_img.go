package character_img

import (
	"fmt"
	"github.com/tomoish/readme/funcs"
)


func CreateCharacterImg() {

	// キャラクター画像のパス
	characterPath := "images/character.png"

	// ゲージ画像のパス
	gaugePath := "images/gauge.png"

	// キャラクター画像とゲージ画像を合成
	mergedImage, err := funcs.MergeImages(characterPath, gaugePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 画像をファイルに保存
	err = funcs.SaveImage(mergedImage, "images/merged.png")
	if err != nil {
		fmt.Println(err)
		return
	}
}	