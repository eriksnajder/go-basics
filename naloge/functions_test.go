package naloge

import (
	"errors"
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

func TestFibonacci(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    []int
		errExpected error
	}{
		"invalid path": {
			input:       0,
			expected:    nil,
			errExpected: fmt.Errorf("invalid n"),
		},
		"zero path": {
			input:    1,
			expected: []int{0},
		},
		"happy path": {
			input:    5,
			expected: []int{0, 1, 1, 2, 3},
		},
		"happy path 2": {
			input:    10,
			expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Fibonacci(test.input)
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

func TestSumDigits(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    int
		errExpected error
	}{
		"invalid path": {
			input:       -3,
			expected:    -1,
			errExpected: fmt.Errorf("non-negative numbers only"),
		},
		"zero path": {
			input:    0,
			expected: 0,
		},
		"single path": {
			input:    5,
			expected: 5,
		},
		"happy path": {
			input:    51,
			expected: 6,
		},
		"happy path 2": {
			input:    8598,
			expected: 30,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SumDigits(test.input)
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

func TestIsArmstrong(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    bool
		errExpected error
	}{
		"invalid path": {
			input:       -3,
			expected:    false,
			errExpected: fmt.Errorf("non-negative numbers only"),
		},
		"zero path": {
			input:    0,
			expected: true,
		},
		"single path": {
			input:    5,
			expected: true,
		},
		"unhappy path": {
			input:    51,
			expected: false,
		},
		"happy path 2": {
			input:    1634,
			expected: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IsArmstrong(test.input)
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

func TestMultiplicationTable(t *testing.T) {
	tests := map[string]struct {
		inputNumber int
		inputX      int
		expected    []int
		errExpected error
	}{
		"happy path": {
			inputNumber: 4,
			inputX:      4,
			expected:    []int{4, 8, 12, 16},
		},
		"happy path.2": {
			inputNumber: 15,
			inputX:      8,
			expected:    []int{15, 30, 45, 60, 75, 90, 105, 120},
		},
		"negative path": {
			inputNumber: -3,
			inputX:      5,
			expected:    []int{-3, -6, -9, -12, -15},
		},
		"invalid x": {
			inputNumber: 5,
			inputX:      -3,
			expected:    nil,
			errExpected: fmt.Errorf("invalid x-"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MultiplicationTable(test.inputNumber, test.inputX)
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

func TestCountSubstring(t *testing.T) {
	tests := map[string]struct {
		inputS         string
		inputSubstring string
		expected       int
	}{
		"happy path": {
			inputS:         "foobarbaz",
			inputSubstring: "ba",
			expected:       2,
		},
		"happy path.2": {
			inputS:         "foobarbaz",
			inputSubstring: "bar",
			expected:       1,
		},
		"zero path": {
			inputS:         "foobarbaz",
			inputSubstring: "bay",
			expected:       0,
		},
		"nil path": {
			inputS:         "",
			inputSubstring: "ba",
			expected:       0,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountSubstring(test.inputS, test.inputSubstring)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestGCD(t *testing.T) {
	tests := map[string]struct {
		a           int
		b           int
		expected    int
		errExpected error
	}{
		"21": {
			a:        252,
			b:        105,
			expected: 21,
		},
		"common divisor": {
			a:        48,
			b:        18,
			expected: 6,
		},
		"one divides the other": {
			a:        100,
			b:        25,
			expected: 25,
		},
		"co-primes": {
			a:        7,
			b:        3,
			expected: 1,
		},
		"zero b – should error": {
			a:           42,
			b:           0,
			errExpected: errors.New("b cannot be zero"),
		},
		"both zero – should error": {
			a:           0,
			b:           0,
			errExpected: fmt.Errorf("b cannot be zero"),
		},
		"negative input": {
			a:        -8,
			b:        12,
			expected: 4,
		},
		"larger numbers": {
			a:        270,
			b:        192,
			expected: 6,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := GCD(test.a, test.b)
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
func TestLCM(t *testing.T) {
	tests := map[string]struct {
		a           int
		b           int
		expected    int
		errExpected error
	}{
		"basic LCM": {
			a:        4,
			b:        6,
			expected: 12,
		},
		"same number": {
			a:        7,
			b:        7,
			expected: 7,
		},
		"one is multiple of other": {
			a:        3,
			b:        9,
			expected: 9,
		},
		"co-primes": {
			a:        5,
			b:        7,
			expected: 35,
		},
		"zero input – should error": {
			a:           0,
			b:           5,
			errExpected: fmt.Errorf("b cannot be zero"),
		},
		"both zero – should error": {
			a:           0,
			b:           0,
			errExpected: fmt.Errorf("b cannot be zero"),
		},
		"negative input": {
			a:        -4,
			b:        6,
			expected: 12,
		},
		"large numbers": {
			a:        21,
			b:        6,
			expected: 42,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := LCM(test.a, test.b)
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

func TestUniqueInts(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"no duplicates": {
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		"some duplicates": {
			input:    []int{1, 2, 2, 3, 1, 4},
			expected: []int{1, 2, 3, 4},
		},
		"all duplicates": {
			input:    []int{5, 5, 5, 5},
			expected: []int{5},
		},
		"empty list": {
			input:    []int{},
			expected: []int{},
		},
		"already unique but out of order": {
			input:    []int{9, 7, 3, 1},
			expected: []int{9, 7, 3, 1},
		},
		"negative and positive mix": {
			input:    []int{-1, 2, -1, 2, 0},
			expected: []int{-1, 2, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := UniqueInts(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"already sorted": {
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		"reverse sorted": {
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		"random order with duplicates": {
			input:    []int{3, 1, 4, 5, 9, 2},
			expected: []int{1, 2, 3, 4, 5, 9},
		},
		"empty slice": {
			input:    []int{},
			expected: []int{},
		},
		"single element": {
			input:    []int{7},
			expected: []int{7},
		},
		"all same values": {
			input:    []int{2, 2, 2, 2},
			expected: []int{2, 2, 2, 2},
		},
		"negatives and positives": {
			input:    []int{-1, 3, -2, 0, 5},
			expected: []int{-2, -1, 0, 3, 5},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := BubbleSort(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMergeSorted(t *testing.T) {
	tests := map[string]struct {
		input1   []int
		input2   []int
		expected []int
	}{
		"success": {
			input1:   []int{1, 2, 4, 6},
			input2:   []int{3, 4, 5, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		"success duplicates": {
			input1:   []int{1, 2, 4, 6},
			input2:   []int{3, 4, 5, 7, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := MergeSorted(test.input1, test.input2)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSumMatrix(t *testing.T) {
	tests := map[string]struct {
		input    [][]int
		expected int
	}{
		"success": {
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: 45,
		},
		"success 2": {
			input:    [][]int{{1, 2, 6}, {4, 5, 5}, {7, 8, 19}},
			expected: 57,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := SumMatrix(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMainDiagonal(t *testing.T) {
	tests := map[string]struct {
		input    [][]int
		expected []int
	}{
		"success": {
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 5, 9},
		},
		"success 2": {
			input:    [][]int{{1, 2, 6}, {4, 5, 5}, {7, 8, 19}},
			expected: []int{1, 5, 19},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := MainDiagonal(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
