package naloge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestEven(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{2, 4, 2, 5, 6, 6},
			expected: 2,
		},

		"unhappy path": {
			input:       []int{3, 5, 7, 9},
			expected:    0,
			errExpected: fmt.Errorf("no repeating even values"),
		},

		"empty path": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("no repeating even values"),
		},

		"decending path": {
			input:    []int{10, 10, 8, 8, 6},
			expected: 8,
		},

		"happy path.2": {
			input:    []int{1, 2, 3, 2, 4},
			expected: 2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SmallestEven(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {
				assert.Equal(t, test.expected, actual)
			}
		})
	}

}

func TestPrimeAndRepeating(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"basic primes repeated – 3 and 5": {
			input:    []int{2, 3, 3, 5, 5, 6},
			expected: []int{3, 5},
		},
		"no primes repeated": {
			input:    []int{4, 6, 8, 9},
			expected: []int{},
		},
		"2 and 3 repeated, 4 is not prime": {
			input:    []int{2, 2, 2, 3, 3, 4, 4},
			expected: []int{2, 3},
		},
		"only ones – not prime": {
			input:    []int{1, 1, 1},
			expected: []int{},
		},
		"empty input": {
			input:    []int{},
			expected: []int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := PrimeAndRepeating(test.input)

			assert.ElementsMatch(t, test.expected, actual)
		})

	}
}
