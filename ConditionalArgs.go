/* A program to take four arguments as variables and then use if statements. */

package main

import (
	"fmt"
	"os"
	)
		
func main() {

	val1 := os.Args[1]
	val2 := os.Args[2]
	val3 := os.Args[3]
	val4 := os.Args[4]
	fmt.Printf("%s, %s, %s, %s\n", val1, val2, val3, val4)
}









