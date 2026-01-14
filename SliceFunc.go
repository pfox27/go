/* A programme that uses slices to alter a list of comic book heroes. */

package main

import(
		"fmt"
      )

func main(){
	
	heroSlice := []string{"Kal-el", "Mon-el", "Prince Namor", "Atom", "Hulk"}
	fmt.Printf("Original List: %s\n" , heroSlice)
	changeSlice(heroSlice)
	fmt.Printf("Final List: %v\n", heroSlice)
	
}



func changeSlice(heroes [] string){
	
	//Change an item in the list
	
	heroes[4] = "Thor"
	fmt.Printf("changeSlice Function List: %v\n", heroes)
}

