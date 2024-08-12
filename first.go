package first

import (
	"fmt"
	"math"
)

func Greet() {
	fmt.Printf("Good Morning")
	fmt.Printf("value of e = %e", math.E)
}

func Evening() {
	fmt.Printf("Good Evening")
	fmt.Printf("value of e = %e", math.Sqrt(25.7))
}

func Day() {
	fmt.Printf("Good Day")
	fmt.Printf("value of e = %e", math.Sqrt(25))
}
