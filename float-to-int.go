package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Enter any number: ")
	_, _ = fmt.Scan(&x)   //add 'blanks' to avoid error caused due to 'assigned but unused vairables'
	var y int = int(x)
	fmt.Println("The integer part of the number typed is:", y)
}
