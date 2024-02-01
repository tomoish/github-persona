package funcs

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func MergeCharacterLanguageImages(
	characterPath, languagePath string,
) ([]byte, error) {
	characterFile, err := os.Open(characterPath)
	if err != nil {
		return nil, err
	}
	defer characterFile.Close()

	characterImage, _, err := image.Decode(characterFile)
	if err != nil {
		return nil, err
	}

	languageFile, err := os.Open(languagePath)
	if err != nil {
		return nil, err
	}
	defer languageFile.Close()

	languageImage, _, err := image.Decode(languageFile)
	if err != nil {
		return nil, err
	}

	// 画像のサイズを取得
	characterBounds := characterImage.Bounds()
	languageBounds := languageImage.Bounds()

	// 新しい画像のサイズを計算（横並び）
	width := characterBounds.Dx() + languageBounds.Dx()
	height := max(characterBounds.Dy(), languageBounds.Dy())

	// 新しい画像を作成
	mergedImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// 最初の画像（characterImage）を描画
	draw.Draw(mergedImage, characterBounds, characterImage, image.Point{}, draw.Src)

	// 2番目の画像（languageImage）を描画
	languagePosition := image.Point{X: characterBounds.Dx(), Y: 0}
	draw.Draw(mergedImage, image.Rectangle{Min: languagePosition, Max: languagePosition.Add(languageBounds.Size())}, languageImage, image.Point{}, draw.Src)

	// 画像をバイト列に変換
	var imgBytes bytes.Buffer
	err = png.Encode(&imgBytes, mergedImage)
	if err != nil {
		return nil, err
	}

	return imgBytes.Bytes(), nil
}
