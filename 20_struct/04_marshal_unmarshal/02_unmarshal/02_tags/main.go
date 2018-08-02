package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p person

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	json.Unmarshal(bs, &p)

	fmt.Printf("%s %s is %d years old.", p.First, p.Last, p.Age)
}
