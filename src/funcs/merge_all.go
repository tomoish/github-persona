package funcs

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// 画像を読み込む関数
func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

//原点（0, 0）は背景画像の左上隅

// すべての画像のパスを受け取り、それらを合成した画像を生成する関数
func Merge_all(nameImg, statsImg, characterImg, languageImg, dateImg, filepath string) {
	// 背景画像を黒で生成
	// backgroundWidth := 2000 // 背景画像の幅
	// backgroundHeight := 1900 // 背景画像の高さ
	// backgroundColor := color.White
	// backgroundImage := image.NewRGBA(image.Rect(0, 0, backgroundWidth, backgroundHeight))
	// draw.Draw(backgroundImage, backgroundImage.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
	// 背景画像を読み込む
	backgroundImage, err := loadImage(nameImg)
	if err != nil {
		fmt.Println("Error loading background image:", err)
		return
	}

	// 背景画像の幅と高さを取得
	backgroundWidth := backgroundImage.Bounds().Dx()
	backgroundHeight := backgroundImage.Bounds().Dy()
	// 合成する画像の読み込み
	overlayImages := []string{statsImg, characterImg, languageImg, dateImg}

	// 画像の幅と高さを取得
	// width0, height0, _ := getImageSize(statsImg)
	// width1, height1, _ := getImageSize(characterImg)
	// width2, height2, _ := getImageSize(languageImg)
	// width3, height3, _ := getImageSize(dateImg)

	// アスペクト比を計算
	// aspectRatio0 := float64(width0) / float64(height0)
	// aspectRatio1 := float64(width1) / float64(height1)
	// aspectRatio2 := float64(width2) / float64(height2)
	// aspectRatio3 := float64(width3) / float64(height3)

	// 画像のリサイズ

	ResizeImage(statsImg, uint(46*backgroundWidth)/100, 0)
	ResizeImage(characterImg, uint(40*backgroundWidth)/100, uint(40*backgroundWidth)/100)
	ResizeImage(languageImg, uint((46*backgroundWidth)/100), 0)
	ResizeImage(dateImg, uint((80*backgroundWidth)/100), 0)
	// 画像の幅と高さを取得
	// width0, height0, _ := getImageSize(statsImg)
	_, height1, _ := getImageSize(characterImg)
	_, height2, _ := getImageSize(languageImg)
	// width3, height3, _ := getImageSize(dateImg)

	// 画像を背景の上に重ねる
	for i, overlayImage := range overlayImages {
		overlay, err := loadImage(overlayImage)
		if err != nil {
			fmt.Println("Error loading overlay image:", err)
			return
		}
		overlayX := 0
		overlayY := 0

		switch i {
		case 0:
			// 配置する座標を計算
			overlayX = (2 * backgroundWidth) / 100
			overlayY = height1 + (10*backgroundHeight)/100
		case 1:
			overlayX = (50 * backgroundWidth) / 100
			overlayY = (5 * backgroundHeight) / 100
		case 2:
			overlayX = (52 * backgroundWidth) / 100
			overlayY = height1 + (10*backgroundHeight)/100
		case 3:
			overlayX = (10 * backgroundWidth) / 100
			overlayY = height1 + height2 + (15*backgroundHeight)/100
		default:
			overlayX = 0
			overlayY = 0
		}
		overlayPos := image.Point{overlayX, overlayY}

		// 合成画像を背景画像に描画
		draw.Draw(backgroundImage.(*image.RGBA), overlay.Bounds().Add(overlayPos), overlay, image.Point{}, draw.Over)
	}

	// 合成した画像を保存
	saveImage(backgroundImage, filepath)
	ResizeImage("result.png", 700, 0)
}

// 画像を保存する関数
func saveImage(img image.Image, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating image file:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img) // PNG形式で保存する
	if err != nil {
		fmt.Println("Error encoding image as PNG:", err)
		return
	}

	fmt.Println("Image saved as", filename)
}

func ResizeImage(imageName string, width uint, height uint) {
	// 元の画像の読み込み
	file := imageName
	fileData, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	// 画像をimage.Image型にdecodeします
	img, data, err := image.Decode(fileData)
	if err != nil {
		log.Fatal(err)
	}
	fileData.Close()

	// ここでリサイズします
	// 片方のサイズを0にするとアスペクト比固定してくれます
	resizedImg := resize.Resize(width, height, img, resize.NearestNeighbor)

	// 書き出すファイル名を指定します
	createFilePath := imageName
	output, err := os.Create(createFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// 最後にファイルを閉じる
	defer output.Close()

	// 画像のエンコード(書き込み)
	switch data {
	case "png":
		if err := png.Encode(output, resizedImg); err != nil {
			log.Fatal(err)
		}
	case "jpeg", "jpg":
		opts := &jpeg.Options{Quality: 100}
		if err := jpeg.Encode(output, resizedImg, opts); err != nil {
			log.Fatal(err)
		}
	default:
		if err := png.Encode(output, resizedImg); err != nil {
			log.Fatal(err)
		}
	}
}

// 画像の幅と高さを取得する関数
func getImageSize(filename string) (int, int, error) {
	// 画像ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	// 画像形式を識別し、画像データをデコード
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
