package main

import (
	"fmt"
)

// AreEqual checks if two values are equal using generics
func AreEqual[T comparable](a, b T) bool {
	return a == b
}

func main() {
	intA, intB := 9,9
	areEqualInt := AreEqual(intA, intB)
	fmt.Printf("%d and %d are equal: %t\n", intA, intB, areEqualInt)

	intC, intD := 2, 3
	areEqualInt = AreEqual(intC, intD)
	fmt.Printf("%d and %d are equal: %t\n", intC, intD, areEqualInt)

	floatA, floatB := 7.00, 7.00
	areEqualFloat := AreEqual(floatA, floatB)
	fmt.Printf("%f and %f are equal: %t\n", floatA, floatB, areEqualFloat)

	floatC, floatD := 1.2, 2.3
	areEqualFloat = AreEqual(floatC, floatD)
	fmt.Printf("%f and %f are equal: %t\n", floatC, floatD, areEqualFloat)

}
