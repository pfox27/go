/* A program to output the result of concatenating some strings and substrings. */

package main

import "fmt"

var(
	string1 = "Constant"
	string2 = "inople"
		)
		
func main() {

	fmt.Println(string1[0:3] + string2[2:6])
	fmt.Printf("%s + %s = %s\n", string1, string2, string1 + string2)
	
}









