package main

import (
	"bufio"
	"log"
	"strings"
)

type direction string

const (
	directionNorth direction = "north"
	directionEast  direction = "east"
	directionSouth direction = "south"
	directionWest  direction = "west"
)

// CityMap represents a world map
type CityMap struct {
	cities map[string]*City
}

// City describes a single city and its neighbours
type City struct {
	name           string
	neighbours     map[direction]*City // neighbours routes
	neighbourNames map[string]bool     // neighbours uniqueness map
}

// NewCityMap creates an entire map of cities and the roads between them
func NewCityMap(config *bufio.Scanner) *CityMap {
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
			// validate direction
			dir := direction(d[0])
			if !(dir == directionNorth || dir == directionEast || dir == directionSouth || dir == directionWest) {
				continue
			}
			rawMap[name][dir] = d[1]
			citiesList[d[1]] = true
		}
	}

	// add all mentioned cities to the list
	cityMap := &CityMap{cities: make(map[string]*City)}
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

// linkCities adds a road from one city to another in the given direction
func linkCities(c *CityMap, city, neighbour string, dir direction) {
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

// destroyCity removes the city from the map along with all the roads leading in and out of it
func (cm *CityMap) destroyCity(name string) {
	city, ok := cm.cities[name]
	if !ok {
		log.Printf("DestroyCity failed: city '%s' does not exist", name)
		return
	}

	// destroy all roads leading to the current city
	for dir, neighbour := range city.neighbours {
		delete(neighbour.neighbours, getOppositeDirection(dir))
		delete(neighbour.neighbourNames, name)
	}

	// delete the city itself from the map
	delete(cm.cities, name)
}
