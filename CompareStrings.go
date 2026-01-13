/* A program to compare two string values and output the results of the comparison. */

package main

import "fmt"

var(
	string1 = "elm"
	string2 = "maple"
		)
		
func main() {

	if string1 > string2 {
		fmt.Println(string1 + " is greater than " + string2 + ".")	
}	else {
		fmt.Println(string2 + " is greater than " + string1 + ".")
	}
}









