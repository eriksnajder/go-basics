package naloge

import (
	"fmt"
	"math"
)

func NalogeIzFunctions() {
	// 1 Write a function `Factorial` that calculates the factorial of a non‑negative integer using a `for` loop and an `if` to handle zero.
	// n int

	n := 5

	result, err := Factorial(n)
	if err != nil {
		fmt.Printf("error calculating factorial: %v\n", err)
	} else {
		fmt.Println(result)
	}

	// 2 Write a function `IsPrime` that checks whether a given integer is prime. Loop from 2 up to `n-1`, use a conditional to return false if any divisor is found.
	// n int

	// 3 Write a function `MaxInSlice` that finds and returns the largest number in a slice of integers. Initialize a variable to the first element, then loop and update it.
	// nums []int

	// 4 Write a function `MinInSlice` that finds and returns the smallest number in a slice of integers. Initialize to the first element, loop and compare.
	// nums []int

	// 5 Write a function `SumSlice` that returns the sum of all elements in a slice of integers. Start a total at 0, loop and add each element.
	// nums []int

	// 6 Write a function `AverageSlice` that computes the average of numbers in a slice. Use the `SumSlice` logic, then use an `if` to avoid dividing by zero.
	// nums []int

	// 7 Write a function `CountEven` that counts how many even numbers are in a slice. Loop through each element and use `if num%2 == 0` to increment a counter.
	// nums []int

	// 8 Write a function `CountOdd` that counts how many odd numbers are in a slice. Loop and use `if num%2 != 0` to increment.
	// nums []int

	// 9 Write a function `ReverseString` that returns a new string which is the reverse of the input. Use a loop from the end index down to zero.
	// s string

	// 10 Write a function `IsPalindrome` that checks if a string reads the same forward and backward. Loop and compare characters at symmetrical positions.
	// s string

	// 11 Write a function `IndexOf` that finds the index of a target integer in a slice, returning -1 if not found. Loop with index and value and use an `if` check.
	// nums []int
	// target int

	// 12 Write a function `FilterEven` that returns a new slice containing only the even numbers from the input. Loop, test `num%2 == 0`, and append matches.
	// nums []int

	// 13 Write a function `CountVowels` that counts how many vowels (`a, e, i, o, u`) are in a string. Loop over runes and use a conditional with multiple `||` tests.
	// s string

	// 14 Write a function `CountConsonants` that counts consonants (letters that are not vowels) in a string. Loop and use `if` to skip non-letters and vowels.
	// s string

	// 15 Write a function `IsLeapYear` that determines whether a given year is a leap year. Use `if` to check divisibility by 4, 100, and 400 in the correct order.
	// year int

	// 16 Write a function `Fibonacci` that generates and returns the first `n` Fibonacci numbers in a slice. Use a `for` loop and handle `n` less than 2 with a conditional.
	// n int

	// 17 Write a function `SumDigits` that computes the sum of digits of a non‑negative integer. Use a loop dividing by 10 and adding `n % 10` each iteration.
	// n int

	// 18 Write a function `IsArmstrong` that checks whether a number is an Armstrong number (sum of its own digits each raised to the power of the number of digits). Loop to compute digit powers and compare.
	// n int

	// 19 Write a function `MultiplicationTable` that returns a slice of ints representing `n × 1` through `n × 10`. Use a loop from 1 to 10 and multiply.
	// n int

	// 20 Write a function `CountSubstring` that counts how many times `substr` appears in `s`. Loop stepping one rune at a time and use `if` with `strings.HasPrefix`.
	// s string
	// substr string

	// 21 Write a function `GCD` that computes the greatest common divisor of `a` and `b` using the Euclidean algorithm. Loop while `b` is not zero, swap and compute `a % b`.
	// a int
	// b int

	// 22 Write a function `LCM` that computes the least common multiple of `a` and `b`. Use the formula `abs(a*b)/GCD(a,b)` and a conditional to handle zero.
	// a int
	// b int

	// 23 Write a function `UniqueInts` that removes duplicate integers from a slice and returns a new slice. Loop over `nums`, use a map to track seen values, and append unseen.
	// nums []int

	// 24 Write a function `BubbleSort` that sorts a slice of integers using the bubble sort algorithm. Use nested loops and swap with a conditional `if nums[j] > nums[j+1]`.
	// nums []int

	// 25 Write a function `MergeSorted` that merges two sorted slices into one sorted slice. Use two indices, loop while both have elements, and compare to append.
	// slice1 []int
	// slice2 []int

	// 26 Write a function `SumMatrix` that sums all elements in a 2D slice (matrix). Use nested loops: outer for rows, inner for columns, add each value.
	// matrix [][]int

	// 27 Write a function `MainDiagonal` that returns the main diagonal elements of a square matrix as a slice. Loop once using index `i` and append `matrix[i][i]`.
	// matrix [][]int

	// 28 Write a function `IsSymmetric` that checks if a square matrix is symmetric (`matrix[i][j] == matrix[j][i]`). Use nested loops and return false on first mismatch.
	// matrix [][]int

	// 29 Write a function `PrimesInRange` that returns all prime numbers between `start` and `end` inclusive. Loop each number and call `IsPrime` to test, appending primes.
	// start int
	// end int

	// 30 Write a function `Power` that computes `base` raised to `exponent` for non‑negative integers. Use a loop multiplying an accumulator and handle exponent zero with a conditional.
	// base int
	// exponent int
}

// 1 Write a function `Factorial` that calculates the factorial of a non‑negative integer using a `for` loop and an `if` to handle zero.
// n int
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number: %d", n)

	}
	if n == 0 {
		return 0, nil
	}

	fact := n
	for i := n - 1; i > 0; i-- {
		fact *= i
	}
	return fact, nil
}

// 2 Write a function `IsPrime` that checks whether a given integer is prime. Loop from 2 up to `n-1`, use a conditional to return false if any divisor is found.
// n int //

func IsPrime(n int) (bool, error) {
	if n < 0 {
		return false, fmt.Errorf("negative number: %d", n)
	}

	if n == 0 {
		return false, nil
	}

	if n == 1 {
		return true, nil
	}

	for i := 2; i <= (n/2)+1; i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

// 3 Write a function `MaxInSlice` that finds and returns the largest number in a slice of integers. Initialize a variable to the first element, then loop and update it.
// nums []int

func MaxInSlice(nums []int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}

	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	number := nums[0]

	for _, i := range nums {
		if i > number {
			number = i
		}
	}
	return number, nil

}

// 4 Write a function `MinInSlice` that finds and returns the smallest number in a slice of integers. Initialize to the first element, loop and compare.
// nums []int

func MinInSlice(nums []int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}

	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	number := nums[0]

	for _, i := range nums {
		if i < number {
			number = i
		}
	}
	return number, nil
}

// 5 Write a function `SumSlice` that returns the sum of all elements in a slice of integers. Start a total at 0, loop and add each element.
// nums []int

func SumSlice(nums []int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}

	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	number := 0

	for _, i := range nums {
		number += i
	}
	return number, nil
}

// 6 Write a function `AverageSlice` that computes the average of numbers in a slice. Use the `SumSlice` logic, then use an `if` to avoid dividing by zero.
// nums []int

func AverageSlice(nums []int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}

	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	number := 0

	for _, i := range nums {
		number += i
	}

	return number / len(nums), nil
}

// 7 Write a function `CountEven` that counts how many even numbers are in a slice. Loop through each element and use `if num%2 == 0` to increment a counter.
// nums []int

func CountEven(nums []int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}

	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	number := 0

	for _, i := range nums {
		if i%2 == 0 {
			number += 1
		}
	}
	return number, nil
}

// 8 Write a function `CountOdd` that counts how many odd numbers are in a slice. Loop and use `if num%2 != 0` to increment.
// nums []int
func CountOdd(nums []int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}
	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	number := 0

	for _, i := range nums {
		if i%2 != 0 {
			number += 1
		}
	}
	return number, nil
}

// 9 Write a function `ReverseString` that returns a new string which is the reverse of the input. Use a loop from the end index down to zero.
// s string
func ReverseString(word string) string {
	if word == "" {
		return ""
	}
	newWord := ""

	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}

	return newWord
}

// 10 Write a function `IsPalindrome` that checks if a string reads the same forward and backward. Loop and compare characters at symmetrical positions.
// s string
func IsPalindrome(word string) bool {
	if word == "" {
		return false
	}
	newWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}
	return newWord == word
}

// 11 Write a function `IndexOf` that finds the index of a target integer in a slice, returning -1 if not found. Loop with index and value and use an `if` check.
// nums []int
// target int

func IndexOf(nums []int, target int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("input is nil slice")
	}
	if len(nums) == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	for index, item := range nums {
		if target == item {
			return index, nil
		}
	}
	return -1, nil
}

// 12 Write a function `FilterEven` that returns a new slice containing only the even numbers from the input. Loop, test `num%2 == 0`, and append matches.
// nums []int

func FilterEven(nums []int) ([]int, error) {
	if nums == nil {
		return nil, fmt.Errorf("input is nil slice")
	}
	if len(nums) == 0 {
		return nil, fmt.Errorf("input is empty")
	}
	newNums := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			newNums = append(newNums, num)
		}

	}
	return newNums, nil
}

// 13 Write a function `CountVowels` that counts how many vowels (`a, e, i, o, u`) are in a string. Loop over runes and use a conditional with multiple `||` tests.
// s string

func IsVowel(s rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		if s == vowel {
			return true
		}
	}
	return false
}

func CountVowels(s string) int {
	count := 0
	for _, letter := range s {
		if IsVowel(letter) {
			count += 1
		}
	}
	return count

}

// 14 Write a function `CountConsonants` that counts consonants (letters that are not vowels) in a string. Loop and use `if` to skip non-letters and vowels.
// s string

func CountConsonants(s string) int {
	count := 0
	for _, letter := range s {
		if !IsVowel(letter) {
			count += 1
		}
	}
	return count
}

// 15 Write a function `IsLeapYear` that determines whether a given year is a leap year. Use `if` to check divisibility by 4, 100, and 400 in the correct order.
// year int

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 {
			return true
		} else if year%400 == 0 {
			return true
		}
	}
	return false
}

// 16 Write a function `Fibonacci` that generates and returns the first `n` Fibonacci numbers in a slice. Use a `for` loop and handle `n` less than 2 with a conditional.
// n int

func Fibonacci(n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("invalid n")
	}

	if n == 1 {
		return []int{0}, nil
	}

	out := []int{0, 1}

	x := 0
	y := 1
	z := 0

	for i := 3; i <= n; i++ {
		z = x + y
		x = y
		y = z
		out = append(out, z)
	}
	return out, nil
}

// 17 Write a function `SumDigits` that computes the sum of digits of a non‑negative integer. Use a loop dividing by 10 and adding `n % 10` each iteration.
// n int

func SumDigits(num int) (int, error) {
	if num < 0 {
		return -1, fmt.Errorf("non-negative numbers only")
	}

	if num < 10 {
		return num, nil
	}

	sum := 0
	for {
		currentDigit := num % 10
		num /= 10
		sum += currentDigit
		if num == 0 {
			break
		}

	}
	return sum, nil
}

// 18 Write a function `IsArmstrong` that checks whether a number is an Armstrong number (sum of its own digits each raised to the power of the number of digits). Loop to compute digit powers and compare.
// n int

func IsArmstrong(num int) (bool, error) {
	if num < 0 {
		return false, fmt.Errorf("non-negative numbers only")
	}

	if num < 10 {
		return true, nil
	}

	stevke := []int{}
	original := num

	for num != 0 {
		stevke = append(stevke, num%10)
		num /= 10
	}

	sum := 0

	for _, stevka := range stevke {
		sum += int(math.Pow(float64(stevka), float64(len(stevke))))
	}

	return sum == original, nil

}

// 19 Write a function `MultiplicationTable` that returns a slice of ints representing `n × 1` through `n × 10`. Use a loop from 1 to 10 and multiply.
// n int

func MultiplicationTable(number int, x int) ([]int, error) {
	intSlices := []int{}
	if x < 0 {
		return nil, fmt.Errorf("invalid x")
	}
	for i := 1; i <= x; i++ {
		intSlices = append(intSlices, number*i)
	}
	return intSlices, nil

}
