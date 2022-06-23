package Task3

import (
	"fmt"
	"math"
)

func quadratic(a float64, b float64, c float64) (float64, float64) {
	var root1, root2 float64
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)

	// calculate the roots of the polynomial
	// x^2 + bx + c using the quadratic formula

	discriminant := (b * b) - (4 * a * c)
	d := math.Sqrt(discriminant)

	root1 = ((-b + d) / (2 * a))
	root2 = ((-b - d) / (2 * a))
	return root1 , root2
}
