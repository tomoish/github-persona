package funcs

import (
    "bytes"
    "github.com/fogleman/gg"
    "math"
)

// ゲージを描画してバイトスライスで返す
func DrawGauge(percentage float64) ([]byte, error) {
    const S = 1024
    dc := gg.NewContext(S, S)
    dc.SetRGB(1, 1, 1)
    dc.Clear()

    radius := (float64(S) - 30.0) / 2
    lineWidth := 30.0
    startAngle := -math.Pi / 2

    // ゲージ背景の描画
    dc.SetRGBA(0, 0, 0, 0)
    dc.SetLineWidth(lineWidth)
    dc.DrawArc(S/2, S/2, radius, 0, math.Pi*2)
    dc.Stroke()
    dc.Clear()

    // ゲージの全体を描画
    allAngle := 2 * math.Pi * 1.0
    dc.SetRGB(0.5, 0.5, 0.5)
    dc.SetLineWidth(lineWidth)
    dc.DrawArc(S/2, S/2, radius, startAngle, startAngle+allAngle)
    dc.Stroke()

    // ゲージの進捗部分の描画
    progressAngle := 2 * math.Pi * percentage
    dc.SetRGB(0, 0.5, 0.7)
    dc.SetLineWidth(lineWidth)
    dc.DrawArc(S/2, S/2, radius, startAngle, startAngle+progressAngle)
    dc.Stroke()

    // 画像をバイトスライスにエンコード
    var buf bytes.Buffer
    err := dc.EncodePNG(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}