package day3

import "testing"

func TestDay_P1(t *testing.T) {
	tests := map[string]struct {
		wires    []wire
		expected int
	}{
		"firstExample": {
			wires: []wire{
				newWire([]string{"R8", "U5", "L5", "D3"}),
				newWire([]string{"U7", "R6", "D4", "L4"}),
			},
			expected: 6,
		},
		"secondExample": {
			wires: []wire{
				newWire([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}),
				newWire([]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}),
			},
			expected: 159,
		},
		"thirdExample": {
			wires: []wire{
				newWire([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}),
				newWire([]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}),
			},
			expected: 135,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{wires: test.wires}

			err := d.P1()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.closestInterestion != test.expected {
				t.Errorf("got %d, want %d", d.closestInterestion, test.expected)
			}
		})
	}
}

func TestDay_P2(t *testing.T) {
	tests := map[string]struct {
		wires    []wire
		expected int
	}{
		"firstExample": {
			wires: []wire{
				newWire([]string{"R8", "U5", "L5", "D3"}),
				newWire([]string{"U7", "R6", "D4", "L4"}),
			},
			expected: 30,
		},
		"secondExample": {
			wires: []wire{
				newWire([]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}),
				newWire([]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}),
			},
			expected: 610,
		},
		"thirdExample": {
			wires: []wire{
				newWire([]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}),
				newWire([]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}),
			},
			expected: 410,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{wires: test.wires}

			err := d.P2()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.closestInterestion != test.expected {
				t.Errorf("got %d, want %d", d.closestInterestion, test.expected)
			}
		})
	}
}
