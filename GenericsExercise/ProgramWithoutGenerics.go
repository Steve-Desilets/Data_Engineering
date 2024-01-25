package main
import (
	"fmt"
)
// AreIntsEqual checks if two integers are equal
func AreIntsEqual(a, b int) bool {
	return a == b
}
// AreFloatsEqual checks if two floats are equal
func AreFloatsEqual(a, b float64) bool {
	return a == b
}
// run the main function that does not leverage generics
func main() {
	intA, intB := 9, 9
	areEqualInt := AreIntsEqual(intA, intB)
	fmt.Printf("%d and %d are equal: %t\n", intA, intB, areEqualInt)

	intC, intD := 2, 3
	areEqualInt = AreIntsEqual(intC, intD)
	fmt.Printf("%d and %d are equal: %t\n", intC, intD, areEqualInt)

	floatA, floatB := 7.00, 7.00
	areEqualFloat := AreFloatsEqual(floatA, floatB)
	fmt.Printf("%f and %f are equal: %t\n", floatA, floatB, areEqualFloat)

	floatC, floatD := 1.2, 2.3
	areEqualFloat = AreFloatsEqual(floatC, floatD)
	fmt.Printf("%f and %f are equal: %t\n", floatC, floatD, areEqualFloat)
}
