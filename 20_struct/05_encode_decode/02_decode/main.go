package main

import (
	"strings"
	"encoding/json"
	"fmt"
	)

type person struct {
	First string
	Last string
	Age int
	secret_id string
}

func main() {
	var p person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p)

	fmt.Printf("%s %s is %d years old.", p.First, p.Last, p.Age)
}