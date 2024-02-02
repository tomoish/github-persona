package graphs

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	// "github.com/fogleman/gg"
	// "log"
)

func DrawCommitChart(commitsHistory []int, maxCommits int, width int, height int) error {
	y := make([]float64, len(commitsHistory))
	for i := range y {
		y[i] = float64(commitsHistory[i])
	}
	x := make([]float64, len(commitsHistory))
	for i := range x {
		x[i] = float64(i-len(commitsHistory)) + 1
	}

	if len(x) != len(y) {
		fmt.Println("x and y arrays must have the same length")
	}

	p := plot.New()

	bgColor := color.RGBA{R: 51, G: 51, B: 51, A: 255}
	p.BackgroundColor = bgColor

	points := make(plotter.XYs, len(x))
	for i := range x {
		points[i].X = x[i]
		points[i].Y = y[i]
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{R: 135, G: 206, B: 235, A: 255}
	p.Add(line)
	plot.DefaultFont = font.Font{
		Typeface: "Roboto-Medium.ttf",
		Variant:  "Roboto-Medium.ttf",
		Size:     12.0,
	}
	p.Title.Text = "Contribution History"
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "Commits"

	p.X.Min = -366
	p.X.Max = 5
	p.Y.Min = -0
	p.Y.Max = float64(maxCommits) + 5

	// ラベルの外に余白を持つ
	p.Title.Padding = 10   // タイトル周りの余白
	p.X.Label.Padding = 2 // X軸ラベル周りの余白
	p.Y.Label.Padding = 2 // Y軸ラベル周りの余白

	// DefaultTextHandler is the default text handler used for text processing.

	white := color.White
	p.Title.TextStyle.Color = white
	p.X.Tick.Color = white
	p.Y.Tick.Color = white
	p.X.Label.TextStyle.Color = white
	p.Y.Label.TextStyle.Color = white
	p.X.Tick.Label.Color = white
	p.Y.Tick.Label.Color = white
	p.X.LineStyle.Color = white
	p.Y.LineStyle.Color = white

	p.X.Padding, p.Y.Padding = 0, 0

	if err := p.Save(5*vg.Inch, 2*vg.Inch, "./images/commits_history.png"); err != nil {
		panic(err)
	}

	// ggを使用してテキストをオーバーレイ
	// dc := gg.NewContext(width, height)
	// err = dc.LoadFontFace("Roboto-Medium.ttf", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 背景を透明に設定
	// dc.SetRGBA(0, 0, 0, 0) // 背景を透明に
	// dc.SetRGB(0, 0, 0) // テキストの色を設定（黒色）
	// dc.DrawString("kkdkdkdkkdd", 50, 10) // テキストを描画する位置を指定
	// err = dc.SavePNG("./images/commits_history_with_text.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return err
}
