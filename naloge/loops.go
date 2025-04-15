package naloge

import (
	"fmt"
	"strings"
)

func NalogeIzLoops() {
	// 0. Prints elements of a slice and map
	l := []int{1, 2, 3, 4, 5}
	m := map[string]string{"Hello": "World", "name": "erik"}

	for _, item := range l {
		fmt.Println(item)
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	// 1. Use a for loop to print numbers from 1 to 10
	// No variables needed
	fmt.Println("----------")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. Check if a number is even or odd using if-else
	fmt.Println("----------")
	num := 7

	if num%2 != 0 {
		fmt.Println("2. number is odd")
	} else {
		fmt.Println("2. number is even")
	}
	// 3. Print all even numbers from 1 to 20 using a loop and condition
	// No variables needed
	fmt.Println("----------")
	for i := 0; i <= 20; i++ {
		if i%2 == 0 && i != 0 {
			fmt.Println(i)
		}
	}
	// 4. Use a loop to calculate the sum of numbers from 1 to 100
	fmt.Println("----------")
	sum := 0

	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 5. Use a loop to count how many numbers from 1 to 50 are divisible by 3
	count := 0
	fmt.Println("----------")
	for i := 0; i <= 50; i++ {
		if i%3 == 0 {
			count += 1
		}
	}
	fmt.Println(count)

	// 6. Write a loop that prints "Hello" 5 times
	// No variables needed
	fmt.Println("----------")
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}
	// 7. Print the multiplication table of 5 (from 5x1 to 5x10)
	// No variables needed
	fmt.Println("----------")
	for i := 0; i <= 10; i++ {
		if i != 0 {
			rez := 5 * i
			fmt.Printf("%d X %d = %d\n", 5, i, rez)
		}
	}
	// 8. Use an if statement to check if a person is eligible to vote (age >= 18)
	age := 17
	fmt.Println("----------")
	if age >= 18 {
		fmt.Println("eligible to vote")
	} else {
		fmt.Println("not eligible to vote")
	}
	// 9. Print numbers from 1 to 20, but skip even numbers using continue
	// No variables needed
	fmt.Println("----------")
	for i := 0; i <= 20; i++ {
		if i%2 == 0 || i == 0 {
			continue
		}
		fmt.Println(i)
	}
	// 10. Sum only odd numbers between 1 and 100
	sum = 0
	fmt.Println("----------")
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println(sum)
	// 11. Create a countdown from 10 to 1 using a loop
	// No variables needed
	fmt.Println("----------")
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
	// 12. Use a nested loop to print a 5x5 star grid
	// No variables needed
	fmt.Println("----------")
	visina := 5
	sirina := 5

	for vrstice := 0; vrstice <= visina; vrstice++ {
		vrsta := ""
		for zvezdice := 0; zvezdice < sirina; zvezdice++ {
			vrsta = fmt.Sprintf("%s*", vrsta)
		}
		fmt.Println(vrsta)
	}
	// 13. Check if a number is positive, negative or zero
	num = -3
	fmt.Println("----------")

	if num == 0 {
		fmt.Println("number is zero")
	} else if num > 0 {
		fmt.Println("number is positive")
	} else {
		fmt.Println("number is negative")
	}
	// 14. Use a loop to find the factorial of a number
	n := 5
	factorial := 0
	fmt.Println("----------")

	for i := 0; i <= n; i++ {
		factorial += i
	}
	fmt.Println(factorial)

	// 15. Loop through numbers 1 to 30 and print "Fizz" if divisible by 3, "Buzz" if divisible by 5, and "FizzBuzz" if both
	// No variables needed
	fmt.Println("----------")
	for i := 0; i <= 30; i++ {
		if i != 0 {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else if i%3 == 0 {
				fmt.Println("Fizz")
			} else {
				fmt.Println(i)
			}
		}
	}
	// 16. Use a loop to find the largest number in a slice
	numbers := []int{4, 10, 2, 99, 23}
	spremenljivka := 0
	fmt.Println("----------")
	for _, number := range numbers {
		if number > spremenljivka {
			spremenljivka = number
		}
	}
	fmt.Println(spremenljivka)

	// 17. Ask the user for 5 test scores and calculate the average using a loop
	scores := []int{80, 90, 85, 75, 95}
	sum = 0

	fmt.Println("----------")
	for _, number := range scores {
		sum += number
	}
	avg := sum / len(scores)
	fmt.Printf("%d / %d = %d\n", sum, len(scores), avg)

	// 18. Use an if-else to determine the grade based on score
	score := 83
	fmt.Println("----------")
	if score < 70 {
		fmt.Println("F")
	} else if score < 80 {
		fmt.Println("C")
	} else if score < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
	// 19. Create a loop that doubles a number until it’s greater than 1000
	n = 2
	i := 1
	fmt.Println("----------")
	for {
		i *= 2
		fmt.Println(i)
		if i > 1000 {
			break
		}
	}

	// 20. Loop through a slice and print only the numbers greater than 50
	numbers = []int{23, 67, 45, 89, 10, 102}
	fmt.Println("----------")
	for _, number := range numbers {
		if number > 50 {
			fmt.Println(number)
		}
	}
	// 21. Use a loop to reverse a string
	input := "hello"
	reversed := ""
	crke := strings.Split(input, "")
	fmt.Println("----------")

	for i := len(crke) - 1; i >= 0; i-- {
		reversed = fmt.Sprintf("%s%s", reversed, crke[i])
	}
	fmt.Println(reversed)
	// 22. Create a simple login check using if-else
	username := "admin"
	password := "pass123"

	fmt.Println("----------")
	if username == "admin" {
		if password == "pass123" {
			fmt.Println("Welcome, admin!")
		} else {
			fmt.Println("Wrong password!")
		}
	}
	// 23. Use a for loop with range to print all characters in a string
	input = "golang"
	fmt.Println("----------")
	crke = strings.Split(input, "")

	for _, i := range crke {
		fmt.Println(i)
	}

	// 24. Create a loop that finds the sum of digits of a number
	num = 1234
	sum = 0
	fmt.Println("----------")

	for {
		currentDigit := num % 10
		num /= 10
		sum += currentDigit
		if num == 0 {
			break
		}
	}
	fmt.Println(sum)
	// 25. Create a menu system using switch:
	//"a" = "Add Item"
	//"b" = "Remove Item"
	//"c" = "Checkout"
	choice := "b"
	fmt.Println("----------")

	switch choice {
	case "a":
		fmt.Println("Add Item")
	case "b":
		fmt.Println("Remove Item")
	case "c":
		fmt.Println("Checkout")
	}
	// 26. Print all prime numbers between 1 and 50
	// No variables needed
	fmt.Println("----------")

	for stevilo := 1; stevilo <= 50; stevilo++ {
		prime := true
		for i := 2; i <= (stevilo/2)+1; i++ {
			if stevilo%i == 0 {
				prime = false
			}
		}
		if prime {
			fmt.Println(stevilo, prime)
		}

	}

	// 27. Find the smallest number in a slice
	numbers = []int{4, 10, 2, 99, 23}
	fmt.Println("----------")

	min := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}

	}
	fmt.Println(min)
	// 28. Loop through numbers 1 to 100 and stop the loop if number is divisible by 77
	// No variables needed
	fmt.Println("----------")
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		if i >= 77 && i%77 == 0 {
			break
		}
	}
	// 29. Use a loop to generate the first 10 numbers of the Fibonacci sequence
	// a := 0, b := 1
	fmt.Println("----------")
	x := 0
	y := 1
	z := 0

	fmt.Println(x)
	fmt.Println(y)
	for i := 1; i <= 10; i++ {
		z = x + y
		x = y
		y = z
		fmt.Println(z)
	}
}
