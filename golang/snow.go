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
	height = height - 1
	width := tm.Width()
	flakes := []Flake{}
	rand.Seed(2020)
	floorState := make([]int, width)

	for {
		tm.Clear()

		// Update the flake pos
		for i := 0; i < len(flakes); i++ {
			if flakes[i].Lock == false {
				if flakes[i].Y >= height-floorState[flakes[i].X]+1 {
					fmt.Println("Lock!")
					floorState[flakes[i].X] += 1
					flakes[i].Lock = true
				} else {
					flakes[i].Y = flakes[i].Y + 1
				}
			}
		}

		newFlake := Flake{X: rand.Intn(width), Y: 0, Lock: false}

		flakes = append(flakes, newFlake)

		for _, f := range flakes {
			tm.Print(tm.MoveTo("*", f.X, f.Y))
			tm.MoveCursor(1, 1)
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
