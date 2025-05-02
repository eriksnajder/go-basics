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

func TestCountEven(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 2,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 2,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: 1,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3, 3},
			expected: 0,
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
			actual, err := CountEven(test.input)
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
func TestCountOdd(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 5,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 3,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: 4,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3, 3},
			expected: 5,
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
			actual, err := CountOdd(test.input)
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

func TestReverseString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"happy path": {
			input:    "pork",
			expected: "krop",
		},
		"contains negative ints": {
			input:    "abeceda",
			expected: "adeceba",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ReverseString(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"happy path": {
			input:    "pericarezeracirep",
			expected: true,
		},
		"unhappy path": {
			input:    "abeceda",
			expected: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsPalindrome(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := map[string]struct {
		nums        []int
		target      int
		expected    int
		errExpected error
	}{
		"happy path": {
			nums:     []int{1, 10, 22, 3, 7, 19, 5},
			target:   3,
			expected: 3,
		},
		"contains negative ints": {
			nums:     []int{1, 10, -3, -5, 0},
			target:   -3,
			expected: 2,
		},
		"only negative ints": {
			nums:     []int{-15, -10, -3, -5, -7},
			target:   -7,
			expected: 4,
		},
		"not found": {
			nums:     []int{-15, -10, -3, -5, -7},
			target:   -8,
			expected: -1,
		},
		"nil slice": {
			nums:        nil,
			target:      0,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			nums:        []int{},
			target:      0,
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IndexOf(test.nums, test.target)
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

func TestFilterEven(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    []int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: []int{10, 22},
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: []int{10, 0},
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: []int{-10},
		},
		"not found": {
			input:    []int{-15, -11, -3, -5, -7},
			expected: []int{},
		},
		"nil slice": {
			input:       nil,
			expected:    nil,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    nil,
			errExpected: fmt.Errorf("input is empty"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := FilterEven(test.input)
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

func TestCountVowels(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"happy path": {
			input:    "pericarezeracirep",
			expected: 8,
		},
		"happy path.2": {
			input:    "abeceda",
			expected: 4,
		},
		"unhappy path": {
			input:    "krt",
			expected: 0,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountVowels(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCountConsonants(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"happy path": {
			input:    "pericarezeracirep",
			expected: 9,
		},
		"happy path.2": {
			input:    "abeceda",
			expected: 3,
		},
		"unhappy path": {
			input:    "oooooo",
			expected: 0,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountConsonants(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected bool
	}{
		"happy path": {
			input:    2024,
			expected: true,
		},
		"unhappy path": {
			input:    2025,
			expected: false,
		},
		"misterious path": {
			input:    800,
			expected: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsLeapYear(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
