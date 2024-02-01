package funcs
import (
	"fmt"

)

func CreateCharacterImg() {
	// キャラクター画像のパス
	characterPath := "images/character.png"
	// ゲージ画像のパス
	gaugePath := "images/gauge.png"

	// ゲージ画像生成のためのチャネル
	gaugeImageChan := make(chan []byte)

	// ゲージ画像を非同期で生成
	go func() {
		GaugeBytes, _ := DrawGauge(0.5)
		gaugeImageChan <- GaugeBytes
	}()

	// ゲージ画像をファイルに保存
	gaugeBytes := <-gaugeImageChan
	err := SaveImage(gaugePath, gaugeBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// キャラクター画像とゲージ画像を合成
	mergedImage, err := MergeImages(characterPath, gaugePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 画像をファイルに保存
	err = SaveImage("images/merged.png", mergedImage)
	if err != nil {
		fmt.Println(err)
		return
	}
}
