package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const maxMovementsNum = 10000

func main() {
	// parse flags
	aliens := flag.Int("aliens", 1, "number of aliens to invade the map")
	flag.Parse()

	// init city config
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// init dependencies
	scanner := bufio.NewScanner(file)
	cityMap := NewCityMap(scanner)

	// start the app
	app := NewApplication(cityMap, NewRandomStrategy(), *aliens, maxMovementsNum)
	app.Run()

	// export the results
	f, err := os.Create("export.txt")
	if err != nil {
		fmt.Errorf("failed to export map: %s", err.Error())
		return
	}
	defer f.Close()

	mapString := app.cityMap.ExportCityMap()
	n, err := f.WriteString(mapString)
	fmt.Printf("wrote %d bytes\n", n)

}
