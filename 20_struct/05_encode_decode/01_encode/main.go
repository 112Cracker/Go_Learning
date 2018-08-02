package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First string
	Last string
	Age int
	secret_id string
}

func main() {
	p := person{"James", "Bond", 20,"007"}
	json.NewEncoder(os.Stdout).Encode(p)
}