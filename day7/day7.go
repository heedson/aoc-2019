package day7

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 7.
type Day struct {
	intCodes      []int
	maxThrust     int
	maxPhaseOrder []int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day7/day7.txt")
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
	d.maxThrust = 0
	numAmps := 5
	phases := make([]int, numAmps)
	for i := 0; i < numAmps; i++ {
		phases[i] = i
	}

	d.maxThrust, d.maxPhaseOrder = perm(d.intCodes, phases, runAmplifiers, 0)

	fmt.Println("Max thrust:", d.maxThrust)
	fmt.Println("Phase order:", d.maxPhaseOrder)
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	d.maxThrust = 0
	numAmps := 5
	phases := make([]int, numAmps)
	for i := 5; i < numAmps+5; i++ {
		phases[i-5] = i
	}

	d.maxThrust, d.maxPhaseOrder = perm(d.intCodes, phases, runAmplifiers, 0)

	fmt.Println("Max thrust:", d.maxThrust)
	fmt.Println("Phase order:", d.maxPhaseOrder)
	return nil
}

func perm(intCodes, phases []int, f func([]int, []int) (int, []int), i int) (int, []int) {
	if i > len(phases) {
		return f(intCodes, phases)
	}
	maxValue, maxPhases := perm(intCodes, phases, f, i+1)
	for j := i + 1; j < len(phases); j++ {
		phases[i], phases[j] = phases[j], phases[i]
		v, vPhases := perm(intCodes, phases, f, i+1)
		if v > maxValue {
			maxValue = v
			maxPhases = vPhases
		}
		phases[i], phases[j] = phases[j], phases[i]
	}
	return maxValue, maxPhases
}

func runAmplifiers(intCodes, phases []int) (int, []int) {
	amps := make([]computer, 0, len(phases))
	for _, p := range phases {
		c := newComputer(intCodes)
		c.input(p)
		amps = append(amps, c)
	}
	var haltCount int
	var output int
feedback:
	for {
		for i := range amps {
			amps[i].input(output)
			err := amps[i].run(false)
			if err != nil && err != errHalted {
				panic(err)
			}
			if err == errHalted {
				haltCount++
			}
			output = amps[i].output
			if haltCount == len(phases) {
				break feedback
			}
		}
	}
	return output, phases
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

var errHalted = errors.New("received halt code")

type computer struct {
	memory     []int
	programIdx int
	inputIdx   int
	inputs     []int
	output     int
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

func (c *computer) input(i int) {
	c.inputs = append(c.inputs, i)
}

func (c *computer) run(debug bool) error {
	for c.programIdx < len(c.memory) {
		opcode, paramModeMap := getOpcode(util.SplitInt(c.memory[c.programIdx]))
		switch opcode {
		case Add:
			if c.programIdx+3 >= len(c.memory) {
				return errors.New("not enough parameters for addition opcode")
			}
			if debug {
				log.Println("op: addition", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[c.programIdx+2], c.programIdx+2, paramModeMap[2])
			if err != nil {
				return err
			}
			err = c.poke(c.memory[c.programIdx+3], v1+v2)
			if err != nil {
				return err
			}
			c.programIdx += 4
		case Mul:
			if c.programIdx+3 >= len(c.memory) {
				return errors.New("not enough parameters for multiplication opcode")
			}
			if debug {
				log.Println("op: multiplication", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[c.programIdx+2], c.programIdx+2, paramModeMap[2])
			if err != nil {
				return err
			}
			err = c.poke(c.memory[c.programIdx+3], v1*v2)
			if err != nil {
				return err
			}
			c.programIdx += 4
		case In:
			if c.programIdx+1 >= len(c.memory) {
				return errors.New("not enough parameters for input opcode")
			}
			if c.inputIdx >= len(c.inputs) {
				return errors.New("not enough input given")
			}
			if debug {
				log.Println("op: input", "paramModeMap:", paramModeMap)
			}
			err := c.poke(c.memory[c.programIdx+1], c.inputs[c.inputIdx])
			if err != nil {
				return err
			}
			c.inputIdx++
			c.programIdx += 2
		case Out:
			if c.programIdx+1 >= len(c.memory) {
				return errors.New("not enough parameters for output opcode")
			}
			if debug {
				log.Println("op: output", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			c.output = v1
			if debug {
				fmt.Println("Output:", v1)
			}
			c.programIdx += 2
			return nil
		case JIT:
			if c.programIdx+2 >= len(c.memory) {
				return errors.New("not enough parameters for jump-if-true opcode")
			}
			if debug {
				log.Println("op: jump-if-true", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			if v1 != 0 {
				v2, err := c.peekMode(c.memory[c.programIdx+2], c.programIdx+2, paramModeMap[2])
				if err != nil {
					return err
				}
				c.programIdx = v2
			} else {
				c.programIdx += 3
			}
		case JIF:
			if c.programIdx+2 >= len(c.memory) {
				return errors.New("not enough parameters for jump-if-false opcode")
			}
			if debug {
				log.Println("op: jump-if-false", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			if v1 == 0 {
				v2, err := c.peekMode(c.memory[c.programIdx+2], c.programIdx+2, paramModeMap[2])
				if err != nil {
					return err
				}
				c.programIdx = v2
			} else {
				c.programIdx += 3
			}
		case Lt:
			if c.programIdx+3 >= len(c.memory) {
				return errors.New("not enough parameters for less than opcode")
			}
			if debug {
				log.Println("op: less than", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[c.programIdx+2], c.programIdx+2, paramModeMap[2])
			if err != nil {
				return err
			}
			var res int
			if v1 < v2 {
				res = 1
			}
			err = c.poke(c.memory[c.programIdx+3], res)
			if err != nil {
				return err
			}
			c.programIdx += 4
		case Eq:
			if c.programIdx+3 >= len(c.memory) {
				return errors.New("not enough parameters for equal opcode")
			}
			if debug {
				log.Println("op: equals", "paramModeMap:", paramModeMap)
			}
			v1, err := c.peekMode(c.memory[c.programIdx+1], c.programIdx+1, paramModeMap[1])
			if err != nil {
				return err
			}
			v2, err := c.peekMode(c.memory[c.programIdx+2], c.programIdx+2, paramModeMap[2])
			if err != nil {
				return err
			}
			var res int
			if v1 == v2 {
				res = 1
			}
			err = c.poke(c.memory[c.programIdx+3], res)
			if err != nil {
				return err
			}
			c.programIdx += 4
		case Halt:
			if debug {
				log.Println("op: halt", "paramModeMap:", paramModeMap)
			}
			c.programIdx++
			return errHalted
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
