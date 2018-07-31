package main

import "fmt"

func main() {
	// var uninitialized map
	var map_var_nil  map[string]string
	fmt.Println(map_var_nil == nil) // => true

	// var make initialized map
	var map_var_make = make(map[string]string)
	fmt.Println(map_var_make == nil) // => false

	// make shorthand initialized map
	map_var_make_shorthand := make(map[string]string)
	fmt.Println(map_var_make_shorthand == nil) // => false

	// shorthand composite literal
	map_shorthand_composite := map[string]string{}
	fmt.Println(map_shorthand_composite == nil) // => false

	map_shorthand_composite_with_inits := map[string]int {
		"foo":1,
		"bar":1, // do not forget the comma
	}
	fmt.Printf("%v\n", map_shorthand_composite_with_inits) // map[foo:1 bar:1]
	map_shorthand_composite_with_inits["foo"] ++
	fmt.Printf("%v\n", map_shorthand_composite_with_inits) // map[foo:, bar:1]

}



