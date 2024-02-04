package funcs

import (
	"fmt"
)

func CreateCharacterImg(characterPath, gaugePath string, total, level int, username string) {
	// ゲージ画像生成のためのチャネル
	gaugeImageChan := make(chan []byte)
	percentage := float64(15-(total%15)) / float64(15)
	fmt.Println(percentage)
	// ゲージ画像を非同期で生成
	go func() {
		GaugeBytes, _ := DrawGauge(percentage)
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
	imageFileName := fmt.Sprintf("./images/generate_character_%s.png", username)
	err = SaveImage(imageFileName, mergedImage)
	if err != nil {
		fmt.Println(err)
		return
	}
}
