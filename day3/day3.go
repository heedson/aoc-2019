package day3

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 3.
type Day struct {
	wires              []wire
	closestInterestion int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day3/day3.txt")
	if err != nil {
		panic(err)
	}
	wires, err := util.StringsFromReader(reader, "\n")
	if err != nil {
		panic(err)
	}
	d := &Day{wires: make([]wire, 0, len(wires))}
	for _, w := range wires {
		d.wires = append(d.wires, newWire(strings.Split(w, ",")))
	}
	return d
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
	var closestManDist int
	for i, w1 := range d.wires {
		for j, w2 := range d.wires[i:] {
			if j == 0 {
				continue
			}
			intersections := w1.intersections(w2)
			sort.Slice(intersections, func(i, j int) bool {
				// Put the 0,0 intersections at the end.
				leftManDist := intersections[i].manDist(newVec2(0, 0))
				if leftManDist == 0 {
					return false
				}
				rightManDist := intersections[j].manDist(newVec2(0, 0))
				if rightManDist == 0 {
					return true
				}
				return leftManDist < rightManDist
			})
			if len(intersections) != 0 {
				intersectionManDist := intersections[0].manDist(newVec2(0, 0))
				if closestManDist == 0 {
					closestManDist = intersectionManDist
				}
				if intersectionManDist < closestManDist {
					closestManDist = intersectionManDist
				}
			}
		}
	}
	d.closestInterestion = closestManDist
	fmt.Println("Closest intersection:", d.closestInterestion)
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	var closestLenDist int
	for i, w1 := range d.wires {
		for j, w2 := range d.wires[i:] {
			if j == 0 {
				continue
			}
			intersections := w1.intersections(w2)
			sort.Slice(intersections, func(i, j int) bool {
				// Put the 0,0 intersections at the end.
				leftManDist := intersections[i].manDist(newVec2(0, 0))
				if leftManDist == 0 {
					return false
				}
				rightManDist := intersections[j].manDist(newVec2(0, 0))
				if rightManDist == 0 {
					return true
				}
				return intersections[i].lengthToPoint < intersections[j].lengthToPoint
			})
			if len(intersections) != 0 {
				intersectionLenDist := intersections[0].lengthToPoint
				if closestLenDist == 0 {
					closestLenDist = intersectionLenDist
				}
				if intersectionLenDist < closestLenDist {
					closestLenDist = intersectionLenDist
				}
			}
		}
	}
	d.closestInterestion = closestLenDist
	fmt.Println("Closest intersection:", d.closestInterestion)
	return nil
}

type wire struct {
	head    vec2
	headLen int
	cells   map[vec2]int
}

func newWire(directions []string) wire {
	w := wire{cells: make(map[vec2]int), head: newVec2(0, 0)}
	for _, dir := range directions {
		dist, err := strconv.Atoi(dir[1:])
		if err != nil {
			panic(err)
		}
		var newHead vec2
		newHeadLen := w.headLen
		switch dir[0] {
		case 'U':
			for y := 0; y < dist; y++ {
				newHead = newVec2(w.head.x, w.head.y+y+1)
				newHeadLen++
				w.cells[newHead] = newHeadLen
			}
		case 'R':
			for x := 0; x < dist; x++ {
				newHead = newVec2(w.head.x+x+1, w.head.y)
				newHeadLen++
				w.cells[newHead] = newHeadLen
			}
		case 'D':
			for y := 0; y < dist; y++ {
				newHead = newVec2(w.head.x, w.head.y-y-1)
				newHeadLen++
				w.cells[newHead] = newHeadLen
			}
		case 'L':
			for x := 0; x < dist; x++ {
				newHead = newVec2(w.head.x-x-1, w.head.y)
				newHeadLen++
				w.cells[newHead] = newHeadLen
			}
		}
		w.head = newHead
		w.headLen = newHeadLen
	}
	return w
}

func (w *wire) intersections(other wire) []intersection {
	var intersections []intersection
	for k, otherLen := range other.cells {
		if ourLen, ok := w.cells[k]; ok {
			intersections = append(intersections, intersection{
				vec2:          k,
				lengthToPoint: ourLen + otherLen,
			})
		}
	}
	return intersections
}

type intersection struct {
	vec2
	lengthToPoint int
}

type vec2 struct {
	x int
	y int
}

func newVec2(x, y int) vec2 {
	return vec2{
		x: x,
		y: y,
	}
}

func (v vec2) add(other vec2) vec2 {
	return newVec2(v.x+other.x, v.y+other.y)
}

func (v vec2) equals(other vec2) bool {
	return v.x == other.x && v.y == other.y
}

func (v vec2) manDist(other vec2) int {
	xComp := v.x - other.x
	if xComp < 0 {
		xComp *= -1
	}
	yComp := v.y - other.y
	if yComp < 0 {
		yComp *= -1
	}
	return xComp + yComp
}
