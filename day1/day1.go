package day1

import (
	"fmt"
	"math"
	"os"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 1.
type Day struct {
	masses    []int
	totalFuel int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day1/day1.txt")
	if err != nil {
		panic(err)
	}
	masses, err := util.IntsFromReader(reader, "\n")
	if err != nil {
		panic(err)
	}
	return &Day{
		masses: masses,
	}
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
	d.totalFuel = 0
	for _, mass := range d.masses {
		d.totalFuel += calculatP1Fuel(mass)
	}
	fmt.Println("Total fuel:", d.totalFuel)
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	d.totalFuel = 0
	for _, mass := range d.masses {
		d.totalFuel += calculateAllP2Fuel(mass)
	}
	fmt.Println("Total fuel:", d.totalFuel)
	return nil
}

func calculatP1Fuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func calculateAllP2Fuel(mass int) int {
	return calculateP2Fuel(mass, 0)
}

func calculateP2Fuel(mass, totalFuel int) int {
	fuel := calculatP1Fuel(mass)
	if fuel <= 0 {
		return totalFuel
	}
	return calculateP2Fuel(fuel, totalFuel+fuel)
}
