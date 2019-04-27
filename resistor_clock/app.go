package main

import (
	"time"

	"github.com/russelltsherman/blinkt"
)

func breakOut(colours []int) (int, int, int) {
	r := colours[0]
	g := colours[1]
	b := colours[2]
	return r, g, b
}

func main() {
	brightness := 0.5
	blinky := blinkt.NewBlinkt(brightness)
	blinky.SetClearOnExit(true)

	colours := [10][3]int{
		{0, 0, 0},       //0 black
		{139, 69, 19},   //1 brown
		{255, 0, 0},     //2 red
		{255, 69, 0},    //3 orange
		{255, 255, 0},   //4 yellow
		{0, 255, 0},     //5 green
		{0, 0, 255},     //6 blue
		{128, 0, 128},   //7 violet
		{255, 255, 100}, //8 grey
		{255, 255, 255}, //9 white
	}

	blinky.Setup()
	blinkt.Delay(100)

	r := 0
	g := 0
	b := 0

	for {

		t := time.Now()
		hour := t.Hour()
		minute := t.Minute()

		hourten := hour / 10
		hourunit := hour % 10
		minuteten := minute / 10
		minuteunit := minute % 10

		r, g, b = breakOut(colours[hourten][:])
		blinky.SetPixel(0, r, g, b)
		blinky.SetPixel(1, r, g, b)

		r, g, b = breakOut(colours[hourunit][:])
		blinky.SetPixel(2, r, g, b)
		blinky.SetPixel(3, r, g, b)

		r, g, b = breakOut(colours[minuteten][:])
		blinky.SetPixel(4, r, g, b)
		blinky.SetPixel(5, r, g, b)

		r, g, b = breakOut(colours[minuteunit][:])
		blinky.SetPixel(6, r, g, b)
		blinky.SetPixel(7, r, g, b)

		blinky.Show()
		blinkt.Delay(500)

		blinky.SetPixel(7, 0, 0, 0)

		blinky.Show()
		blinkt.Delay(500)

	}
}
