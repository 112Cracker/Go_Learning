// CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/tree/master/05_blank-identifier/02_http-get_example
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

/*
Blank identifier
	- tell the compiler you aren't using something
		- you are obliged to use everything in your code
	- Get
		- throwing away an error
 */