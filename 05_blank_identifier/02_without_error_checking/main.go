// CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/tree/master/05_blank-identifier/02_http

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com/") // _ is a blank identifier
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
