package graphs

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	bgColor := color.RGBA{R: 51, G: 61, B: 79, A: 255}
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
	line.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	p.Add(line)

	p.Title.Text = "Contribution History"
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "Commits"

	p.X.Min = -366
	p.X.Max = 5
	p.Y.Min = -0
	p.Y.Max = float64(maxCommits) + 5

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

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "./images/commits_history.png"); err != nil {
		panic(err)
	}

	return err
}
