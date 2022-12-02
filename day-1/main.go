package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	checkErr(err)

	var totals []int
	currentIndex := 0

	splits := bytes.Split(input, []byte("\n"))

	// Can add an additional step to avoid the 0 out indexes but
	// then I don't really care about it right now

	totals = make([]int, len(splits))

	largest := 0

	for _, n := range splits {

		if len(bytes.TrimSpace(n)) == 0 {

			if largest < totals[currentIndex] {
				largest = totals[currentIndex]
			}

			currentIndex += 1
			continue
		}

		num, err := strconv.Atoi(string(n))
		checkErr(err)

		totals[currentIndex] += num
	}

	log.Println(largest)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
