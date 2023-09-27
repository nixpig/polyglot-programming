package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	lines := strings.Split(getInput(), "\n")

	treesHit := 0

	for r, line := range lines {
		if string(line[(r*3)%len(line)]) == "#" {
			treesHit++
		}
	}

	fmt.Printf("Number of trees hit: %v\n", treesHit)

}
