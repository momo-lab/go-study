package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

type LissajousParams struct {
	Cycles  float64
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// TODO: HTTP Requestのパラメータ解析についてはもうちょっと定石を知りたい
func handler(w http.ResponseWriter, r *http.Request) {
	param := LissajousParams{
		Cycles:  5,
		Res:     0.001,
		Size:    100,
		Nframes: 64,
		Delay:   8,
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}
	for k, v := range r.Form {
		log.Print(k, v)
		switch k {
		case "cycles":
			param.Cycles, _ = strconv.ParseFloat(v[0], 64)
		case "res":
			param.Res, _ = strconv.ParseFloat(v[0], 64)
		case "size":
			param.Size, _ = strconv.Atoi(v[0])
		case "nframes":
			param.Nframes, _ = strconv.Atoi(v[0])
		case "delay":
			param.Delay, _ = strconv.Atoi(v[0])
		}
	}
	log.Print(param)
	lissajous(param, w)
}

func lissajous(param LissajousParams, out io.Writer) {
	freq := rand.Float64() * 3.0 // 発信器 y の相対周波数
	anim := gif.GIF{LoopCount: param.Nframes}
	phase := 0.0 // 位相差
	for i := 0; i < param.Nframes; i++ {
		rect := image.Rect(0, 0, 2*param.Size+1, 2*param.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < param.Cycles*2*math.Pi; t += param.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(param.Size+int(x*float64(param.Size)+0.5), param.Size+int(y*float64(param.Size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, param.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
