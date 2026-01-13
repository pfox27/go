/* A programme to accept four numerical arguments and check to see if they are odd or even. */

package main

import(
		"fmt"
		"os"
		"strconv"
      )

func main(){

	val1, _ := strconv.Atoi(os.Args[1])
	val2, _ := strconv.Atoi(os.Args[2])
  	val3, _ := strconv.Atoi(os.Args[3])
	val4, _ := strconv.Atoi(os.Args[4])
	
	fmt.Println("The four numerical values are: ")
	fmt.Printf("%d, %d, %d, %d\n", val1, val2, val3, val4)
	fmt.Println("\n")
	fmt.Printf("%s is %s\n", os.Args[1], oddOrEven(os.Args[1]))
	fmt.Printf("%s is %s\n", os.Args[2], oddOrEven(os.Args[2]))
	fmt.Printf("%s is %s\n", os.Args[3], oddOrEven(os.Args[3]))
	fmt.Printf("%s is %s\n", os.Args[4], oddOrEven(os.Args[4]))


}

func oddOrEven(value string) string{
	num, _ := strconv.Atoi(value)
		if num % 2 == 0 {
			return "even"
		} else{
			return "odd"
		}
}

