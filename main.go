package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/heedson/aoc-2019/day1"
	"github.com/heedson/aoc-2019/day2"
	"github.com/heedson/aoc-2019/day3"
	"github.com/heedson/aoc-2019/day4"
	"github.com/heedson/aoc-2019/day5"
	"github.com/heedson/aoc-2019/day6"
	"github.com/heedson/aoc-2019/day7"
	"github.com/heedson/aoc-2019/day8"
)

type day interface {
	P1() error
	P2() error
}

var days = []day{
	day1.New(),
	day2.New(),
	day3.New(),
	day4.New(),
	day5.New(),
	day6.New(),
	day7.New(),
	day8.New(),
}

func main() {
	var dayNumber = flag.Int("day", 0, "Specify what days challenge to run.")
	var part = flag.Bool("final", false, "If set, indicates that the second part of the challenge should run.")
	flag.Parse()

	if *dayNumber == 0 {
		log.Fatal("A day needs to be selected to continue")
	}

	day, err := getDay(*dayNumber)
	if err != nil {
		log.Fatal(err)
	}

	err = runPart(day, *part)
	if err != nil {
		log.Fatal(err)
	}
}

func getDay(dayNumber int) (day, error) {
	if dayNumber > len(days) {
		return nil, fmt.Errorf("Day '%d' is not implemented (yet)", dayNumber)
	}

	fmt.Println("Day", dayNumber)

	return days[dayNumber-1], nil
}

func runPart(day day, finalPart bool) error {
	if finalPart {
		fmt.Println("Part 2")
		return day.P2()
	}

	fmt.Println("Part 1")
	return day.P1()
}
