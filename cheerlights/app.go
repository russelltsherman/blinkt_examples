package main

import (
	"os"
	"strconv"

	"github.com/russelltsherman/blinkt"
)

func getEnv(envVar string, assumed int) int {
	if value, exists := os.LookupEnv(envVar); exists {
		if period, err := strconv.Atoi(value); err == nil {
			return period
		}
	}
	return assumed
}

const defaultRefreshSeconds = 60
const envRefreshSeconds = "refresh_seconds"

func main() {
	brightness := 0.5
	blinky := blinkt.NewBlinkt(brightness)
	checkPeriodSeconds := getEnv(envRefreshSeconds, defaultRefreshSeconds)
	blinky.SetClearOnExit(true)
	blinky.Setup()
	blinkt.Delay(100)

	for {
		r, g, b := getCheerlightColours()
		blinky.Clear()
		blinky.SetAll(r, g, b)
		blinky.Show()
		blinkt.Delay(checkPeriodSeconds * 1000)
	}
}
