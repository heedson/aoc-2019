package day2

import "testing"

func TestDay_P1(t *testing.T) {
	tests := map[string]struct {
		intCodes []int
		expected int
	}{
		"firstExample": {
			intCodes: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expected: 3500,
		},
		"secondExample": {
			intCodes: []int{1, 0, 0, 0, 99},
			expected: 2,
		},
		"thirdExample": {
			intCodes: []int{2, 3, 0, 3, 99},
			expected: 2,
		},
		"fourthExample": {
			intCodes: []int{2, 4, 4, 5, 99, 0},
			expected: 2,
		},
		"fifthExample": {
			intCodes: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: 30,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{intCodes: test.intCodes}

			err := d.P1()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.pos0Val != test.expected {
				t.Errorf("got %d, want %d", d.pos0Val, test.expected)
			}
		})
	}
}
