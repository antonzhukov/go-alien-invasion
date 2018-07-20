package main

type Application struct {
	cityMap   *CityMap
	aliensNum int
}

func NewApplication(cityMap *CityMap, aliensNum int) *Application {
	return &Application{
		cityMap:   cityMap,
		aliensNum: aliensNum,
	}
}
