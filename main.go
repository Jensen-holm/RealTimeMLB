package main

import (
	"github.com/Jensen-holm/RealTimeMLB/app"
	"log"
)

var a = app.NewApp()

func main() {

	err := a.RunApp()
	if err != nil {
		log.Fatalf("error running application: %v", err)
	}

}
