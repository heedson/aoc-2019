package day1

import "testing"

func TestDay_P1(t *testing.T) {
	tests := map[string]struct {
		masses   []int
		expected int
	}{
		"firstExample": {
			masses:   []int{12},
			expected: 2,
		},
		"secondExample": {
			masses:   []int{14},
			expected: 2,
		},
		"thirdExample": {
			masses:   []int{1969},
			expected: 654,
		},
		"fourthExample": {
			masses:   []int{100756},
			expected: 33583,
		},
		"addition": {
			masses:   []int{12, 14, 1969},
			expected: 658,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{masses: test.masses}

			err := d.P1()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.totalFuel != test.expected {
				t.Errorf("got %d, want %d", d.totalFuel, test.expected)
			}
		})
	}
}

func TestDay_P2(t *testing.T) {
	tests := map[string]struct {
		masses   []int
		expected int
	}{
		"firstExample": {
			masses:   []int{14},
			expected: 2,
		},
		"secondExample": {
			masses:   []int{1969},
			expected: 966,
		},
		"thirdExample": {
			masses:   []int{100756},
			expected: 50346,
		},
		"addition": {
			masses:   []int{14, 1969, 100756},
			expected: 51314,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{masses: test.masses}

			err := d.P2()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.totalFuel != test.expected {
				t.Errorf("got %d, want %d", d.totalFuel, test.expected)
			}
		})
	}
}
