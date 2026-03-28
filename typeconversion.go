package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

/* Output: 3 4 5
note: In Go, you cannot directly convert an int to a float64 or a float64 to a uint without an explicit conversion. You need to use the appropriate type conversion functions to perform these conversions.
note: The math.Sqrt function returns a float64, so we need to convert the result to uint before assigning it to z. This is because z is declared as a uint, and we cannot assign a float64 value directly to it without an explicit conversion.
note: The math.Sqrt function returns a float64, so we need to convert the result to uint before assigning it to z. This is because z is declared as a uint, and we cannot assign a float64 value directly to it without an explicit conversion.
note: The math.Sqrt function returns a float64, so we need to convert the result to uint before assigning it to z. This is because z is declared as a uint, and we cannot assign a float64 value directly to it without an explicit conversion.
*/
