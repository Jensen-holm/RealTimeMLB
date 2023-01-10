package main

import (

	// Packages made specific to this project
	"github.com/Jensen-holm/RealTimeMLB/baseball"
	stadium "github.com/Jensen-holm/RealTimeMLB/models/SoccerStadium"

	// Wrapper packages for g3n that I made
	"github.com/Jensen-holm/g3n-Wrapper/apper"
	"github.com/Jensen-holm/g3n-Wrapper/apper/model"
	"github.com/Jensen-holm/g3n-Wrapper/apper/phys"

	"log"
)

func main() {
	a := apper.NewApp(true)
	Init(a)
	a.Run()
}

// Init -> Customizing the application wrapper (apper)
// for our specific use case. Ran once before running the app.
func Init(a *apper.App) {

	a.AddBg(.5, .75, 2, .5)

	// Creating physics simulation
	sim := phys.NewSim()
	a.AddSim(sim)
	sim.SetGravity(0, -32.2, 0)

	// Creating objects to render
	ball := baseball.NewBaseball()

	// eventually this will be outside the init function
	// and inside our own custom update function that is run each frame
	// and forces and balls will be added based on scraped baseball data
	// in real time while a certain mlb game is going on
	sim.AddSphere(ball)
	ball.ApplyForce(
        baseball.CalculateVector(20, 150, 15),
    )

	// making the plane, added it to the simulation but not to the
	// scene so that it is invisible for testing our stadium model
	ground := model.NewPlane(10000, 10000, 90, "slategray", false)
	sim.SetPlane(ground)

	// Creating lights
	l1 := apper.Light("white", 1, 100, 100, 100)
	l2 := apper.Light("white", 1, -100, 100, -100)
	l3 := apper.Light("white", 1, -100, -100, -100)
	l4 := apper.Light("white", 1, 100, -100, 100)

	// Importing models built with other tools (blender)
	// The stadium package is one that is not a part of the wrapper
	arena, err := stadium.ImportStadium()
	if err != nil {
		log.Fatalf("error importing the stadium: %v", err)
	}

	// Adding everything to the scene
	a.Add2Scene(
		a.Cam.Self,
		l1, l2, l3, l4,
		ball.Mesh,
		arena,
	)
}
