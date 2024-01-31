package character

import (
	"fmt"
	"github.com/tomoish/readme/funcs"
)

func CreateCharacterImg() {
	// キャラクター画像のパス
	characterPath := "funcs/images/character.png"
	// ゲージ画像のパス
	gaugePath := "funcs/images/gauge.png"

	// ゲージ画像生成のためのチャネル
	gaugeImageChan := make(chan []byte)

	// ゲージ画像を非同期で生成
	go func() {
		GaugeBytes, _ := funcs.DrawGauge(0.5)
		gaugeImageChan <- GaugeBytes
	}()

	// ゲージ画像をファイルに保存
	gaugeBytes := <-gaugeImageChan
	err := funcs.SaveImage(gaugePath, gaugeBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// キャラクター画像とゲージ画像を合成
	mergedImage, err := funcs.MergeImages(characterPath, gaugePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 画像をファイルに保存
	err = funcs.SaveImage("funcs/images/merged.png", mergedImage)
	if err != nil {
		fmt.Println(err)
		return
	}
}
