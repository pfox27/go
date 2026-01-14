/* A programme to traverse (parse) a string and count the number of 
 * letters in three separate groups in the alphabet. */

package main

import(
		"fmt"
		"os"
      )

func main(){

	numAtoI, numJtoR, numStoZ, numSpaces, numOther := 0, 0, 0, 0, 0
	sentence := os.Args[1] 
	
	for i:= 0; i < len(sentence); i++{
		letter := string(sentence[i])
		fmt.Print(" ", string(sentence[i]), " ", sentence[i])
		
		switch letter{
			case "a", "b", "c", "d", "e", "f", "g", "h", "i":
			numAtoI++
			case "j", "k", "l", "m", "n", "o", "p", "q", "r":
			numJtoR++
			case "s", "t", "u", "v", "w", "x", "y", "z":
			numStoZ++
			case " ":
			numSpaces++
			default:
			numOther++
		}
	}
		fmt.Printf("Number of letters from A to I: %d\n", numAtoI)
		fmt.Printf("Number of letters from J to R: %d\n", numJtoR)
		fmt.Printf("Number of letters from S to Z: %d\n", numStoZ)
		fmt.Printf("Number of spaces: %d\n", numSpaces)
		fmt.Printf("Number of \"other\": %d\n", numOther)
}

