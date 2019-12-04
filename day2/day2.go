package day2

import (
	"errors"
	"fmt"
	"os"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 2.
type Day struct {
	intCodes []int
	pos0Val  int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day2/day2.txt")
	if err != nil {
		panic(err)
	}
	intCodes, err := util.IntsFromReader(reader, ",")
	if err != nil {
		panic(err)
	}
	return &Day{
		intCodes: intCodes,
	}
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
parseLoop:
	for i, o := range d.intCodes {
		if i%4 != 0 {
			continue
		}
		switch o {
		case 1:
			err := validatePointers(i, 4, d.intCodes)
			if err != nil {
				return err
			}
			d.intCodes[d.intCodes[i+3]] = d.intCodes[d.intCodes[i+1]] + d.intCodes[d.intCodes[i+2]]
		case 2:
			err := validatePointers(i, 4, d.intCodes)
			if err != nil {
				return err
			}
			d.intCodes[d.intCodes[i+3]] = d.intCodes[d.intCodes[i+1]] * d.intCodes[d.intCodes[i+2]]
		case 99:
			break parseLoop
		}
	}
	d.pos0Val = d.intCodes[0]
	fmt.Println("Index 0:", d.pos0Val)
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			c := newComputer(d.intCodes)
			err := c.poke(1, noun)
			if err != nil {
				return err
			}
			err = c.poke(2, verb)
			if err != nil {
				return err
			}
			err = c.run()
			if err != nil {
				return err
			}
			v, err := c.peek(0)
			if err != nil {
				return err
			}
			if v == 19690720 {
				fmt.Println("Answer:", 100*noun+verb)
				return nil
			}
		}
	}
	return fmt.Errorf("didn't find a noun verb combination to make %d", 19690720)
}

func validatePointers(instructionPointer, instructionSize int, intCodes []int) error {
	maxIdx := len(intCodes) - 1
	if instructionPointer+instructionSize > maxIdx {
		return errors.New("not enough intcodes for instruction")
	}
	for i := 0; i < instructionSize; i++ {
		err := validatePointer(intCodes[instructionPointer+i], maxIdx)
		if err != nil {
			return err
		}
	}
	return nil
}

func validatePointer(idx, maxIdx int) error {
	if idx < 0 || idx > maxIdx {
		return fmt.Errorf("%d is pointing to outside of the intcode list", idx)
	}
	return nil
}

type computer struct {
	memory []int
}

func newComputer(initialMemory []int) computer {
	c := computer{memory: make([]int, len(initialMemory))}
	copy(c.memory, initialMemory)
	return c
}

func (c *computer) poke(address, value int) error {
	if address < 0 || address >= len(c.memory) {
		return fmt.Errorf("address %d outside of memory", address)
	}
	c.memory[address] = value
	return nil
}

func (c *computer) peek(address int) (int, error) {
	if address < 0 || address >= len(c.memory) {
		return 0, fmt.Errorf("address %d outside of memory", address)
	}
	return c.memory[address], nil
}

func (c *computer) run() error {
	for i := 0; i < len(c.memory); {
		switch c.memory[i] {
		case 1:
			if i+3 >= len(c.memory) {
				return errors.New("not enough parameters for addition opcode")
			}
			v1, err := c.peek(c.memory[i+1])
			if err != nil {
				return err
			}
			v2, err := c.peek(c.memory[i+2])
			if err != nil {
				return err
			}
			err = c.poke(c.memory[i+3], v1+v2)
			if err != nil {
				return err
			}
			i += 4
		case 2:
			if i+3 >= len(c.memory) {
				return errors.New("not enough parameters for multiplication opcode")
			}
			v1, err := c.peek(c.memory[i+1])
			if err != nil {
				return err
			}
			v2, err := c.peek(c.memory[i+2])
			if err != nil {
				return err
			}
			err = c.poke(c.memory[i+3], v1*v2)
			if err != nil {
				return err
			}
			i += 4
		case 99:
			i++
			return nil
		}
	}
	return nil
}
