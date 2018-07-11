package main

import (
	"fmt"
)

func tozero(z *int) {
	*z = 0
}

func main() {
	i := 1
	fmt.Printf("%v\n", i)
	fmt.Println(&i)

	var j *int
	j = &i
	fmt.Println(j)

	*j = 2
	fmt.Println(i)

	tozero(j)
	fmt.Println(i)
}
