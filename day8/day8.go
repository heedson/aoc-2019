package day8

import (
	"fmt"
	"math"
	"os"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 8.
type Day struct {
	width  int
	height int
	pixels []int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day8/day8.txt")
	if err != nil {
		panic(err)
	}
	pixels, err := util.IntsFromReader(reader, "")
	if err != nil {
		panic(err)
	}
	return &Day{
		pixels: pixels,
		width:  25,
		height: 6,
	}
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
	colours := []string{" ", "▒", "█", "?"}
	pixelLen := len(d.pixels)
	layers := make([]map[int]int, 0, pixelLen/(d.height*d.width))
	var pixelIdx int
	for pixelIdx < pixelLen {
		layer := make(map[int]int)
		for y := 0; y < d.height; y++ {
			for x := 0; x < d.width; x++ {
				colourIdx := d.pixels[pixelIdx]
				if colourIdx > 3 {
					colourIdx = 3
				}
				layer[colourIdx]++
				fmt.Print(colours[colourIdx])
				pixelIdx++
			}
			fmt.Print("\n")
		}
		layers = append(layers, layer)
		fmt.Print("Layer:", layer, "\n")
	}

	leastZeroCount := math.MaxInt64
	var result int
	for _, l := range layers {
		if l[0] < leastZeroCount {
			leastZeroCount = l[0]
			result = l[1] * l[2]
		}
	}
	fmt.Println("Result:", result)

	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	return nil
}
