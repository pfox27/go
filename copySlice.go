/* A programme that uses slices to alter a list of comic book heroes. */

package main

import(
		"fmt"
      )

func main(){
	
	languageSlice := make([]string, 2, 5)
	fmt.Printf("Initial Length of Slice: %d\n" , len(languageSlice))
	fmt.Printf("Capacity: %d\n" , cap(languageSlice))
	languageSlice = append(languageSlice, "Java", "Python", "C", "C++", "Go", "Javascript")
	fmt.Printf("Contents: %s\n" , languageSlice)

	fmt.Printf("Current Length of Slice: %d\n" , len(languageSlice))
	fmt.Printf("Current Capacity: %d\n" , cap(languageSlice))

}

