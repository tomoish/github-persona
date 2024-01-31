package funcs

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func MergeImages(characterPath, gaugePath string) ([]byte, error) {
	// キャラクター画像をロード
	characterFile, err := os.Open(characterPath)
	if err != nil {
		return nil, err
	}
	defer characterFile.Close()

	characterImage, _, err := image.Decode(characterFile)
	if err != nil {
		return nil, err
	}

	// ゲージ画像をロード
	gaugeFile, err := os.Open(gaugePath)
	if err != nil {
		return nil, err
	}
	defer gaugeFile.Close()

	gaugeImage, _, err := image.Decode(gaugeFile)
	if err != nil {
		return nil, err
	}

	// キャラクター画像の半径を計算
	characterBounds := characterImage.Bounds()
	characterRadius := float64(max(characterBounds.Dx(), characterBounds.Dy())) / 2.0

	// ゲージ画像の半径を計算
	gaugeBounds := gaugeImage.Bounds()
	gaugeRadius := float64(max(gaugeBounds.Dx(), gaugeBounds.Dy())) / 2.0

	// ゲージ画像をリサイズ
	resizeRatio := characterRadius / gaugeRadius
	newDiameter := int((characterRadius + 30.0*resizeRatio) * 2.0)
	newSize := image.Point{newDiameter, newDiameter}
	resizedGaugeImage := resize.Resize(uint(newDiameter), uint(newDiameter), gaugeImage, resize.Lanczos3)

	// 新しいイメージを作成
	newImage := image.NewRGBA(image.Rect(0, 0, newDiameter, newDiameter))

	// ゲージを新しいイメージの中心に描画
	gaugeOffset := image.Pt(int((float64(newDiameter)-float64(newSize.X))/2.0), int((float64(newDiameter)-float64(newSize.Y))/2.0))
	draw.Draw(newImage, newImage.Bounds(), resizedGaugeImage, gaugeOffset, draw.Over)

	// キャラクターを新しいイメージの中心に描画
	characterOffset := image.Pt(int((float64(newDiameter)-float64(characterBounds.Dx()))/2.0), int((float64(newDiameter)-float64(characterBounds.Dy()))/2.0))
	draw.Draw(newImage, characterBounds.Add(characterOffset), characterImage, image.Point{}, draw.Over)

	// バイトスライスにエンコード
	var buf bytes.Buffer
	err = png.Encode(&buf, newImage)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
