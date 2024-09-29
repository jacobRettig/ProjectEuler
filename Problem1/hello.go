package main

import "fmt"

func main() {
	const limit uint32 = 1000000

	var runningSum uint32 = 0
	for i := uint32(3); i < limit; i += 3 {
		runningSum += i
	}
	for i := uint32(5); i < limit; i += 5 {
		if i % 3 != 0 {
			runningSum += i
		}
	}

	fmt.Printf("Running Sum = %d\n", runningSum)
}
