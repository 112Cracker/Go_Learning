package main

import (
	"fmt"
	"sort"
)

// sort int slice
type profits []int

func (p profits) Len() int { return len(p)}
func (p profits) Swap(i, j int) { p[i], p[j] = p[j], p[i]}
func (p profits) Less(i, j int) bool { return p[i] > p[j]}

func main() {
	profits_2018 := profits{11, 33, 22, 44, 77, 55}
	fmt.Printf("Data before sorted: \n%v\n", profits_2018)
	sort.Sort(profits_2018)
	fmt.Printf("Data after sorted: \n%v\n", profits_2018)

	profits_2019 := []int{11, 33, 22, 44, 77, 55}
	fmt.Printf("Data before sorted: \n%v\n", profits_2019)
	sort.Sort(sort.IntSlice(profits_2019)) // conversion from type []int to type IntSlice
	// default implementations sort in ascending order
	fmt.Printf("Data after default sorted: \n%v\n", profits_2019)

	// reverse sort against default implementations
	fmt.Printf("Data before sorted: \n%v\n", profits_2019)
	sort.Sort(sort.Reverse(sort.IntSlice(profits_2019)))
	// reverse sort in descending order
	fmt.Printf("Data after reversed sorte: \n%v\n", profits_2019)


	// another way to sort a []int
	profits_2020 := []int{11, 33, 22, 44, 77, 55}
	fmt.Printf("Data before sorted: \n%v\n", profits_2020)
	sort.Ints(profits_2020)
	fmt.Printf("Data after reversed sorte: \n%v\n", profits_2020)

}
