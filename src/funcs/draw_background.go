package funcs

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
)

func DrawBackground(username, level, kind string) {
	const (
		// Width  = 1200
		// Height = 1800
		Width      = 800
		Height     = 1200
		BorderSize = 10 // ボーダーのサイズ
	)

	// 新しい画像コンテキストを作成
	dc := gg.NewContext(Width, Height)

	// 画像の背景を白に設定
	dc.SetRGB(0.03, 0.03, 0.03)
	dc.Clear()

	// テキストの描画設定
	// フォントの設定

	// 関数呼び出しとエラーチェック
	if err := dc.LoadFontFace("Roboto-Medium.ttf", 28); err != nil {
		fmt.Println("フォントのロードに失敗しました:", err)
	}
	dc.SetColor(color.White)

	// テキストの座標を個別に指定して描画
	text1 := "Hi, there! I'm " + username
	x1, y1 := (25*Width)/100, (15*Height)/100 // テキスト1の座標
	dc.DrawStringAnchored(text1, float64(x1), float64(y1), 0.5, 0.5)

	text2 := level
	x2, y2 := (13*Width)/100, (25*Height)/100 // テキスト2の座標(テキストの中央に持つ)
	dc.DrawStringAnchored(text2, float64(x2), float64(y2), 0.5, 0.5)

	text3 := kind
	if strings.Contains(kind, "ネクロマンサー") {
		if err := dc.LoadFontFace("ipaexg.ttf", 26); err != nil {
			fmt.Println("フォントのロードに失敗しました:", err)
		}
	} else {

		if err := dc.LoadFontFace("ipaexg.ttf", 30); err != nil {
			fmt.Println("フォントのロードに失敗しました:", err)
		}
	}
	x3, y3 := (33*Width)/100, (25*Height)/100 // テキスト3の座標
	dc.DrawStringAnchored(text3, float64(x3), float64(y3), 0.5, 0.5)

	// ボーダーを描画する新しいコンテキストを作成
	dcBorder := gg.NewContext(Width, Height)
	dcBorder.SetRGB(1, 1, 1) // 白いボーダー
	dcBorder.Clear()
	dcBorder.DrawRectangle(0, 0, float64(Width), float64(Height))
	dcBorder.SetLineWidth(BorderSize)
	dcBorder.Stroke()

	// ボーダーの上に背景画像を描画
	dcBorder.DrawImage(dc.Image(), 0, 0)
	// 画像をPNG形式で保存
    imageFileName := fmt.Sprintf("./images/background_%s.png", username)
	if err := dc.SavePNG(imageFileName ); err != nil {
		log.Fatal(err)
	}

	log.Println("Image with multiple text has been created and saved.")
}
