package greeting

import "fmt"

// Salutation is an exported type.
type Salutation struct {
	name     string
	greeting string
}

// Printer is an exported type.
type Printer func(string)

// Greet is an exported function.
func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	if isFormal {
		do(message)
	}

	do(alternate)
}

// CreateMessage is an exported function.
func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!" + " " + name
	return
}

// Print is an exported function.
func Print(s string) {
	fmt.Print(s)
}

// PrintLine is an exported function.
func PrintLine(s string) {
	fmt.Println(s)
}

// CreatePrintCustom is an exported function.
func CreatePrintCustom(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
