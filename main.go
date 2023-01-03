package main

import (
	"github.com/Jensen-holm/RealTimeMLB/apper"
	"github.com/Jensen-holm/g3n-App-Skeleton/apper/model"
	"github.com/Jensen-holm/g3n-App-Skeleton/apper/phys"
)

func main() {
	a := apper.NewApp(false)
	Init(a)
	a.Run()
}

func Init(a *apper.App) {
	a.AddBg(.5, .75, 2, .5)

	sim := phys.NewSim()
	a.AddSim(sim)
	sim.SetGravity(0, -32.2, 0)

	ball := model.NewSphere(0, 50, 0, 3, 14.5, "red")
	ball2 := model.NewSphere(0, 500, 0, 10, 20.0, "green")

	//arena, err := stadium.ImportStadium()
	//if err != nil {
	//	log.Fatalf("error importing the stadium: %v", err)
	//}

	sim.AddSphere(ball, ball2)
	ball.ApplyForce(150, 150, 150)

	ground := model.NewPlane(10000, 10000, 90, "slategray", false)
	sim.SetPlane(ground)

	l1 := apper.Light("white", 1, 100, 100, 100)
	l2 := apper.Light("white", 1, -100, 100, -100)
	l3 := apper.Light("white", 1, -100, -100, -100)
	l4 := apper.Light("white", 1, 100, -100, 100)

	a.Add2Scene(
		a.Cam.Self,
		l1, l2, l3, l4,
		ground.Mesh,
		ball.Mesh,
		ball2.Mesh,
		//arena,
	)
}
