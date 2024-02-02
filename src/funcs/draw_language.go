package funcs

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/fogleman/gg"
)

type LanguageStat struct {
	Name    string
	Percent float64
	Color   string
}

func drawRoundedRectangle(dc *gg.Context, x, y, w, h, r float64) {
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()
}

func GenerateLanguageUsageGraph(languages []LanguageStat, width, height int) ([]byte, error) {
	const cornerRadius = 10.0
	dc := gg.NewContext(width, height)

	// 背景とタイトルの描画
	dc.SetRGB(0.2, 0.2, 0.2) // 背景色（暗い青灰色）
	drawRoundedRectangle(dc, 0, 0, float64(width), float64(height), cornerRadius)
	dc.SetRGB(1, 1, 1)
	err := dc.LoadFontFace("Roboto-Medium.ttf", 45) // フォントとサイズの設定が必要
	if err != nil {
		fmt.Println(err)
	}
	dc.DrawStringAnchored("Most Used Languages", float64(width)/2, 50, 0.5, 0.5)

	// 帯グラフの基準点
	barX := 40.0
	barY := 120.0
	barHeight := 10.0
	barWidth := float64(width) - 80.0

	// 帯グラフの背景
	dc.SetRGB(0.16, 0.20, 0.29) // バーの背景色
	drawRoundedRectangle(dc, barX, barY, barWidth, barHeight, 5)

	// 割合でソート
	sort.SliceStable(languages, func(i, j int) bool {
		return languages[i].Percent > languages[j].Percent
	})


	// 5%未満の言語を抽出し、"others"としてまとめる
	var otherPercent float64
	// 削除対象の言語を探し、スライスから削除
	newLanguages := make([]LanguageStat, 0)
	for _, lang := range languages {
		if lang.Percent < 5.0 {
			otherPercent += lang.Percent

		} else {
			newLanguages = append(newLanguages, lang)
			continue
		}
	}


	// "others"を追加
	if otherPercent > 0 {
		othersStat := LanguageStat{
			Name:    "Others",
			Percent: otherPercent,
			Color:   "#CCCCCC", // 任意の色
		}
		newLanguages = append(newLanguages, othersStat)
	}

	// 帯グラフの各セクションを描画
	for _, lang := range newLanguages {
		dc.SetHexColor(lang.Color)
		sectionWidth := barWidth * lang.Percent / 100
		drawRoundedRectangle(dc, barX, barY, sectionWidth, barHeight, 5)
		barX += sectionWidth
	}

	// 凡例の描画
	legendX := 40.0
	legendY := barY + barHeight + 40.0             // 帯グラフの下に余白をとる
	err = dc.LoadFontFace("Roboto-Medium.ttf", 28) // 凡例のフォントサイズ
	if err != nil {
		fmt.Println(err)
	}

	i := 0.0
	for _, lang := range newLanguages {
		// 色のサンプルを描画
		dc.SetHexColor(lang.Color)
		dc.DrawCircle(legendX+5.0+math.Mod(float64(int(i)), 2.0)*300.0, legendY+30, 5)
		dc.Fill()

		// 言語名と割合を描画
		dc.SetRGB(1, 1, 1)
		dc.DrawString(fmt.Sprintf("%s %.2f%%", lang.Name, lang.Percent), legendX+20.0+math.Mod(float64(int(i)), 2.0)*300.0, legendY+45)
		i += 1.0
		legendY = legendY + math.Mod(float64(int(i)+1), 2.0)*40.0

	}

	// 画像をバイトデータにエンコード
	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
