package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Flake struct {
	X    int
	Y    int
	Lock bool
}

func init() {
	InterruptHandler()
}

func main() {
	height := tm.Height()
	width := tm.Width()
	flakes := []Flake{}
	rand.Seed(2020)
	var floorState [width]int

	tm.Clear()
	for {
		// Update the flake pos
		for _, f := range flakes {
			if f.Lock == false {

				if f.Y == height-floorState[f.X]+1 {
					floorState[f.X] += 1
					f.Lock = true
				} else {
					f.Y += 1
				}
			}
		}

		newFlake := Flake{X: rand.Intn(width), Y: 0, Lock: false}

		flakes = append(flakes, newFlake)

		tm.MoveCursor(1, 1)
		for _, f := range flakes {
			tm.Print(tm.MoveTo("*", f.X, f.Y))
		}
		tm.Flush()
		time.Sleep(time.Second)
	}

}

func InterruptHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
