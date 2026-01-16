/* A programme that creates a map to store the ages of various celebrities. */

package main

import(
		"fmt"
      )

func main(){
	
	celeb := make(map[string]int)
	celeb["Flea"] = 61
	celeb["Borgnine"] = 92
	celeb["Arthur"] = 94
	celeb["Young"] = 79
	celeb["McCartney"] = 84
	celeb["Torvalds"] = 68
	celeb["Lopez"] = 52
	celeb["Notaro"] = 55

	fmt.Printf("Linus Torvalds is %d years old.\n" , celeb["Torvalds"])	
	fmt.Printf("Bea Arthur is %d years old.\n" , celeb["Arthur"])
	fmt.Printf("Ernest Bornine is %d years old.\n" , celeb["Borgnine"])

	for i:= 0; i < 4; i++{
		fmt.Printf("\nRun Number: %d\n" , i)
		for key, value := range celeb {
			fmt.Printf("%s is %d years old.\n", key, value)
		}
	}	
}

