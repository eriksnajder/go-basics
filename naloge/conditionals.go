package naloge

import "fmt"

func NalogeIzConditionals() {
	fmt.Println("hello world")
	// 1. Check if the boolean variable x is true.
	// If it is, print "x is true".
	x := true
	if x {
		fmt.Println("1. x je true")
	} else {
		fmt.Println("q. x je false")
	}

	// 2. Determine whether the given number is even or odd.
	// Print "Even" or "Odd" accordingly.
	number := 7
	if number%2 == 0 {
		fmt.Println("2. even")
	} else {
		fmt.Println("2. odd")
	}

	// 3. Print the appropriate letter grade based on the score.
	// Use: A for 90+, B for 80-89, C for 70-79, F otherwise.

	score := 85

	//if score >= 70 && score <= 79 {
	//	fmt.Println("3. score is C")
	//} else if score >= 80 && score <= 89 {
	//	fmt.Println("3. score je B")
	//} else if score >= 90 {
	//	fmt.Println("3. score je A")
	//} else {
	//	fmt.Println("3. score je F")

	//}
	if score < 70 {
		fmt.Println("3. F")
	} else if score < 80 {
		fmt.Println("3. C")
	} else if score < 90 {
		fmt.Println("3. B")
	} else {
		fmt.Println("3. A")
	}

	// 4. Check whether the number is positive, negative, or zero.
	// Print the result.
	n := -3
	if n < 0 {
		fmt.Println("4. negative")
	} else if n == 0 {
		fmt.Println("4. zero")
	} else {
		fmt.Println("4. positive")
	}

	// 5. Compare two strings and print whether they are equal or different.
	a := "hello"
	b := "world"

	if a == b {
		fmt.Println("5. equal")
	} else {
		fmt.Println("5. different")
	}

	// 6. Use a switch statement to print the day of the week based on a number (1-5).
	// Print "Weekend" for any other value.

	day := 3

	//if day == 1 {
	//	fmt.Println("6. monday")
	//} else if day == 2 {
	//	fmt.Println("6. tuesday")
	//} else if day == 3 {
	//	fmt.Println("6. wednesday")
	//} else if day == 4 {
	//	fmt.Println("6. thursday")
	//} else if day == 5 {
	//	fmt.Println("6. friday")
	//} else if day == 6 || day == 7 {
	//	fmt.Println("6. weekend")
	//}
	switch day {
	case 1:
		fmt.Println("6. monday")
	case 2:
		fmt.Println("6. tuesday")
	case 3:
		fmt.Println("6. wednesday")
	case 4:
		fmt.Println("6. thursday")
	case 5:
		fmt.Println("6. friday")
	default:
		fmt.Println("6. weekend")
	}

	// 7. Check whether the number is divisible by both 3 and 5.
	num := 15
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("7. divisible")
	} else {
		fmt.Println("7. undivisible")
	}

	// 8. Determine if a given year is a leap year.
	// A leap year is divisible by 4 but not by 100, unless also divisible by 400.
	year := 2024
	if year%4 == 0 {
		if year%100 != 0 {
			fmt.Println("8. leap year")
		} else {
			if year%400 == 0 {
				fmt.Println("8. leap year")
			} else {
				fmt.Println("8. not leap year")
			}
		}
	} else {
		fmt.Println("8. not leap year")
	}

	// 9. Check if the given character is a vowel (a, e, i, o, u) or consonant.
	ch := 'e'

	switch ch {
	case 'a':
		fmt.Println("9. vowel")
	case 'e':
		fmt.Println("9. vowel")
	case 'i':
		fmt.Println("9. vowel")
	case 'o':
		fmt.Println("9. vowel")
	case 'u':
		fmt.Println("9. vowel")
	default:
		fmt.Println("9. consonant")
	}

	// 10. Check if a person is eligible to vote (age 18 or older).
	age := 16

	if age < 18 {
		fmt.Println("10. not eligible")
	} else {
		fmt.Println("10. eligible")
	}

	// 11. Use a switch to respond to a traffic light color.
	// For example: "Stop" for red, "Go" for green, "Wait" for yellow.
	color := "green"

	switch color {
	case "red":
		fmt.Println("11. Stop")
	case "yellow":
		fmt.Println("11. Wait")
	case "green":
		fmt.Println("11. Go")
	default:
		fmt.Println("11. Watch the Signs!")
	}

	// 12. Compare three integers and print the largest one.
	xx := 3
	y := 7
	z := 5

	res := xx
	if y > res {
		res = y
	}
	if z > res {
		res = z
	}
	fmt.Printf("12. %d\n", res)

	//list := []int{xx, y, z}
	//res := list[0]
	//for _, i := range enumerate(list) {
	//	if i > res {
	//		res = i
	//	}
	//}

	// 13. Check if a number is within the range 10 to 20 (inclusive).
	val := 15
	if val >= 10 && val <= 20 {
		fmt.Println("13. inclusive")
	} else {
		fmt.Println("13. not inclusive")
	}
	// 14. Classify the temperature:
	//below 0: "Freezing", 0-20: "Cold", 21-30: "Warm", above 30: "Hot"
	temperature := 30
	classify := ""
	if temperature <= 0 {
		classify = "Freezing"
	} else if temperature <= 20 {
		classify = "Cold"
	} else if temperature <= 30 {
		classify = "Warm"
	} else if temperature > 30 {
		classify = "Hot"
	}
	fmt.Printf("14.%s\n", classify)

	// 15. Use a switch statement to describe a grade:
	// A = "Excellent", B = "Good", C = "Average", F = "Fail"
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("15. Excellent")
	case "B":
		fmt.Println("15. Good")
	case "C":
		fmt.Println("15. Average")
	case "F":
		fmt.Println("15. Fail")
	}

	// 16. Check whether a given string is empty or not.
	text := ""
	if text == "" {
		fmt.Println("16. empty")
	} else {
		fmt.Println("16. not empty")
	}

	// 17. Check if a number is divisible by 2, 5, both, or neither.
	numm := 10
	if numm%2 != 0 && numm%5 != 0 {
		fmt.Println("17. neither")
	} else if numm%2 == 0 && numm%5 == 0 {
		fmt.Println("17. both")
	} else if numm%5 == 0 {
		fmt.Println("17. divisible by 5")
	} else if numm%2 == 0 {
		fmt.Println("17. divisible by 2")
	}

	// 18. Use a switch to print an emoji based on mood:
	// "happy" = 😀, "sad" = 😢, "angry" = 😡, default = 😐
	mood := "happy"

	switch mood {
	case "happy":
		fmt.Println("18. 😀")
	case "sad":
		fmt.Println("18. 😢")
	case "angry":
		fmt.Println("18. 😡")
	default:
		fmt.Println("18. 😐")
	}

	// 19. Check if a password is long enough (8 characters or more).
	password := "secret"

	if len(password) >= 8 {
		fmt.Println("19. long enough")
	} else {
		fmt.Println("19. not long enough")
	}

	// 20. Use if-else to determine if the given number is a multiple of 10.
	number = 10
	if number%10 == 0 {
		fmt.Println("20. multiple of 10")
	} else {
		fmt.Println("20. not multiple of 10")
	}

	// 21. Check if a number is odd AND greater than 10.
	number = 13
	if number%2 != 0 && number > 10 {
		fmt.Println("21. Check OK")
	} else {
		fmt.Println("21. Check not OK")
	}

	// 22. Use a switch to simulate a basic calculator with operators "+", "-", "*", "/".
	operator := "+"
	aaa := 5
	bbb := 3

	switch operator {
	case "+":
		ccc := aaa + bbb
		fmt.Printf("22. %d + %d = %d\n", aaa, bbb, ccc)
	case "-":
		ccc := aaa - bbb
		fmt.Printf("22. %d - %d = %d\n", aaa, bbb, ccc)
	case "*":
		ccc := aaa * bbb
		fmt.Printf("22. %d * %d = %d\n", aaa, bbb, ccc)
	case "/":
		ccc := float64(aaa) / float64(bbb)
		fmt.Printf("22. %d / %d = %f\n", aaa, bbb, ccc)
	}

	// 23. Determine if a number is a single-digit, two-digit, or larger.
	number = 42
	numberS := fmt.Sprintf("%d", number)
	if len(numberS) == 1 {
		fmt.Println("23. single-digit")
	} else if len(numberS) == 2 {
		fmt.Println("23. two-digit")
	} else if len(numberS) > 2 {
		fmt.Println("23. larger")
	}
	// 24. Print a special message if a user's name matches a secret name.
	name := "Harry"
	user_name := "Harry"
	if user_name == name {
		fmt.Println("24. ko se kaka, se tudi lula")
	} else {
		fmt.Println("24. wrong user")
	}
	// 25. Use if-else to check if a number is prime (up to 20 is enough).
	// number := 17

	// 26. Check if two numbers are both positive.
	aa := 5
	bb := 12

	if aa > 0 {
		fmt.Println("26. a = positive")
	} else {
		fmt.Println("26. a = negative")
	}
	if bb > 0 {
		fmt.Println("26. b = positive")
	} else {
		fmt.Println("26. b = negative")
	}
	// 27. Use a switch to determine the season based on the month number (1-12).
	month := 4
	fmt.Print("27. ")
	switch month {
	case 1:
		fmt.Println("winter")
	case 2:
		fmt.Println("winter")
	case 3:
		fmt.Println("spring")
	case 4:
		fmt.Println("spring")
	case 5:
		fmt.Println("spring")
	case 6:
		fmt.Println("summer")
	case 7:
		fmt.Println("summer")
	case 8:
		fmt.Println("summer")
	case 9:
		fmt.Println("autumn")
	case 10:
		fmt.Println("autumn")
	case 11:
		fmt.Println("autumn")
	case 12:
		fmt.Println("winter")
	}
	// 28. Check if a given year is in the 20th or 21st century.
	year = 1999
	if year < 2001 {
		fmt.Println("28. 20th century")
	} else {
		fmt.Println("28. 21th century")
	}

	// 29. Check if a temperature is suitable for swimming (between 22 and 30°C).
	temperature = 25

	if temperature > 22 && temperature < 30 {
		fmt.Println("29. swimmable")
	} else {
		fmt.Println("29. not swimmable")
	}

	// 30. Use a switch to simulate a mini menu system:
	// "1" = "Start Game", "2" = "Load Game", "3" = "Exit"
	choice := "2"

	switch choice {
	case "1":
		fmt.Println("30. Start game")
	case "2":
		fmt.Println("30. Load game")
	case "3":
		fmt.Println("30. Exit")
	}
}
