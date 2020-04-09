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

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0xf6, 0x37, 0x0e, 0xff},
	color.RGBA{0x0e, 0xd4, 0x31, 0xee},
	color.RGBA{0x96, 0x0e, 0xd4, 0xee},
	color.RGBA{0xef, 0xf9, 0x0e, 0xff}}

const (
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1 // next color in palette
	redIndex    = 2
	greenIndex  = 3
	purpleIndex = 4
	yellowIndex = 5
)

// go run images.go >out.gif
func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	rand.Seed(time.Now().UnixNano())
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				rnd())
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func rnd() uint8 {

	min := 0
	max := len(palette)

	return uint8(rand.Intn(max-min) + min)
}
