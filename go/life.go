package main

import (
    "./life"
    "fmt"
)

func itLives() {
    var cTimer, dTimer life.Timer
    var world life.World

    world.Init(true)
    fmt.Print(world.ToString())
    displayTimings(cTimer, dTimer, world.Step())

    for {
        cTimer.Start()
        world.Calculate()
        cTimer.Stop()

        dTimer.Start()
        fmt.Print(world.ToString())
        dTimer.Stop()

        displayTimings(cTimer, dTimer, world.Step())
    }
}

func displayTimings(cTimer, dTimer life.Timer, step int64) {
    fmt.Printf("#%d | Calc %s | Disp %s\n", step, cTimer.ToString(), dTimer.ToString())
}

func main() {
    itLives()
}
