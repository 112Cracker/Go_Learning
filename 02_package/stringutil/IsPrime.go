package stringutil

// capitalized -> exported -> visible to other packages
func IsPrime(num int) bool {
	if num <= 2 {
		return false
	}
	if isEven(num) {
		return false
	}
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}