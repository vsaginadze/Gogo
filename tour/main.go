// Hello, 世界

/*
package main

import "fmt" # formatting package

func main() {
    fmt.Println("Hello, 世界") # formatting package includes printing method
}
*/

// Exported Names
// Exported names starts with the capital letter, like Pi,
/*
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("PI: ", math.Pi)
}
*/

// Functions
// add(int x, int y) int <- returing value {
// }

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// if variables share type you can ommit the redeclaration
func subtract(x, y, z int) int {
	return x - y - z
}

// functions can return any number of results
func swap(x, y, z string) (string, string, string) {
	return z, x, y
}

// named return values
func split(sum int) (x, y int) {
	x = sum - (sum % 10)
	y = sum - (sum - 1)
	return
}

func main() {
	ten, one := split(12)
	fmt.Println(ten, one)
}
