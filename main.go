package main

import (
	win "github.com/Jensen-holm/RealTimeMLB/app"
	"log"
)

func main() {

	w := win.NewWindow()
	w.Init()

	err := w.RunApp()
	if err != nil {
		log.Fatalf("error running the app: %v", err)
	}

}
