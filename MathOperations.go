/* A program to output a winning message to the screen, using variable substitution. */

package main

import "fmt"

var(
	integer1 = 5
	integer2 = -3
	float1 = -8.1
	float2 = 6.32
		)
		
func main() {

	fmt.Printf("(%d) + (%d) = %d\n",  integer1, integer2, integer1 + integer2)
	fmt.Printf("(%f) / (%f) = %f\n",  float1, float2, float1 + float2)
}









