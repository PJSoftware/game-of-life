package main

import (
	"./life"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("../life-go.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Game of Life: Go version")
	var cTimer, dTimer life.Timer
	var world life.World

	world.Init("default", "conway", true)
	fmt.Print(world.ToString())
	displayTimings(cTimer, dTimer, world.Step(), false)

	for world.Step() < 1000 || !cTimer.IsAverageStable() {
		cTimer.Start()
		world.Calculate()
		cTimer.Stop()

		dTimer.Start()
		fmt.Print(world.ToString())
		dTimer.Stop()

		displayTimings(cTimer, dTimer, world.Step(), false)
	}
	displayTimings(cTimer, dTimer, world.Step(), true)
}

func displayTimings(cTimer, dTimer life.Timer, step int64, toLog bool) {
	str := fmt.Sprintf("Calc %s | Disp %s", cTimer.ToString(), dTimer.ToString())
	fmt.Printf("#%d | %s\n", step, str)
	if toLog {
		log.Printf("Stable at %d passes: %s", step, str)
	}
}
