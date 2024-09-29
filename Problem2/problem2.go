package main

import "fmt"

func main() {
	const limit uint32 = 4e6
	var a uint32 = 1
	var b uint32 = 2
	var c uint32
	var runningSum uint32 = 0

	for b < limit {
		if b % 2 == 0 {
			runningSum += b
		}
		c = b
		b = a + b
		a = c
	}

	fmt.Printf("Running Sum = %d\n", runningSum)
}
