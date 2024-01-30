package funcs

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"

)

type LanguageStat struct {
	Name   string
	Percent float64
	Color  string
}

func drawRoundedRectangle(dc *gg.Context, x, y, w, h, r float64) {
    dc.DrawRectangle(x, y, w, h)
    dc.Fill()

}

func generateLanguageUsageGraph(languages []LanguageStat, width, height int) error {
	const cornerRadius = 10.0
	dc := gg.NewContext(width, height)

	// 背景とタイトルの描画
	dc.SetRGB(0.2, 0.24, 0.31) // 背景色（暗い青灰色）
	drawRoundedRectangle(dc, 0, 0, float64(width), float64(height), cornerRadius)
	dc.SetRGB(1, 1, 1)
	dc.LoadFontFace("Roboto-Medium.ttf", 30) // フォントとサイズの設定が必要
	dc.DrawStringAnchored("Most Used Languages", float64(width)/2, 30, 0.5, 0.5)

	// 帯グラフの基準点
	barX := 40.0
	barY := 60.0
	barHeight := 10.0
	barWidth := float64(width) - 80.0

	// 帯グラフの背景
	dc.SetRGB(0.16, 0.20, 0.29) // バーの背景色
	drawRoundedRectangle(dc, barX, barY, barWidth, barHeight, 5)

	// 帯グラフの各セクションを描画
	for _, lang := range languages {
		dc.SetHexColor(lang.Color)
		sectionWidth := barWidth * lang.Percent / 100
		drawRoundedRectangle(dc, barX, barY, sectionWidth, barHeight, 5)
		barX += sectionWidth
	}

	// 凡例の描画
	legendX := 40.0
	legendY := barY + barHeight + 20.0 // 帯グラフの下に余白をとる
	dc.LoadFontFace("Roboto-Medium.ttf", 25) // 凡例のフォントサイズ
	i := 0.0
	for _, lang := range languages {
		// 色のサンプルを描画
		dc.SetHexColor(lang.Color)
		dc.DrawCircle(legendX+5.0+math.Mod(float64(int(i)), 2.0)*300.0, legendY+6, 5)
		dc.Fill()

		// 言語名と割合を描画
		dc.SetRGB(1, 1, 1)
		dc.DrawString(fmt.Sprintf("%s %.2f%%", lang.Name, lang.Percent), legendX+20.0+math.Mod(float64(int(i)), 2.0)*300.0, legendY+15)
		i += 1.0
		legendY = legendY + math.Mod(float64(int(i)+1), 2.0)*40.0
		
		
	}

	// 画像をファイルに保存
	dc.SavePNG("language_usage_graph.png")
	return nil
}
