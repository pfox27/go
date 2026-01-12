/* A program to output a winning message to the screen, using variable substitution. */

package main

import "fmt"

const(
	prizeDay = "Tuesday"
	prizeFund = 22400
		)
		
func main() {

	msg := "%s's prize-winning entry is %d, and the amount of the prize is $%d!!\n"
	winner := 6
	fmt.Printf(msg, prizeDay, winner, prizeFund)
	
}









