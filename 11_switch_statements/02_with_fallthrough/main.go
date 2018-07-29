package main
// CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/blob/master/11_switch-statements/02_fallthrough/main.go
import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}