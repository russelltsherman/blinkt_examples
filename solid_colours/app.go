package main

import "github.com/russelltsherman/blinkt"

func main() {
	brightness := 0.1
	blinky := blinkt.NewBlinkt(brightness)
	blinky.SetClearOnExit(true)
	blinky.Setup()
	blinkt.Delay(100)
	step := 0
	for {
		step = step % 3
		switch step {
		case 0:
			blinky.SetAll(128, 0, 0)
		case 1:
			blinky.SetAll(0, 128, 0)
		case 2:
			blinky.SetAll(0, 0, 128)
		}
		step++
		blinky.Show()
		blinkt.Delay(500)
	}
}
