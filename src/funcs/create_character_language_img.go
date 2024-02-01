package funcs

import (
	"fmt"
)

func CreateCharacterLanguageImg() {
	// キャラクター画像のパス
	characterPath := "funcs/images/merged.png"
	// ゲージ画像のパス
	languagePath := "funcs/images/language.png"

	// キャラクター画像とゲージ画像を合成
	mergedImage, err := MergeCharacterLanguageImages(characterPath, languagePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 画像をファイルに保存
	err = SaveImage("funcs/images/characterLanguageMerged.png", mergedImage)
	if err != nil {
		fmt.Println(err)
		return
	}
}
