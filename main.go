package main

import (
	"github.com/Jensen-holm/RealTimeMLB/app"
	"log"
)

func main() {

	w := app.NewWindow()

	//w.ToggleHelp(false)

	w.SetUP()

	err := w.RunApp()
	if err != nil {
		log.Fatalf("error running the app: %v", err)
	}

}
