package funcs

import (
	"bytes"
	"strconv"

	"github.com/fogleman/gg"
)

// GenerateGitHubStatsImage 関数はユーザー統計情報を受け取り、グラフィカルな表現を画像として生成します。
func GenerateGitHubStatsImage(stats UserStats, star ,width, height int) ([]byte, error) {
	// const padding = 20.0
	// const lineHeight = 30.0

	dc := gg.NewContext(width, height)

	// 背景色を設定します。
	dc.SetRGB(0.2, 0.2, 0.2)
	dc.Clear()

	// タイトルを描画します。
	dc.SetRGB(0, 0.749, 1)
	if err := dc.LoadFontFace("Roboto-Medium.ttf", 50); err != nil {
		return nil, err
	}
	dc.DrawStringAnchored("Github-Stats", float64(width)/2, 10*float64(height)/100, 0.5, 0.5)

	// 統計情報をリストとして描画します。
	statsList := []struct {
		Label string
		Value int
	}{
		{"Total Stars Earned:", star},
		{"Total Commits:", stats.TotalCommits},
		{"Total PRs:", stats.TotalPRs},
		{"Total Issues:", stats.TotalIssues},
		{"Contributed to:", stats.ContributedTo},
	}

	dc.SetRGB(1, 1, 1) // 白色

	// 関数呼び出しとエラーチェック
	if err := dc.LoadFontFace("Roboto-Medium.ttf", 30); err != nil {
		return nil, err
	}

	// 各統計情報をリスト形式で描画します。
	for i, stat := range statsList {
		y := 0.3 + 0.15*float64(i)     // 垂直位置を計算
		valueX := 0.7 * float64(width) // 値を右揃えにするためのX座標
		dc.DrawString(stat.Label, 0.1*float64(width), y*float64(height))
		dc.DrawString(strconv.Itoa(stat.Value), valueX, y*float64(height))
	}
	// 画像をバッファにエンコードします。
	buf := bytes.Buffer{}
	err := dc.EncodePNG(&buf)
	if err != nil {
		return nil, err
	}

	// エンコードしたバッファをバイト配列として返します。
	return buf.Bytes(), nil
}
