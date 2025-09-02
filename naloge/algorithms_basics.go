package naloge

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 01 Return the smallest even number that appears more than once
// Test cases:
//   []int{2, 4, 2, 5, 6, 6} => 2
//   []int{3, 5, 7, 9} => "no repeating even values"
//   []int{} => "no repeating even values"
//   []int{10, 10, 8, 8, 6} => 8
//   []int{1, 2, 3, 2, 4} => 2

func SmallestRepeatingEven(numbers []int) (int, error) {
	frequencies := map[int]int{}
	for _, number := range numbers {
		if number%2 != 0 {
			continue
		}
		_, ok := frequencies[number] // ok = true, če number obstaja v mapu
		if !ok {                     // če ni v mapu, dodaj key & value
			frequencies[number] = 1
		} else { // če je v mapu, prištej eno vrednosti key-a
			frequencies[number]++
		}
	}

	smallest := 0
	overriden := false
	for number, frequency := range frequencies {
		if !overriden && frequency > 1 {
			smallest = number
			overriden = true
			continue
		}

		if frequency > 1 && number < smallest {
			smallest = number
		}
	}

	if !overriden {
		return 0, fmt.Errorf("no repeating even values")
	}

	return smallest, nil
}

// 02 Return all numbers that are both prime and appear at least twice
// Test cases:
//   []int{2, 3, 3, 5, 5, 6} => [3 5]
//   []int{4, 6, 8, 9} => []
//   []int{2, 2, 2, 3, 3, 4, 4} => [2 3]
//   []int{1, 1, 1} => []
//   []int{} => []

func isPrime(number int) bool {
	if number == 2 {
		return true
	}

	if number == 1 {
		return false
	}

	for i := 2; i <= (number/2)+1; i++ {
		if number%i == 0 {
			return false
		}

	}
	return true
}

func PrimeAndRepeating(numbers []int) []int {
	frequencies := map[int]int{}
	out := []int{}

	for _, number := range numbers {

		_, ok := frequencies[number]
		if !ok {
			frequencies[number] = 1
		} else {
			frequencies[number]++
		}
	}
	for number, frequency := range frequencies {
		if isPrime(number) && frequency > 1 {
			out = append(out, number)
		}
	}

	return out
}

// 03 Return the average of all values that are at odd indices and divisible by 3
// Test cases:
//   []int{3, 6, 9, 12, 15, 18} => 12.0
//   []int{1, 2, 3, 4} => "no matching values at odd indices"
//   []int{} => "no matching values at odd indices"
//   []int{3, 3, 3, 3} => 3.0
//   []int{5, 6, 7, 9, 11, 12} => 9.0

// 1. naredi seznam, ki bo vseboval numbers[n], ki bodo na lihem indeksu.
// 2. naredi seznam, ki bo vseboval oddNumbers[n], ki so deljiva s 3.
// 3. naredi loop, ki bo izračunal vsoto divThreeNumbers[n], in na koncu izračunal average[divThreeNumbers]

func AverageOddIndiceDivByThree(numbers []int) (float64, error) {
	oddNumbers := []int{}
	divThreeNumbers := []int{}
	var sum float64

	if len(numbers) == 0 {
		return 0, fmt.Errorf("no matching values at odd indices")
	}

	for i, number := range numbers {
		if i%2 != 0 {
			oddNumbers = append(oddNumbers, number)
		}
	}

	for _, number := range oddNumbers {
		if number%3 == 0 {
			divThreeNumbers = append(divThreeNumbers, number)
		}
	}

	if len(divThreeNumbers) == 0 {
		return 0, fmt.Errorf("no matching values at odd indices")
	}

	for _, number := range divThreeNumbers {
		sum += float64(number)
	}

	return sum / float64(len(divThreeNumbers)), nil
}

// 04 Return the number of unique values that are palindromes and odd
// Test cases:
//   []int{1, 3, 5, 7, 9, 11, 33, 44} => 7
//   []int{2, 4, 6, 8} => 0
//   []int{121, 131, 141, 151} => 4
//   []int{} => 0
//   []int{1, 1, 1, 1} => 1

// 1. naredi funkcijo, ki prepozna števila, ki so palindromi

func numberPalindrome(number int) bool {
	strVal := fmt.Sprintf("%d", number)
	if strVal == "" {
		return false
	}
	newStrVal := ""
	for i := len(strVal) - 1; i >= 0; i-- {
		newStrVal += string(strVal[i])
	}
	return newStrVal == strVal
}

// 2. naredi seznam, ki bo vseboval numbers[n], ki so palindromi
// 3. naredi seznam, ki bo vseboval palindromes[n], ki so liha števila
// 4. vrni len(oddPalindromes)

func OddPalindromes(numbers []int) int {
	palindromes := []int{}
	oddPalindromes := map[int]bool{}

	for _, number := range numbers {
		if numberPalindrome(number) {
			palindromes = append(palindromes, number)
		}
	}

	for _, number := range palindromes {
		if number%2 != 0 {
			oddPalindromes[number] = true

		}
	}

	return len(oddPalindromes)
}

// 05 Return the second largest value that appears more than once
// Test cases:
//   []int{5, 5, 6, 6, 7, 7} => 6
//   []int{1, 1, 2, 3, 4, 5} => "not enough repeating values"
//   []int{9, 9, 8, 8, 7, 7, 6, 6} => 8
//   []int{} => "not enough repeating values"
//   []int{2, 2, 2, 1, 1} => 1

// 1. napiši funkcijo, ki uporabila slovar 'key : value' in bo prepoznala, kateri keyi imajo value > 1.
// 2.

func SecondLargestRepeatingValue(numbers []int) (int, error) {
	frequencies := map[int]int{}

	for _, number := range numbers {
		frequencies[number]++
	}

	repeated := []int{}
	for number, frequency := range frequencies {
		if frequency > 1 {
			repeated = append(repeated, number)
		}
	}

	if len(repeated) < 2 {
		return 0, fmt.Errorf("not enough repeating values")
	}

	var highest int
	var secondHighest int

	if repeated[0] > repeated[1] {
		highest = repeated[0]
		secondHighest = repeated[1]
	} else {
		secondHighest = repeated[0]
		highest = repeated[1]
	}

	for _, number := range repeated[2:] {
		if number > highest {
			secondHighest = highest
			highest = number
		} else if number > secondHighest {
			secondHighest = number
		}
	}

	return secondHighest, nil

}

// 06 Return the longest increasing sub-slice from a list of integers
// Test cases:
//   []int{1, 2, 3, 2, 3, 4, 5} => [2 3 4 5]
//   []int{5, 4, 3, 2, 1} => [5]
//   []int{} => "empty input"
//   []int{1, 2, 1, 2, 3} => [1 2 3]
//   []int{10} => [10]

// 1. iz slicea "s" bomo ven jemali sub-slice "subS"
// 2. naredimo loop, ki bo preverjal, ali številke naraščajo in si zapomnil najdaljši "len(subs)"
// 3. loop bo vrnil "longestSubS"

func LongestIncreasingSubSlice(s []int) ([]int, error) {

	if len(s) == 0 {
		return []int{}, fmt.Errorf("empty input")
	}

	subS := []int{s[0]}
	longestSubS := []int{}
	growingInt := s[0]

	for _, number := range s[1:] {
		if number > growingInt {
			subS = append(subS, number)
			growingInt = number
		} else {
			if len(subS) > len(longestSubS) {
				longestSubS = subS
				subS = []int{number}
				growingInt = number
			} else {
				subS = []int{number}
				growingInt = number
			}
		}
	}
	if len(subS) > len(longestSubS) {
		longestSubS = subS
	}

	return longestSubS, nil

}

// 07 Return the product of all odd numbers that are also perfect squares
// Test cases:
//   []int{1, 4, 9, 16, 25} => 225
//   []int{2, 3, 5, 7} => "no matching values"
//   []int{} => "no matching values"
//   []int{1, 1, 1} => 1
//   []int{49, 36, 25} => 1225

// 1. preveri, ali so v sliceu "s" liha števila, ki so hkrati kvadratna potenca nekega naravnea števila, in jih dodaj v nov slice.
// 2. vsa števila v novem seznamu pomnoži in na koncu vrni "product"

func PerfectSquareOddNumberProduct(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no matching values")
	}

	oddNumbers := []int{}
	product := 1

	for _, number := range numbers {

		if number%2 != 0 {
			if number == 1 {
				oddNumbers = append(oddNumbers, number)
				continue
			}

			sqrt := int(math.Sqrt(float64(number)))
			if sqrt*sqrt == number {
				oddNumbers = append(oddNumbers, number)
			}
		}
	}

	if len(oddNumbers) == 0 {
		return 0, fmt.Errorf("no matching values")
	}

	for _, number := range oddNumbers {
		product *= number
	}

	return product, nil
}

// 08 Return the first value that is both a multiple of 3 and appears exactly twice
// Test cases:
//   []int{3, 3, 6, 6, 9, 9} => 3
//   []int{2, 4, 6, 8, 10} => "no valid match found"
//   []int{} => "no valid match found"
//   []int{6, 6, 3, 3} => 6
//   []int{12, 12, 15} => 12

// 1. v sliceu "numbers" preveri številke, če so deljive s 3.
// 2. v kolikor je, jo dodaj v nov map, kjer boš za vsak key iskal value.
// 3. čim je eden izmed value > 1, --> return key in zaključi.

func FirstMultOfThreeExactlyTwice(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no valid match found")
	}

	frequencies := map[int]int{}

	for _, number := range numbers {
		frequencies[number]++
	}

	for _, number := range numbers {
		if number%3 != 0 {
			continue
		}
		value := frequencies[number]
		if value == 2 {
			return number, nil
		}
	}

	return 0, fmt.Errorf("no valid match found")
}

// 09 From two slices, return all values that are odd and only appear in one of the slices
// Test cases:
//   A: [1, 2, 3], B: [3, 4, 5] => [1 5]
//   A: [2, 4, 6], B: [1, 3, 5] => [1 3 5]
//   A: [1, 3], B: [1, 3] => []
//   A: [], B: [1, 2, 3] => [1 3]
//   A: [7, 9], B: [] => [7 9]

// 1. 2 prazna slicea "numbersA" in "numbersB", v katere se bo dajalo liha števila iz "numbers1" in "numbers2"
// 2. Preveri "numberA" in "numbersB", če se katerokoli število pojavi v obeh sliceih
// 3. Vrni slice lihih števil "numbersOK", ki se ne ponavljajo.

func OddNotRepeatingNumbers(numbersA, numbersB []int) []int {
	numbersOK := []int{}

	if len(numbersA) == 0 && len(numbersB) == 0 {
		return numbersOK
	}

	for _, number := range numbersA {
		if number%2 == 0 {
			continue
		}

		found := false

		for _, num := range numbersB {
			if number == num {
				found = true
			}
		}

		if !found {
			numbersOK = append(numbersOK, number)
		}

	}

	for _, number := range numbersB {
		if number%2 == 0 {
			continue
		}

		found := false

		for _, num := range numbersA {
			if number == num {
				found = true
			}
		}

		if !found {
			numbersOK = append(numbersOK, number)
		}

	}

	return numbersOK

}

// 10 Return the index of the smallest number that appears more than once and is a multiple of 5
// Test cases:
//   []int{10, 20, 10, 30, 20, 40} => 0
//   []int{1, 2, 3, 4} => "no valid repeating multiple of 5"
//   []int{} => "no valid repeating multiple of 5"
//   []int{5, 5, 10, 10} => 0
//   []int{25, 30, 25, 30, 15, 15} => 4

func IndexSmallestMultipleOfFive(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no valid repeating multiple of 5")
	}

	seen := map[int]int{}

	var smallest *int
	var smallestIdx *int
	for idx, number := range numbers {
		if number%5 == 0 {
			i, ok := seen[number]
			if !ok {
				seen[number] = idx
				continue
			}
			if smallest == nil && ok {
				smallest = &number
				smallestIdx = &i
				continue
			}

			if number < *smallest && ok {
				smallest = &number
				smallestIdx = &i
			}

			seen[number] = idx
		}
	}

	if smallestIdx == nil {
		return 0, fmt.Errorf("no valid repeating multiple of 5")
	}

	return *smallestIdx, nil
}

// 11 Return the value that occurs most frequently, but only if it is a perfect square
// Test cases:
//   []int{4, 4, 9, 9, 9, 16} => 9
//   []int{2, 3, 5} => "no perfect square values"
//   []int{} => "no perfect square values"
//   []int{1, 1, 2, 2, 4, 4} => 1
//   []int{25, 36, 25, 36} => 25

// 1. naredi map "frequencies", ki bo vrnil 'key : value' števil iz slicea "numbers", ki so perfect square.
// 2. poišči in vrni iz funkcije število, ki ima najvišji value.

func MostFrequentPerfectSquare(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no perfect square values")
	}

	frequencies := map[int]int{}

	for _, number := range numbers {
		sqrt := int(math.Sqrt(float64(number)))
		if sqrt*sqrt == number {
			frequencies[number]++
		}

	}

	if len(frequencies) == 0 {
		return 0, fmt.Errorf("no perfect square values")
	}

	highestKey := 0
	highestValue := 0

	for _, number := range numbers {
		frequency, ok := frequencies[number]
		if !ok {
			continue
		}
		if frequency > highestValue {
			highestKey = number
			highestValue = frequency
		}
	}

	return highestKey, nil
}

// 12 Return the number of distinct primes that divide at least two different numbers in the list
// Test cases:
//   []int{6, 10, 15} => 3       // 2 and 5 divide multiple
//   []int{2, 3, 5, 7} => 0
//   []int{4, 15, 8, 10} => 2     // 2
//   []int{} => 0
//   []int{9, 15, 21} => 1       // 3

func primes(num int) []int {
	out := []int{}
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			out = append(out, i)
		}
	}
	return out
}

func DistinctPrimeDivisors(numbers []int) int {
	seen := map[int]int{}

	for _, number := range numbers {
		allPrimes := primes(number)

		for _, prime := range allPrimes {
			if number%prime == 0 {
				seen[prime]++
			}
		}
	}

	total := 0
	for _, frequency := range seen {
		if frequency >= 2 {
			total++
		}
	}

	return total
}

// 13 Return the sum of the first two most frequent even numbers that are also divisible by 4
// Test cases:
//   []int{4, 4, 8, 8, 12, 12} => 12
//   []int{2, 6, 10} => "not enough matching values"
//   []int{4, 4, 8} => 12
//   []int{16, 16, 20, 20, 24} => 36
//   []int{} => "not enough matching values"

// 1. preglej čez slice "numbers" in naredi map "frequencies", ki bo vseboval 'key : value' vseh sodih števil, ki so deljiva s 4
// 2. poišči prvi dve števili, ki se pojavita največkrat, in ju vrni iz funkcije kot vsoto "sum".

func TwoMostFrequentNumbersDivByFour(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("not enough matching values")
	}

	frequencies := map[int]int{}

	for _, number := range numbers {
		if number%4 == 0 {
			frequencies[number]++
		}
	}

	if len(frequencies) <= 1 {
		return 0, fmt.Errorf("not enough matching values")
	}

	highestKey := 0
	highestValue := 0
	secondHighestKey := 0
	secondHighestValue := 0

	for _, number := range numbers {
		frequency, ok := frequencies[number]
		if !ok || number == highestKey || number == secondHighestKey {
			continue
		}
		if frequency > highestValue {
			secondHighestKey = highestKey
			secondHighestValue = highestValue
			highestKey = number
			highestValue = frequency
		} else if frequency > secondHighestValue {
			secondHighestKey = number
			secondHighestValue = frequency
		}

	}
	sum := highestKey + secondHighestKey

	return sum, nil
}

// 14 Given a list of strings, return the shortest one that contains only digits and has no repeated characters
// Test cases:
//   []string{"123", "112", "789", "56"} => "56"
//   []string{"abc", "999", "111"} => "no valid string found"
//   []string{} => "no valid string found"
//   []string{"9876", "1234", "88"} => "1234" ?????????
//   []string{"9", "98", "987"} => "9"

// 1. naredi map "frequencies", ki bo vseboval stringe iz sliceov "numbers" in bo podal 'key(string):value(dolzina stringa)'
// 2. preveri ali so elementi v stringu števila
// 3. preveri, kateri je najkrajši in nima ponavljajočih elementov.
// 4. funkcija naj vrne string(key) "number" ali 'error'

func ShortestStringNoRepeatingCharacters(numbers []string) (string, error) {
	errInvalid := fmt.Errorf("no valid string found")
	if len(numbers) == 0 {
		return "", errInvalid
	}

	shortest := ""

	for _, s := range numbers {
		isNumeric := true
		repeating := false
		frequencies := map[int]int{}

		for _, char := range s {
			currentNum, err := strconv.Atoi(string(char))
			if err != nil {
				isNumeric = false
				break
			}

			frequencies[currentNum]++
			if frequencies[currentNum] > 1 {
				repeating = true
				break
			}
		}

		if !isNumeric || repeating {
			continue
		}
		if len(shortest) == 0 || len(s) <= len(shortest) {
			shortest = s
		}
	}

	if len(shortest) > 0 {
		return shortest, nil
	}

	return "", errInvalid
}

// 15 From a list of numbers, return the average of all values that:
// Test cases:
//   []int{3, 33, 6, 121, 303, 9} => 105.0     // 303 only one that passes
//   []int{3, 3, 6, 6, 9, 9} => "no matching values"
//   []int{6, 33, 303, 12321} => 169.5         // 33 and 303
//   []int{} => "no matching values"
//   []int{111, 222, 303, 303} => 111.0

// 16 Count how many groups of consecutive odd numbers sum to an even number
// Test cases:
//   []int{1, 3, 5, 2, 7, 9, 11} => 0      // [1 3 5] = 9 (odd), [7 9 11] = 27 (odd), skip; no even
//   []int{1, 3, 2, 5, 7, 2, 9, 11} => 3   // [5 7] = 12 (even)
//   []int{} => 0
//   []int{2, 4, 6} => 0
//   []int{1, 3, 5, 7} => 1

// 1. preverjaj skozi slice 'numbers' in seštevaj vsoto števil, dokler ne naletiš na sodo število.
// 2. ko prideš do sodega, preveri če je vsota liho ali sodo število.
// 3. če je sodo, prištej v 'consGroups += 1', če ne, continue.
// 4. preden začneš preverjati naslednja števila v 'numbers', spravi vsoto na 0 in ponovi.

func GroupsOfEvenSumsOfConsecutiveOddNums(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0
	consGroups := 0

	for _, number := range numbers {
		if number%2 != 0 {
			sum += number
		} else {
			if sum == 0 {
				continue
			} else if sum%2 == 0 {
				consGroups += 1
				sum = 0
			} else {
				sum = 0
			}
		}

	}
	if sum%2 == 0 && sum != 0 {
		consGroups += 1
		sum = 0
	} else {
		sum = 0
	}

	return consGroups
}

// 17 From a slice of strings, return the most common word after lowercasing and removing punctuation
// Test cases:
//   ["Hello!", "hello", "HELLO.", "Hi"] => "hello"
//   [] => "no input provided"
//   ["Yes?", "yes!", "YES", "no"] => "yes"
//   ["One", "Two", "Two", "Three."] => "two"
//   ["What's", "what's", "Whats"] => "whats"

func MostCommonWordLowered(words []string) (string, error) {
	specialChars := "!.,?;:=<>'*"
	errInvalid := fmt.Errorf("no input provided")
	if len(words) == 0 {
		return "", errInvalid
	}

	z := 0
	mostCommonWord := ""
	frequencies := map[string]int{}
	for _, word := range words {
		s := strings.ToLower(word)
		for _, special := range specialChars {
			s = strings.Replace(s, string(special), "", -1)
		}

		frequencies[s]++
	}
	for word, frequency := range frequencies {
		if frequency > z {
			mostCommonWord = word
			z = frequency
		}

	}

	return mostCommonWord, nil
}

// 18. Given a slice of strings, check if these two results are the same:
//
// 1) Reverse every word, then sort the list.
// 2) Sort the list first, then reverse every word.
//
// Return true if both results are identical, otherwise false.
//
// Test cases:
//   ["abc", "def"] => true
//   ["abc", "cba"] => false
//   []              => true
//   ["a", "b", "c"] => true
//   ["xy", "yx"]    => false

// 19. Given a list of integers, return the sum of all values that are prime
// and occur exactly once in the list.
//
// If no such values exist, return "no qualifying values".
//
// Test cases:
//   []int{2, 3, 5, 5, 7, 10} => 12      // 2 + 3 + 7
//   []int{4, 6, 8, 10}       => "no qualifying values"
//   []int{}                  => "no qualifying values"
//   []int{13, 15, 17, 13}    => 17      // 17 only
//   []int{1, 2, 3, 4}

func SumOfPrimeNumsNotRepeating(numbers []int) (int, error) {
	errInvalid := fmt.Errorf("no qualifying values")
	if len(numbers) == 0 {
		return 0, errInvalid
	}

	frequencies := map[int]int{}
	sum := 0

	for _, number := range numbers {
		if isPrime(number) {
			frequencies[number]++
		}
	}

	for number, frequency := range frequencies {
		if frequency == 1 {
			sum += number
		}
	}

	return sum, nil
}

// 20 From a list of strings, return all that are palindromes after removing non-letter characters and lowercasing
// Test cases:
//   ["A man, a plan, a canal: Panama", "Racecar", "hello"] => ["A man, a plan, a canal: Panama", "Racecar"]
//   ["abc", "def"] => []
//   [] => []
//   ["No 'x' in Nixon"] => ["No 'x' in Nixon"]
//   ["Madam!", "madam", "MaDaM"] => ["Madam!", "madam", "MaDaM"]

// 21 Compute the average of all even non-negative numbers, rounded down.
// Test cases:
//   []int{2, 4, -1, -3, 6, 7} => 4
//   []int{1, 3, 5, -2} => "no even values"
//   []int{} => "no even values"
//   []int{-10, -20, -30} => "no even values"
//   []int{10, 20, 30} => 20

// 22 Return the longest string that starts and ends with the same letter (case-insensitive).
// Test cases:
//   []string{"Anna", "civic", "", "racecar", "apple"} => "racecar"
//   []string{"", "", ""} => "no valid word found"
//   []string{} => "no valid word found"
//   []string{"level", "stats", "bob"} => "stats"
//   []string{"wow", "deed", "deed"} => "deed"

// 23 Filter out all numbers less than or equal to 10.
// Find the smallest and largest of the remaining numbers.
// Count how many of the original numbers are strictly between those two.
// Test cases:
//   []int{5, 12, 15, 20, 25} => 2
//   []int{5, 8, 9} => "not enough values"
//   []int{11, 11, 11} => "not enough values"
//   []int{10, 20, 30} => 1
//   []int{50, 100, 70, 85} => 1

// 24 Return a list of strings that appear only once sorted by length.
// Test cases:
//   []string{"a", "bb", "ccc", "a", "bb", "dddd"} => ["ccc" "dddd"]
//   []string{"same", "same", "same"} => []
//   []string{} => []
//   []string{"one", "two", "three"} => ["one" "two" "three"]
//   []string{"a", "ab", "abc", "ab", "abc"} => ["a"]

// 25 Remove all non-positive numbers.
// Count how many digits each number has.
// Return the most common digit length.
// Test cases:
//   []int{1, 12, 123, 45, 678, 9, 10} => 2
//   []int{-1, -2, 0} => "no positive numbers"
//   []int{1000, 10000, 100000} => 5
//   []int{} => "no positive numbers"
//   []int{7, 88, 9, 101, 5} => 1

// 26 From a list of integers, collect all values divisible by 3.
// From those, remove all values that are also divisible by 5.
// Return them sorted in descending order.
// Test cases:
//   []int{3, 6, 10, 15, 18, 20} => [18 6 3]
//   []int{5, 10, 20, 25} => []
//   []int{} => []
//   []int{30, 45, 60} => []
//   []int{9, 12, 16} => [12 9]

// 27 From a list of words, remove all words shorter than 5 characters.
// Then, filter out any word that contains the letter 'z' or 'Z'.
// Return the remaining words sorted alphabetically.

// Test cases:
//   []string{"apple", "zebra", "banana", "Zoo", "plum"} => ["apple" "banana"]
//   []string{"zebra", "zoo", "zap"} => []
//   []string{} => []
//   []string{"grape", "melon", "berry"} => ["berry" "grape" "melon"]
//   []string{"aaazz", "bbbbb", "ccccz"} => ["bbbbb"]

// 28 From a list of integers, remove all duplicates.
// Then, split the numbers into two groups: even and odd.
// Return a slice containing first all even numbers, then all odd numbers, both sorted in ascending order.

// Test cases:
//   []int{1, 2, 3, 4, 5, 2, 3, 4} => [2 4 1 3 5]
//   []int{2, 4, 6, 8} => [2 4 6 8]
//   []int{1, 3, 5, 7} => [1 3 5 7]
//   []int{} => []
//   []int{5, 5, 5, 5} => [5]

// 29 From a list of positive integers, filter out values with an odd number of digits.
// From the rest, compute the product of the smallest and largest.
// Return the result as an integer.
// Test cases:
//   []int{10, 22, 333, 4444, 55555} => 220
//   []int{1, 12, 123, 1234, 12345} => 14808
//   []int{1, 3, 5, 7} => "no even-digit numbers"
//   []int{1000, 2000, 3000} => 3000000
//   []int{} => "no even-digit numbers"

// 30 From a list of words, remove all words that appear more than once.
// From the remaining words, collect the ones whose length is a prime number.
// Return them in reverse order of their appearance in the input.
// Test cases:
//   []string{"dog", "cat", "bird", "dog", "eagle"} => ["eagle" "cat"]
//   []string{"red", "red", "red"} => []
//   []string{} => []
//   []string{"hi", "there", "friend", "hi"} => ["friend" "there"]
//   []string{"prime", "seven", "nine", "ten"} => ["nine" "seven" "prime"]
