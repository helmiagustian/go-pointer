package main

import (
	"fmt"
)

func tozero(z *int) {
	*z = 0 // assign value 0 to memory that have address value of z
}

func main() {
	i := 1
	fmt.Printf("%v\n", i) // print variable i value
	fmt.Println(&i)       // print variable i address value

	var j *int     // create pointer variable j
	j = &i         // variable j now have variable i address value
	fmt.Println(j) // print variable j value

	*j = 2         // assign value 2 to memory that have address value of j
	fmt.Println(i) // print variable i value

	tozero(j)
	fmt.Println(i)
}
