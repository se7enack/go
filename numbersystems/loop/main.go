package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
	for i := 97; i < 123; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
}
