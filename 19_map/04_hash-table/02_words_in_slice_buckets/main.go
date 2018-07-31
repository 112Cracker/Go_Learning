//CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/tree/master/19_map/14_hash-table/03_words-in-buckets/01_slice-bucket
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

const num_buckets = 12

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanner
	scanner.Split(bufio.ScanWords)
	// create 2D slice of string to hold words
	buckets := make([][]string, num_buckets)

	// loop through the words
	for scanner.Scan() {
		word := scanner.Text()
		index := hashBucket(word, num_buckets)
		buckets[index] = append(buckets[index], word)
	}
	// print the length and elements of each bucket
	for i := 0; i < num_buckets; i++ {
		fmt.Printf("%d: ", len(buckets[i]))
		fmt.Println(buckets[i])
	}
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, char := range word {
		sum += int(char)
	}
	return sum % buckets
}
