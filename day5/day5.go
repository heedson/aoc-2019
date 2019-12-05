package day5

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 5.
type Day struct {
	intCodes    []int
	input       int
	customInput bool
	output      int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day5/day5.txt")
	if err != nil {
		panic(err)
	}
	intCodes, err := util.IntsFromReader(reader, ",")
	if err != nil {
		panic(err)
	}
	return &Day{intCodes: intCodes}
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
	c := newComputer(d.intCodes)
	input := 1
	if d.customInput {
		input = d.input
	}
	err := c.run(input, false)
	if err != nil {
		return err
	}
	d.output = c.output
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	c := newComputer(d.intCodes)
	input := 5
	if d.customInput {
		input = d.input
	}
	err := c.run(input, false)
	if err != nil {
		return err
	}
	d.output = c.output
	return nil
}

type opcode int

// These are the opcodes available to this iteration of the computer from Day 2.
const (
	Add  opcode = 1
	Mul  opcode = 2
	In   opcode = 3
	Out  opcode = 4
	JIT  opcode = 5
	JIF  opcode = 6
	Lt   opcode = 7
	Eq   opcode = 8
	Halt opcode = 99
)

type computer struct {
	memory []int
	output int
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

func (c *computer) peekMode(paramVal, paramIdx, mode int) (int, error) {
	switch mode {
	case 0:
		return c.peek(paramVal)
	case 1:
		return c.peek(paramIdx)
	default:
		return 0, fmt.Errorf("parameter mode %d not recognised", mode)
	}
}

func (c *computer) run(input int, debug bool) error {
	for i := 0; i < len(c.memory); {
		opcode, paramModeMap := getOpcode(util.SplitInt(c.memory[i]))
		switch opcode {
		case Add:
			if i+3 >= len(c.memory) {
				return errors.New("not enough parameters for addition opcode")
			}
			if debug {
				log.Println("op: addition", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[i+2], i+2, paramModeMap[2])
			if err != nil {
				return err
			}
			err = c.poke(c.memory[i+3], v1+v2)
			if err != nil {
				return err
			}
			i += 4
		case Mul:
			if i+3 >= len(c.memory) {
				return errors.New("not enough parameters for multiplication opcode")
			}
			if debug {
				log.Println("op: multiplication", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[i+2], i+2, paramModeMap[2])
			if err != nil {
				return err
			}
			err = c.poke(c.memory[i+3], v1*v2)
			if err != nil {
				return err
			}
			i += 4
		case In:
			if i+1 >= len(c.memory) {
				return errors.New("not enough parameters for input opcode")
			}
			if debug {
				log.Println("op: input", "paramModeMap:", paramModeMap)
			}
			err := c.poke(c.memory[i+1], input)
			if err != nil {
				return err
			}
			i += 2
		case Out:
			if i+1 >= len(c.memory) {
				return errors.New("not enough parameters for output opcode")
			}
			if debug {
				log.Println("op: output", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			c.output = v1
			fmt.Println("Output:", v1)
			i += 2
		case JIT:
			if i+2 >= len(c.memory) {
				return errors.New("not enough parameters for jump-if-true opcode")
			}
			if debug {
				log.Println("op: jump-if-true", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			if v1 != 0 {
				v2, err := c.peekMode(c.memory[i+2], i+2, paramModeMap[2])
				if err != nil {
					return err
				}
				i = v2
			} else {
				i += 3
			}
		case JIF:
			if i+2 >= len(c.memory) {
				return errors.New("not enough parameters for jump-if-false opcode")
			}
			if debug {
				log.Println("op: jump-if-false", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			if v1 == 0 {
				v2, err := c.peekMode(c.memory[i+2], i+2, paramModeMap[2])
				if err != nil {
					return err
				}
				i = v2
			} else {
				i += 3
			}
		case Lt:
			if i+3 >= len(c.memory) {
				return errors.New("not enough parameters for less than opcode")
			}
			if debug {
				log.Println("op: less than", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[i+2], i+2, paramModeMap[2])
			if err != nil {
				return err
			}
			var res int
			if v1 < v2 {
				res = 1
			}
			err = c.poke(c.memory[i+3], res)
			if err != nil {
				return err
			}
			i += 4
		case Eq:
			if i+3 >= len(c.memory) {
				return errors.New("not enough parameters for equal opcode")
			}
			if debug {
				log.Println("op: equals", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[i+1], i+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[i+2], i+2, paramModeMap[2])
			if err != nil {
				return err
			}
			var res int
			if v1 == v2 {
				res = 1
			}
			err = c.poke(c.memory[i+3], res)
			if err != nil {
				return err
			}
			i += 4
		case Halt:
			if debug {
				log.Println("op: halt", "paramModeMap:", paramModeMap)
			}
			i++
			return nil
		}
	}
	return nil
}

func getOpcode(opcodeParts []int) (opcode, map[int]int) {
	paramModeMap := make(map[int]int)
	if len(opcodeParts) == 1 {
		return opcode(opcodeParts[0]), paramModeMap
	}
	if len(opcodeParts) == 2 {
		return opcode(10*opcodeParts[0] + opcodeParts[1]), paramModeMap
	}
	op := 10*opcodeParts[len(opcodeParts)-2] + opcodeParts[len(opcodeParts)-1]
	for i, paramMode := range opcodeParts[:len(opcodeParts)-2] {
		paramModeMap[len(opcodeParts)-2-i] = paramMode
	}
	return opcode(op), paramModeMap
}
