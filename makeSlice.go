/* A programme that creates a slice to store names of programming languages. */

package main

import(
		"fmt"
      )

func main(){
	
	languageSlice := make([]string, 0, 5)
	fmt.Printf("Initial Length of Slice: %d\n" , len(languageSlice))
	fmt.Printf("Capacity: %d\n" , cap(languageSlice))
	languageSlice = append(languageSlice, "Java", "Python", "C", "C++", "Go", "Javascript")
	fmt.Printf("Contents: %s\n" , languageSlice)

	fmt.Printf("Current Length of Slice: %d\n" , len(languageSlice))
	fmt.Printf("Current Capacity: %d\n" , cap(languageSlice))

	languageSliceCopy := make([]string, 10)
	copy(languageSliceCopy, languageSlice)

	languageSliceCopy[2] = "HTML"
	fmt.Printf("languageSlice: %s\n", languageSlice)
	fmt.Printf("languageSliceCopy: %s\n", languageSliceCopy)

}

