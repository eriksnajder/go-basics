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

func TestOddNotRepeatingNumbers(t *testing.T) {
	tests := map[string]struct {
		input1   []int
		input2   []int
		expected []int
	}{
		"happy path": {
			input1:   []int{1, 2, 3},
			input2:   []int{3, 4, 5},
			expected: []int{1, 5},
		},
		"one even, one odd path": {
			input1:   []int{2, 4, 6},
			input2:   []int{1, 3, 5},
			expected: []int{1, 3, 5},
		},
		"2 repeating slices path": {
			input1:   []int{1, 3},
			input2:   []int{1, 3},
			expected: []int{},
		},
		"A empty path": {
			input1:   []int{},
			input2:   []int{1, 2, 3},
			expected: []int{1, 3},
		},
		"B empty path": {
			input1:   []int{1, 2, 3},
			input2:   []int{},
			expected: []int{1, 3},
		},
		"empty path": {
			input1:   []int{},
			input2:   []int{},
			expected: []int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := OddNotRepeatingNumbers(test.input1, test.input2)

			assert.Equal(t, test.expected, actual)
		})

	}
}

func TestIndexSmallestMultipleOfFive(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{10, 20, 10, 30, 20, 40},
			expected: 0,
		},
		"no repeating values": {
			input:       []int{1, 2, 3, 4},
			errExpected: fmt.Errorf("no valid repeating multiple of 5"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no valid repeating multiple of 5"),
		},
		"repeating values": {
			input:    []int{5, 5, 10, 10},
			expected: 0,
		},
		"yet another succes": {
			input:    []int{25, 30, 25, 30, 15, 15},
			expected: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IndexSmallestMultipleOfFive(test.input)
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

func TestMostFrequentPerfectSquare(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{4, 4, 9, 9, 9, 16},
			expected: 9,
		},
		"no repeating values": {
			input:       []int{2, 3, 5},
			errExpected: fmt.Errorf("no perfect square values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no perfect square values"),
		},
		"same number of repeating values": {
			input:    []int{1, 1, 2, 2, 4, 4},
			expected: 1,
		},
		"yet another same number of repeating values": {
			input:    []int{25, 36, 25, 36},
			expected: 25,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MostFrequentPerfectSquare(test.input)
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

func TestDistinctPrimeDivisors(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"happy path": {
			input:    []int{6, 10, 15},
			expected: 3,
		},
		"no prime divisors repeated": {
			input:    []int{2, 3, 5, 7},
			expected: 0,
		},
		"happy path 2.0": {
			input:    []int{4, 15, 8, 10},
			expected: 2,
		},
		"happy path 3.0": {
			input:    []int{9, 15, 21},
			expected: 1,
		},
		"empty input": {
			input:    []int{},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := DistinctPrimeDivisors(test.input)

			assert.Equal(t, test.expected, actual)
		})

	}
}

func TestTwoMostFrequentNumbersDivByFour(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{4, 4, 8, 8, 12, 12},
			expected: 12,
		},
		"no repeating values": {
			input:       []int{2, 6, 10},
			errExpected: fmt.Errorf("not enough matching values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("not enough matching values"),
		},
		"enough matching values": {
			input:    []int{4, 4, 8},
			expected: 12,
		},
		"not enough matching values": {
			input:       []int{4, 4},
			errExpected: fmt.Errorf("not enough matching values"),
		},
		"yet another same number of repeating values": {
			input:    []int{16, 16, 20, 20, 24},
			expected: 36,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := TwoMostFrequentNumbersDivByFour(test.input)
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

func TestShortestStringNoRepeatingCharacters(t *testing.T) {
	tests := map[string]struct {
		input       []string
		expected    string
		errExpected error
	}{
		"happy path": {
			input:    []string{"123", "112", "789", "56"},
			expected: "56",
		},
		"letters/repeating in slice": {
			input:       []string{"abc", "999", "111"},
			errExpected: fmt.Errorf("no valid string found"),
		},
		"empty slice": {
			input:       []string{},
			errExpected: fmt.Errorf("no valid string found"),
		},
		"the weird one": {
			input:    []string{"9876", "1234", "88"},
			expected: "1234",
		},
		"single elem result": {
			input:    []string{"9", "98", "987"},
			expected: "9",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := ShortestStringNoRepeatingCharacters(test.input)
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

func TestGroupsOfEvenSumsOfConsecutiveOddNums(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"unhappy path": {
			input:    []int{1, 3, 5, 2, 7, 9, 11},
			expected: 0,
		},
		"happy path": {
			input:    []int{1, 3, 2, 5, 7, 2, 9, 11},
			expected: 3,
		},
		"no odd number path": {
			input:    []int{2, 4, 6},
			expected: 0,
		},
		"one group path": {
			input:    []int{1, 3, 5, 7},
			expected: 1,
		},
		"empty input": {
			input:    []int{},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := GroupsOfEvenSumsOfConsecutiveOddNums(test.input)

			assert.Equal(t, test.expected, actual)
		})

	}
}

func TestMostCommonWordLowered(t *testing.T) {
	tests := map[string]struct {
		input       []string
		expected    string
		errExpected error
	}{
		"succes hello": {
			input:    []string{"Hello!", "hello", "HELLO.", "Hi"},
			expected: "hello",
		},
		"succes yes": {
			input:    []string{"Yes?", "yes!", "YES", "no"},
			expected: "yes",
		},
		"empty slice": {
			input:       []string{},
			errExpected: fmt.Errorf("no input provided"),
		},
		"three diff words": {
			input:    []string{"One", "Two", "Two", "Three."},
			expected: "two",
		},
		"single repeated word": {
			input:    []string{"What's", "what's", "Whats"},
			expected: "whats",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MostCommonWordLowered(test.input)
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

func TestSumOfPrimeNumsNotRepeating(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{2, 3, 5, 5, 7, 10},
			expected: 12,
		},
		"no prime numbers": {
			input:       []int{4, 6, 8, 10},
			errExpected: fmt.Errorf("no qualifying values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no qualifying values"),
		},
		"only one correct number": {
			input:    []int{13, 15, 17, 13},
			expected: 17,
		},
		"not enough matching values": {
			input:    []int{1, 2, 3, 4},
			expected: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SumOfPrimeNumsNotRepeating(test.input)
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

func TestIsStringPalindrome(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []string
	}{
		"succes ": {
			input:    []string{"A man, a plan, a canal: Panama", "Racecar", "hello"},
			expected: []string{"A man, a plan, a canal: Panama", "Racecar"},
		},
		"unhappy path": {
			input:    []string{"abc", "def"},
			expected: []string{},
		},
		"single string complex succes": {
			input:    []string{"No 'x' in Nixon"},
			expected: []string{"No 'x' in Nixon"},
		},
		"all strings palindromes & same on newWord": {
			input:    []string{"Madam!", "madam", "MaDaM"},
			expected: []string{"Madam!", "madam", "MaDaM"},
		},
		"empty input": {
			input:    []string{},
			expected: []string{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsStringPalindrome(test.input)

			assert.Equal(t, test.expected, actual)
		})

	}
}

func TestAverageOfEvenNonnegativeNums(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes regular": {
			input:    []int{2, 4, -1, -3, 6, 7},
			expected: 4,
		},
		"no even & positive numbers": {
			input:       []int{1, 3, 5, -2},
			errExpected: fmt.Errorf("no even or/and positive values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no even or/and positive values"),
		},
		"only one correct number": {
			input:       []int{-10, -20, -30},
			errExpected: fmt.Errorf("no even or/and positive values"),
		},
		"not enough matching values": {
			input:    []int{10, 20, 30},
			expected: 20,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := AverageOfEvenNonnegativeNums(test.input)
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

func TestLongestStringFirstAndLastCharacterSame(t *testing.T) {
	tests := map[string]struct {
		input       []string
		expected    string
		errExpected error
	}{
		"succes racecar": {
			input:    []string{"Anna", "civic", "", "racecar", "apple"},
			expected: "racecar",
		},
		"empty strings failure": {
			input:       []string{"", "", ""},
			errExpected: fmt.Errorf("no valid word found"),
		},
		"empty slice": {
			input:       []string{},
			errExpected: fmt.Errorf("no valid word found"),
		},
		"succes two longest words": {
			input:    []string{"level", "stats", "bob"},
			expected: "stats",
		},
		"succes repeated word": {
			input:    []string{"wow", "deed", "deed"},
			expected: "deed",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := LongestStringFirstAndLastCharacterSame(test.input)
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

func TestNumsBetweenSmallestAndLargest(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes regular": {
			input:    []int{5, 12, 15, 20, 25},
			expected: 2,
		},
		"numbers < 10": {
			input:       []int{5, 8, 9},
			errExpected: fmt.Errorf("not enough values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("not enough values"),
		},
		"all same numbers": {
			input:       []int{11, 11, 11},
			errExpected: fmt.Errorf("not enough values"),
		},
		"smallest and largest neighbours 1": {
			input:    []int{10, 20, 30},
			expected: 0,
		},
		"smallest and largest neighbours 2": {
			input:    []int{50, 100, 70, 85},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := NumsBetweenSmallestAndLargest(test.input)
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
