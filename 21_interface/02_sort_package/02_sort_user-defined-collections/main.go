package main

import (
	"fmt"
	"sort"
)

// sort user-defined collections user-defined struct by field
type person struct {
	weight int
	height int
	name string
}

func (p person) String() string {
	return fmt.Sprintf("The person is %d cm tall and weighs %d kg\n",
		p.height, p.weight)
}

// byHeight user-defined collections
// underlying data structure is a person slice
type byHeight []person

func (bh byHeight) Len() int { return len(bh)}
func (bh byHeight) Swap(i, j int) { bh[i], bh[j] = bh[j], bh[i]}
// ">" in descending order
func (bh byHeight) Less(i, j int) bool { return bh[i].height > bh[j].height}

// byWeight user-defined collections
// underlying data structure is a person slice
type byWeight []person

func (bw byWeight) Len() int { return len(bw)}
func (bw byWeight) Swap(i, j int) { bw[i], bw[j] = bw[j], bw[i]}
// "<" in ascending order
func (bw byWeight) Less(i, j int) bool { return bw[i].weight < bw[j].weight}


func main() {
	pre_atheletes := []person {
		{75, 180, "Tom"},
		{60, 165, "Jerry"},
		{80, 185, "Caffe"},
		{90, 187, "Popan"},
	}

	fmt.Println("Data before sorted: ")
	fmt.Println(pre_atheletes)
	sort.Sort(byHeight(pre_atheletes)) // conversion to byHeight
	fmt.Println("Data after sorted: ")
	fmt.Println(pre_atheletes)
}