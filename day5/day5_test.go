package day5

import "testing"

func TestDay_P2(t *testing.T) {
	tests := map[string]struct {
		intCodes []int
		input    int
		expected int
	}{
		"firstExample-equal": {
			intCodes: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    8,
			expected: 1,
		},
		"firstExample-notEqual": {
			intCodes: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    5,
			expected: 0,
		},
		"secondExample-less": {
			intCodes: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    5,
			expected: 1,
		},
		"secondExample-notLess": {
			intCodes: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:    8,
			expected: 0,
		},
		"thirdExample-equal": {
			intCodes: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    8,
			expected: 1,
		},
		"thirdExample-notEqual": {
			intCodes: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:    5,
			expected: 0,
		},
		"fourthExample-less": {
			intCodes: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    5,
			expected: 1,
		},
		"fourthExample-notLess": {
			intCodes: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:    8,
			expected: 0,
		},
		"fifthExample-zero": {
			intCodes: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    0,
			expected: 0,
		},
		"fifthExample-nonZero": {
			intCodes: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:    5,
			expected: 1,
		},
		"sixthExample-zero": {
			intCodes: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    0,
			expected: 0,
		},
		"sixthExample-nonZero": {
			intCodes: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:    5,
			expected: 1,
		},
		"seventhExample-equal": {
			intCodes: []int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			},
			input:    8,
			expected: 1000,
		},
		"seventhExample-less": {
			intCodes: []int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			},
			input:    5,
			expected: 999,
		},
		"seventhExample-greater": {
			intCodes: []int{
				3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
			},
			input:    11,
			expected: 1001,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := Day{intCodes: test.intCodes, input: test.input, customInput: true}

			err := d.P2()
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if d.output != test.expected {
				t.Errorf("got %d, want %d", d.output, test.expected)
			}
		})
	}
}
