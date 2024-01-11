package main

import (
	"fmt"
	"math"
)

func main() {
	var option string
	fmt.Println("What option do you want to do? (Solve equation(a) or find the hypotenuse(b)")
	fmt.Scan(&option)
	switch {
	case option == "a":
		var (
			a float64
			b float64
			c float64
		)
		fmt.Println("Calculator of a quadratic equations\n Enter A coefficient:")
		fmt.Scan(&a)
		fmt.Println("Enter B coefficient")
		fmt.Scan(&b)
		fmt.Println("Enter C coefficient")
		fmt.Scan(&c)

		// Solving logic

		D := (b * b) - 4*a*c
		fmt.Printf("Your D equals: %v\n", D)
		if D > 0 {
			var (
				x1 float64
				x2 float64
			)
			x1 = (-b + math.Sqrt(D)) / (2 * a)
			x2 = (-b - math.Sqrt(D)) / (2 * a)
			fmt.Printf("Your quadratic equation has 2 roots, x1: %v, x2:%v\n", x1, x2)
		} else if D == 0 {
			var x float64
			x = (-b) / (2 * a)
			fmt.Printf("Your quadratic equation has 1 root: %v", x)
		} else if D < 0 {
			fmt.Printf("Your D equals: %v, which is smaller than 0 and if D < 0, quadratic equation can not have any root, it is impossible due to math laws.\n", D)
		}
		main()
	case option == "b":
		var (
			a float64
			b float64
		)
		fmt.Println("Calculator of a hypotenuse\n Enter A coefficient:")
		fmt.Scan(&a)
		fmt.Println("Enter B coefficient")
		fmt.Scan(&b)
		// Solving logic
		if a < 0 && b < 0 {
			fmt.Printf("Coefficients can not be smaller than 0")
		} else {
			var result float64 = math.Sqrt(a*a + b*b)
			fmt.Printf("Result is %v\n", result)
		}
		main()
	}
}
