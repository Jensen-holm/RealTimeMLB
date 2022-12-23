package main

import (
	win "github.com/Jensen-holm/RealTimeMLB/app"
)

func main() {

	w := win.NewWindow()
	w.ToggleHelp(true)
	w.Init()

	err := w.RunApp()
	if err != nil {
		panic(err)
	}

}
