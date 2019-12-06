package day6

import "testing"

func TestDay_P1(t *testing.T) {
	tests := map[string]struct {
		orbits   []string
		expected int
	}{
		"firstExample": {
			orbits: []string{
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
			},
			expected: 42,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{orbits: test.orbits}

			err := d.P1()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.totalOrbits != test.expected {
				t.Errorf("got %d, want %d", d.totalOrbits, test.expected)
			}
		})
	}
}

func TestDay_P2(t *testing.T) {
	tests := map[string]struct {
		orbits   []string
		expected int
	}{
		"firstExample": {
			orbits: []string{
				"COM)B",
				"B)C",
				"C)D",
				"D)E",
				"E)F",
				"B)G",
				"G)H",
				"D)I",
				"E)J",
				"J)K",
				"K)L",
				"K)YOU",
				"I)SAN",
			},
			expected: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{orbits: test.orbits}

			err := d.P2()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.santaDistance != test.expected {
				t.Errorf("got %d, want %d", d.santaDistance, test.expected)
			}
		})
	}
}
