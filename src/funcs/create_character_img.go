package funcs

import (
	"fmt"
	"math"
)

func CreateCharacterImg(characterPath, gaugePath string, total, level int) {
	// ゲージ画像生成のためのチャネル
	gaugeImageChan := make(chan []byte)
	percentage := (float64(total) - float64(math.Pow(float64(level), 2))) / (float64(math.Pow(float64(level+1), 2) - math.Pow(float64(level), 2)))
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
	err = SaveImage("./images/generate_character.png", mergedImage)
	if err != nil {
		fmt.Println(err)
		return
	}
}
