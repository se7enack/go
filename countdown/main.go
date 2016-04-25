package main

import "fmt"

func main()  {
	for i := 10000000; i > 0; i-- {
		print("\033[H\033[2J")
		fmt.Printf("Countdown: ")
		fmt.Println(i)

	}
	print("\033[H\033[2J")
	fmt.Println("Boom!")
}



