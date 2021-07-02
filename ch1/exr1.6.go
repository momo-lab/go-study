package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = make([]color.Color, 256)

func init() {
	// パレットの初期化
	// パレットにnilが含まれると正しくimageが作成できない。
	for i := 0; i < 256; i++ {
		palette[i] = color.RGBA{0x00, byte(i), 0x00, 0xFF}
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発信器 x が完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは [-size..+size] の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms 単位でのフレーム間の遅延
	)

	freq := rand.Float64() * 3.0 // 発信器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := size + int(math.Sin(t)*size+0.5)
			y := size + int(math.Sin(t*freq+phase)*size+0.5)
			c := uint8(x * (256 / (2*size + 1)))
			img.SetColorIndex(x, y, c)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
