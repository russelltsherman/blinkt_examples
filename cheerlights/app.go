package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

type cheerlight struct {
	Colour string `json:"field2"`
}

func getRGBFromColour(hexString string) (int, int, int) {
	//strip off the hash and decode the remaining hex value
	hexCol, _ := hex.DecodeString(strings.TrimPrefix(hexString, "#"))
	//return the ints representing rgb
	return int(hexCol[0]), int(hexCol[1]), int(hexCol[2])
}

func getCheerlightColours() (int, int, int) {
	var netClient = &http.Client{
		Timeout: time.Second * 3,
	}
	resp, getErr := netClient.Get("http://api.thingspeak.com/channels/1417/field/2/last.json")

	if getErr != nil {
		log.Panic(getErr)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Panic(getErr)
	}

	result := cheerlight{}

	parseErr := json.Unmarshal(body, &result)
	if parseErr != nil {
		log.Panic("Can't parse response")
		return 0, 0, 0
	}

	return getRGBFromColour(result.Colour)

}

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
