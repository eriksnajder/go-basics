package naloge

import (
	"fmt"
	"strings"
)

func NalogeIzPrintanja() {
	err := fmt.Errorf("oops")
	formatted := fmt.Sprintf("integers: %d, floats: %f, strings: %s, quoted: %q, verbose: %v", 10, 3.14, "Hello world!", "quoted text", err)
	fmt.Println(formatted)

	// 1. Print a Greeting
	// Task: Write a program that prints: Hello, <your name>! Welcome to Go.

	yourName := "Erik"
	fmt.Println("Hello,", yourName+"!", "Welcome to Go.")

	// 2. Concatenate Strings
	// Task: Combine first name and last name into a full name and print:
	// Full Name := John Doe
	// firstName := John
	// lastName := Doe

	firstName := "Erik"
	lastName := "Snajder"
	fullName := fmt.Sprintf("%s %s\n", firstName, lastName)
	fmt.Println(fullName)

	// 3. String Length
	// Task: Print how many characters does inputString have.

	inputString := "some string foobar"
	a := len(inputString)
	fmt.Printf("inputStrings has %d characters.\n", a)

	// askGoogle := how do i ask for string input in golang?
	// askGoogle := how do i check for string lenght in golang?

	// 4. Change Case
	// Task: Convert a string to uppercase and lowercase and print both.
	lowercaseString := "im a lowercase"
	uppercaseString := "IM AN UPPERCASE"

	lowerCase := strings.ToUpper(lowercaseString)
	fmt.Println(lowerCase)

	fmt.Println(strings.ToLower(uppercaseString))

	// 5. Replace Words
	// Task: Replace the word "boring" with "awesome" in a string like:
	// Go is boring. → Go is awesome.

	originalSentence := "Go is boring."
	newSentence := strings.Replace(originalSentence, "boring", "awesome", -1)
	fmt.Println(originalSentence + "\n" + newSentence)

	// 6. Format Output
	// Task: Format and print a sentence like:
	// Name: Alice | Age: 23 | Country: Canada
	name := "Alice"
	age := 23
	country := "Canada"

	fmt.Printf("Name: %s | Age: %d | Country: %s\n", name, age, country)

	// 7. Print Special Characters
	// Instruction: Print a sentence with quotation marks and a line break

	fmt.Printf("Hello hubba bubba babba, howssa goin'ssa?\nHello habba babba bebba, messa goin'ssa gooooooooodssa!!!\n")

	// 8: Combine Multiple Variables into a Sentence
	// Instruction: Use several string variables to make and print a sentence.
	aa := "The"
	b := "sky"
	c := "is"
	d := "blue"

	//sentenceOne := fmt.Sprintf("%s %s %s %s", aa, b, c, d)
	fmt.Println(aa, b, c, d)
	fmt.Println(d, aa, b, c)
}
