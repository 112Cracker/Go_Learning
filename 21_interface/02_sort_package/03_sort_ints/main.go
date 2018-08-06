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
}
