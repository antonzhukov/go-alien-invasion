package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const maxMovementsNum = 10000

func main() {
	// init city config
	file, err := os.Open("config.txt")
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
