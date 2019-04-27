package main

import (
	"math/rand"

	"github.com/russelltsherman/blinkt"
)

func main() {

	brightness := 0.5
	blinky := blinkt.NewBlinkt(brightness)
	blinky.SetClearOnExit(true)
	blinky.Setup()
	blinkt.Delay(100)

	for {
		for pixel := 0; pixel < 8; pixel++ {
			blinky.SetPixel(pixel, rand.Intn(255), rand.Intn(255), rand.Intn(255))
		}
		blinky.Show()
		blinkt.Delay(50)
	}
}
