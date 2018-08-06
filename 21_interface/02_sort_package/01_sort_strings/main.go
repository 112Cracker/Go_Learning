package main

import (
	"fmt"
	"sort"
)

// sort string slice
type editors []string

func (e editors) Len() int { return len(e)}
func (e editors) Swap(i, j int) { e[i], e[j] = e[j], e[i]}
func (e editors) Less(i, j int) bool { return e[i] < e[j]}

func main() {
	editors_for_HP := editors{"Mary", "John", "Sherry", "Tom", "Zack"}
	fmt.Printf("Data before sorted: \n%v\n", editors_for_HP)
	sort.Sort(editors_for_HP)
	fmt.Printf("Data after sorted: \n%v\n", editors_for_HP)
}