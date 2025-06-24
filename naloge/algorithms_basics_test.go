package naloge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRepeatingEven(t *testing.T) {
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
			actual, err := SmallestRepeatingEven(test.input)
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

func TestAverageOddIndiceDivByThree(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    float64
		errExpected error
	}{
		"happy path": {
			input:    []int{3, 6, 9, 12, 15, 18},
			expected: 12.0,
		},
		"unhappy repeated": {
			input:       []int{1, 2, 3, 4},
			expected:    0.0,
			errExpected: fmt.Errorf("no matching values at odd indices"),
		},
		"3 repeated": {
			input:    []int{3, 3, 3, 3},
			expected: 3.0,
		},
		"happy path 2.0": {
			input:    []int{5, 6, 7, 9, 11, 12},
			expected: 9.0,
		},
		"empty input": {
			input:       []int{},
			expected:    0.0,
			errExpected: fmt.Errorf("no matching values at odd indices"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := AverageOddIndiceDivByThree(test.input)
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

func TestOddPalindromes(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"happy path": {
			input:    []int{1, 3, 5, 7, 9, 11, 33, 44},
			expected: 7,
		},
		"no odds": {
			input:    []int{2, 4, 6, 8},
			expected: 0,
		},
		"100s palindromes": {
			input:    []int{121, 131, 141, 151},
			expected: 4,
		},
		"empty path": {
			input:    []int{},
			expected: 0,
		},
		"repeating values": {
			input:    []int{1, 1, 1, 1},
			expected: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := OddPalindromes(test.input)

			assert.Equal(t, test.expected, actual)
		})

	}
}

func TestSecondLargestRepeatingValue(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{

		"happy path": {
			input:    []int{5, 5, 6, 6, 7, 7},
			expected: 6,
		},
		"unhappy repeated": {
			input:       []int{1, 1, 2, 3, 4, 5},
			expected:    0,
			errExpected: fmt.Errorf("not enough repeating values"),
		},
		"reverse path": {
			input:    []int{9, 9, 8, 8, 7, 7, 6, 6},
			expected: 8,
		},
		"happy path 2.0": {
			input:    []int{2, 2, 2, 1, 1},
			expected: 1,
		},
		"empty input": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("not enough repeating values"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SecondLargestRepeatingValue(test.input)
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

func TestLongestIncreasingSubSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    []int
		errExpected error
	}{

		"happy path": {
			input:    []int{1, 2, 3, 2, 3, 4, 5},
			expected: []int{2, 3, 4, 5},
		},
		"unhappy repeated": {
			input:    []int{10},
			expected: []int{10},
		},
		"reverse path": {
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5},
		},
		"happy path 2.0": {
			input:    []int{1, 2, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
		"empty input": {
			input:       []int{},
			expected:    []int{},
			errExpected: fmt.Errorf("empty input"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := LongestIncreasingSubSlice(test.input)
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

func TestPerfectSquareOddNumberProduct(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{

		"happy path": {
			input:    []int{1, 4, 9, 16, 25},
			expected: 225,
		},
		"zero matching values path": {
			input:       []int{2, 3, 5, 7},
			expected:    0,
			errExpected: fmt.Errorf("no matching values"),
		},
		"repeating path": {
			input:    []int{1, 1, 1},
			expected: 1,
		},
		"repeating path 2.0": {
			input:    []int{9, 9, 9},
			expected: 729,
		},
		"happy path 2.0": {
			input:    []int{49, 36, 25},
			expected: 1225,
		},
		"empty input": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("no matching values"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := PerfectSquareOddNumberProduct(test.input)
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

func TestFirstMultOfThreeExactlyTwice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{

		"happy path": {
			input:    []int{3, 3, 6, 6, 9, 9},
			expected: 3,
		},
		"zero matching values path": {
			input:       []int{2, 4, 6, 8, 10},
			expected:    0,
			errExpected: fmt.Errorf("no valid match found"),
		},
		"happy path 2.0": {
			input:    []int{6, 6, 3, 3},
			expected: 6,
		},
		"repeating path": {
			input:    []int{6, 6, 6, 3, 3},
			expected: 3,
		},
		"happy path 3.0": {
			input:    []int{12, 12, 15},
			expected: 12,
		},
		"empty input": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("no valid match found"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := FirstMultOfThreeExactlyTwice(test.input)
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
