package main

import (
	colorsys "github.com/cckuailong/colorsys-go"
	"github.com/russelltsherman/blinkt"
)

func main() {

	brightness := 0.1
	blinky := blinkt.NewBlinkt(brightness)
	blinky.SetClearOnExit(true)
	blinky.Setup()
	blinkt.Delay(100)

	h := 0
	for {
		h++
		if h > 360 {
			h = 0
		}

		r, g, b := colorsys.Hsv2Rgb(float64(h), 1.0, brightness)
		// fmt.Printf("R: %v G: %v B: %v\n", r, g, b)
		blinky.Clear()
		blinky.SetAll(int(r), int(g), int(b))
		blinky.Show()
		blinkt.Delay(100)
	}
}
