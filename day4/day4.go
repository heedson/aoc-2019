package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 4.
type Day struct {
	lower      int
	upper      int
	totalCount int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day4/day4.txt")
	if err != nil {
		panic(err)
	}
	bounds, err := util.StringsFromReader(reader, "-")
	if err != nil {
		panic(err)
	}
	if len(bounds) != 2 {
		panic(fmt.Sprintf("expected 2 bounds; got %d", len(bounds)))
	}
	bounds[1] = strings.Trim(bounds[1], "\n")
	lower, err := strconv.Atoi(bounds[0])
	if err != nil {
		panic(err)
	}
	upper, err := strconv.Atoi(bounds[1])
	if err != nil {
		panic(err)
	}
	return &Day{
		lower: lower,
		upper: upper,
	}
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
	d.totalCount = 0
	for i := d.lower; i < d.upper; i++ {
		digits := getDigits(i)
		if criteriaMet(digits) {
			d.totalCount++
		}
	}
	fmt.Println("Total possible passwords:", d.totalCount)
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	d.totalCount = 0
	for i := d.lower; i < d.upper; i++ {
		digits := getDigits(i)
		if criteriaMetP2(digits) {
			d.totalCount++
		}
	}
	fmt.Println("Total possible passwords:", d.totalCount)
	return nil
}

func getDigits(password int) []int {
	return getAllDigits(password, nil)
}

func getAllDigits(remainingPassword int, digits []int) []int {
	if remainingPassword != 0 {
		digit := remainingPassword % 10
		digits = append([]int{digit}, digits...)
		return getAllDigits(remainingPassword/10, digits)
	}
	return digits
}

func criteriaMet(passwordDigits []int) bool {
	var hasDouble bool
	for i, d := range passwordDigits[:len(passwordDigits)-1] {
		if d > passwordDigits[i+1] {
			return false
		}
		if d == passwordDigits[i+1] {
			hasDouble = true
		}
	}
	return hasDouble
}

func criteriaMetP2(passwordDigits []int) bool {
	var hasDouble bool
	var repeatCount int
	prevDigit := -1
	for _, d := range passwordDigits {
		if d < prevDigit {
			return false
		}
		if d == prevDigit {
			repeatCount++
		} else {
			if repeatCount == 1 {
				hasDouble = true
			}
			repeatCount = 0
		}
		prevDigit = d
	}
	if repeatCount == 1 {
		hasDouble = true
	}
	return hasDouble
}
