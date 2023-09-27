package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Point struct {
	x int
	y int
}

func parseLine(line string) Point {
	parts := strings.Split(line, " ")

	direction := parts[0]

	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("This should never, ever happen!")
	}

	if direction == "forward" {
		return Point{
			x: amount,
			y: 0,
		}
	} else if direction == "up" {
		return Point{
			x: 0,
			y: -amount,
		}
	}

	return Point{
		x: 0,
		y: amount,
	}

}

func main() {
	lines := strings.Split(getInput(), "\n")

	pos := Point{0, 0}

	for _, line := range lines {
		offset := parseLine(line)

		pos.x += offset.x
		pos.y += offset.y
	}

	fmt.Printf("Point: %+v", pos)
}
