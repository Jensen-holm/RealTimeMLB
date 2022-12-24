package main

import (
	"github.com/Jensen-holm/RealTimeMLB/win"
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
