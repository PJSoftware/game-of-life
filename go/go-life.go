package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"./life"
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
	fmt.Print(world.Render(false))
	displayTimings(cTimer, dTimer, world.Step(), false)

	for world.Step() < 1000 || !cTimer.IsAverageStable(20) {
		cTimer.Start()
		world.Calculate()
		cTimer.Stop()

		dTimer.Start()
		fmt.Print(world.Render(false))
		dTimer.Stop()

		displayTimings(cTimer, dTimer, world.Step(), false)
		time.Sleep(animDelay(cTimer, dTimer))
	}
	displayTimings(cTimer, dTimer, world.Step(), true)
}

func displayTimings(cTimer, dTimer life.Timer, step int64, toLog bool) {
	str := fmt.Sprintf("Calc %s | Disp %s", cTimer.String(), dTimer.String())
	fmt.Printf("Go: #%d | %s\n", step, str)
	if toLog {
		log.Printf("Stable at %d passes: %s", step, str)
	}
}

func animDelay(cTimer, dTimer life.Timer) time.Duration {
	maxDelay := 50.0 // 20 FPS -> 50 ms per frame
	minDelay := 10.0 // in case our calcs are too slow
	loopDuration := cTimer.ElapsedMS() + dTimer.ElapsedMS()

	msDelay := math.Max(maxDelay-loopDuration, minDelay)
	return time.Millisecond * time.Duration(msDelay)
}
