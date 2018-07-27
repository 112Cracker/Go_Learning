// CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/blob/master/06_constants/07_iota/main.go
package main

import "fmt"

const (
	// _ is a blank identifier to throw away unnecessary value
	// iota automatic increment by 1 every time appears in sequence
	_ = iota // 0
	// bit shifting in Golang
	KB = 1 << (iota * 10) // 1 << (1 * 10) get KiloBytes
	MB = 1 << (iota * 10) // 1 << (1 * 20) get MegaBytes
	GB = 1 << (iota * 10) // 1 << (1 * 30) get GigBytes
	TB = 1 << (iota * 10) // 1 << (1 * 40) get TeraBytes
)

func main() {
	// Print KB, MB, GB, TB in binary and decimal
	// %b for binary format
	// %d for decimal format
	fmt.Printf("binary\t\tdecimal\n")

	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\t\n", KB)

	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t\n", MB)

	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t\n", GB)

	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\t\n", TB)
}