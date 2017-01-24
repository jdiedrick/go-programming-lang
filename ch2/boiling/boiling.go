//prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g degrees F or %g degrees C\n", f, c)
	// output: boiling point = 212 degrees F or 100 degrees c
}
