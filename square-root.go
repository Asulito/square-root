package main

import (
	"fmt"
	"math"
)

/*
	Save to a variable if the root of i is less than x
	If the root of i is geater than x then return the last saved variable
*/
func Sqrt(x float64) float64 {
	var i, root, result = 1.0, 1.0, 1.0

	for i = 1; i < x; i++ {
		if root = math.Pow(i, 2); root < x {
			result = i
		} else {
			return result
		}
	}
	return result
}

func main() {
	fmt.Println(Sqrt(30.89))
}
