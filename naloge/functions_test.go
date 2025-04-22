package naloge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    bool
		errExpected error
	}{
		"is prime": {
			input:    13,
			expected: true,
		},
		"is not prime": {
			input:    6,
			expected: false,
		},
		"zero": {
			input:    0,
			expected: false,
		},
		"one": {
			input:    1,
			expected: true,
		},
		"negative number": {
			input:       -5,
			expected:    false,
			errExpected: fmt.Errorf("negative number: %d", -5),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IsPrime(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
func TestMaxInSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 22,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 10,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -3,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3, 3},
			expected: 3,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MaxInSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMinInSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 1,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: -5,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -15,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3, 3},
			expected: 3,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MinInSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSumSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 67,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 3,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -40,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3, 3},
			expected: 15,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SumSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestAverageSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 9,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 0,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -8,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3, 3},
			expected: 3,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := AverageSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}
