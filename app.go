package main

import (
	"bufio"
	"strings"
)

type direction string

const (
	directionNorth direction = "north"
	directionEast  direction = "east"
	directionSouth direction = "south"
	directionWest  direction = "west"
)

type CityMap struct {
	cities map[string]*City
	list   *City
}

type City struct {
	name           string
	neighbours     map[direction]*City // neighbours routes
	neighbourNames map[string]bool     // neighbours uniqueness map
}

func NewCityMap(config *bufio.Scanner) CityMap {
	rawMap := make(map[string]map[direction]string)
	citiesList := make(map[string]bool)
	for config.Scan() {
		s := strings.Split(config.Text(), " ")
		if len(s) < 1 {
			continue
		}
		name := s[0]
		citiesList[name] = true

		rawMap[name] = make(map[direction]string)
		for i := 1; i < len(s); i++ {
			d := strings.Split(s[i], "=")
			if len(d) != 2 {
				continue
			}
			dir := direction(d[0])
			if !(dir == directionNorth || dir == directionEast || dir == directionSouth || dir == directionWest) {
				continue
			}
			rawMap[name][dir] = d[1]
			citiesList[d[1]] = true
		}
	}

	// add all mentioned cities to the list
	cityMap := CityMap{cities: make(map[string]*City)}
	for city := range citiesList {
		cityMap.cities[city] = &City{
			name:           city,
			neighbours:     make(map[direction]*City),
			neighbourNames: make(map[string]bool),
		}
	}

	// go through each city and build the directions map
	for name, cityNeighbours := range rawMap {
		for dir, neighbour := range cityNeighbours {
			// cross-link two cities
			linkCities(cityMap, name, neighbour, dir)
			linkCities(cityMap, neighbour, name, getOppositeDirection(dir))
		}
	}

	return cityMap
}

func linkCities(c CityMap, city, neighbour string, dir direction) {
	if _, ok := c.cities[city].neighbours[dir]; !ok {
		// make sure one city has only 1 link to the other one
		if !c.cities[city].neighbourNames[neighbour] {
			c.cities[city].neighbours[dir] = c.cities[neighbour]
			c.cities[city].neighbourNames[neighbour] = true
		}
	}
}

func getOppositeDirection(d direction) direction {
	switch d {
	case directionWest:
		return directionEast
	case directionEast:
		return directionWest
	case directionNorth:
		return directionSouth
	case directionSouth:
		return directionNorth

	}
	return directionNorth
}
