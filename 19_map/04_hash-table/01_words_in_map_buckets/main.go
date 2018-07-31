//CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/tree/master/19_map/14_hash-table/03_words-in-buckets/02_map-bucket
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

const num_buckets = 12

func main() {
	// get the book
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the book
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create map with a key of int
	// and a value of another map
	// with a key of string, which will be the word
	// and a value of int, which will be the number of ti,es the words occur
	buckets := make(map[int]map[string]int)

	// initialize the map ti hold words
	for i := 0; i < num_buckets; i++ {
		buckets[i] = make(map[string]int)
	}

	// loop through the words
	for scanner.Scan() {
		word := scanner.Text()
		index := hashBucket(word, 12)
		buckets[index][word]++
	}

	// print the words map
	for i := 0; i < num_buckets; i++ {
		fmt.Printf("%d: ", len(buckets[i]))
		fmt.Println(buckets[i])
	}

}

func hashBucket(word string, num_buckets int) int {
	var sum int
	for _, char := range word {
		sum += int(char)
	}
	return sum % num_buckets
}
