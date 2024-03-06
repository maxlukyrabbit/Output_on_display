package main

import (
	"fmt"
	"math/rand"
	"time"
	"bytes"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

func (w *World) Neighbours(x, y int) int {
	count := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			// Adjusting for torus topology
			adjustedI := (i + w.Height) % w.Height
			adjustedJ := (j + w.Width) % w.Width
			if w.Cells[adjustedI][adjustedJ] {
				count++
			}
		}
	}
	return count
}



func (w *World) Next(x, y int) bool {
	n := w.Neighbours(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}
	if n == 3 && !alive {
		return true
	}
	return false
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) Seed() {
	for _, row := range w.Cells {
		for i := range row {
			if rand.Intn(10) == 1 {
				row[i] = true
			}
		}
	}
}

func (w *World) String() string {
    var buffer bytes.Buffer
    for i := 0; i < w.Height; i++ {
        for j := 0; j < w.Width; j++ {
            if w.Cells[i][j] {
                buffer.WriteString("\033[48;5;130m \033[0m") // Зеленый цвет для живых клеток
            } else {
                buffer.WriteString("\033[48;5;52m \033[0m") // Коричневый цвет для мертвых клеток
            }
        }
        buffer.WriteString("\n")
    }
    return buffer.String()
}


func main() {
	height := 20
	width := 40
	currentWorld := NewWorld(height, width)
	nextWorld := NewWorld(height, width)
	currentWorld.Seed()
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println(currentWorld)
		NextState(currentWorld, nextWorld)
		currentWorld, nextWorld = nextWorld, currentWorld
		time.Sleep(100 * time.Millisecond)
	}
}

