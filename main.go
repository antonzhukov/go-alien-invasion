package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// init city config
	file, err := os.Open("data/config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cityMap := NewCityMap(scanner)

	for _, c := range cityMap.cities {
		fmt.Printf("%#v\n", c)
	}
}
