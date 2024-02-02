package funcs

import (
	"bytes"
	"strconv"

	"github.com/fogleman/gg"
)

// GenerateGitHubStatsImage 関数はユーザー統計情報を受け取り、グラフィカルな表現を画像として生成します。
func GenerateGitHubStatsImage(stats UserStats, width, height int) ([]byte, error) {
	const padding = 20.0
	const lineHeight = 30.0

	dc := gg.NewContext(width, height)

	// 背景色を設定します。
	dc.SetRGB(0.2, 0.24, 0.31)
	dc.Clear()

	// タイトルを描画します。
	dc.SetRGB(1, 1, 1) // 白色
	dc.DrawStringAnchored("@hashfx-Github-stats", float64(width)/2, padding, 0.5, 0.5)

	// 統計情報をリストとして描画します。
	statsList := []struct {
		Icon  string
		Label string
		Value int
	}{
		{"★", "Total Stars Earned:", stats.TotalStars},
		{"⧗", "Total Commits:", stats.TotalCommits},
		{"⬆️", "Total PRs:", stats.TotalPRs},
		{"⬇️", "Total Issues:", stats.TotalIssues},
		{"⬈", "Contributed to:", stats.ContributedTo},
	}

	// 各統計情報をリスト形式で描画します。
	for i, stat := range statsList {
		y := padding*2 + lineHeight*float64(i)
		dc.DrawStringAnchored(stat.Icon, padding, y, 0, 0.5)
		dc.DrawStringAnchored(stat.Label+" "+strconv.Itoa(stat.Value), padding+40, y, 0, 0.5)
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
