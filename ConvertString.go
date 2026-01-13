/* A program to convert a string to an integer and then an integer to a string. */

package main

import "fmt"
import "strconv"

var(
	string1 = "14"
	integer1 = 87
		)
		
func main() {

	toIntegerResult, toIntErr := strconv.Atoi(string1)
	toStringResult := strconv.Itoa(integer1)
	fmt.Print("The converted string is now ")
	fmt.Print(toIntegerResult)
	fmt.Println(" as an integer.")
	fmt.Println("The converted integer is now " + toStringResult + " as a string.")
	fmt.Print("The errors encountered when converting the string to an integer are: ")
	fmt.Print(toIntErr)
	fmt.Println(".")
}









