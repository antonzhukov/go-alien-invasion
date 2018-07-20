package main

import (
	"log"
	"strconv"
)

type Application struct {
	cityMap         *CityMap
	aliensPositions map[*City][]int // map of aliens' locations
	aliensStrategy  alienStrategy
	maxMovementsNum int
}

// NewApplication instantiates new aliens invasion app, setting the initial aliens' location
func NewApplication(cityMap *CityMap, aliensStrategy alienStrategy, aliensNum, maxMovementsNum int) *Application {
	a := &Application{
		cityMap:         cityMap,
		aliensPositions: make(map[*City][]int),
		aliensStrategy:  aliensStrategy,
		maxMovementsNum: maxMovementsNum,
	}
	a.disembarkAliens(aliensNum)

	return a
}

func (a *Application) disembarkAliens(aliensNum int) {
	for i := 0; i < aliensNum; i++ {
		city := a.aliensStrategy.disembark(a.cityMap.citiesList, i)
		a.aliensPositions[city] = append(a.aliensPositions[city], i)
	}
}

// Run starts aliens' run through the city
func (a *Application) Run() {
	// try to destroy cities in the beginning if we are lucky to have multiple aliens disembarked at one city
	a.startDestruction()

	for i := 0; i < a.maxMovementsNum; i++ {
		// move aliens
		a.aliensPositions = a.moveAliens()

		// destroy all cities and aliens
		a.startDestruction()
	}
}

func (a *Application) startDestruction() {
	for city, aliens := range a.aliensPositions {
		if len(aliens) > 1 {
			aliensStr := "alien " + strconv.Itoa(aliens[0])
			for j := 1; j < len(aliens); j++ {
				aliensStr += " and alien " + strconv.Itoa(aliens[j])
			}
			log.Printf("%s has been destroyed by %s", city.name, aliensStr)
			a.cityMap.destroyCity(city)
			delete(a.aliensPositions, city)
		}
	}
}

func (a *Application) moveAliens() map[*City][]int {
	newPositions := make(map[*City][]int)
	for city, aliens := range a.aliensPositions {
		for _, alienID := range aliens {
			nextDirection := a.aliensStrategy.move()
			// if there's a city in the given direction, alien moves there
			if nextCity, ok := city.neighbours[nextDirection]; ok {
				newPositions[nextCity] = append(newPositions[nextCity], alienID)
			} else {
				// otherwise he stays at the same location
				newPositions[city] = append(newPositions[city], alienID)
			}
		}
	}
	return newPositions
}
