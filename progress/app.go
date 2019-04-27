package main

import "github.com/russelltsherman/blinkt"

func main() {
	brightness := 0.5
	blinky := blinkt.NewBlinkt(brightness)

	blinky.SetClearOnExit(true)
	blinky.Setup()
	blinkt.Delay(100)

	r := 150
	g := 0
	b := 0
	for {
		for pixel := 0; pixel < 8; pixel++ {
			blinky.Clear()
			blinky.SetPixel(pixel, r, g, b)
			blinky.Show()
			blinkt.Delay(100)
		}
		for pixel := 7; pixel > 0; pixel-- {
			blinky.Clear()
			blinky.SetPixel(pixel, r, g, b)
			blinky.Show()
			blinkt.Delay(100)
		}
	}
}
