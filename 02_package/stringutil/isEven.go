package stringutil

// un-capitalized -> unexported -> invisible to other packages
func isEven(num int) bool {
	return num % 2 == 0
}
