package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNewApplication(t *testing.T) {
	textCfg := "Foo north=Bar"
	config := bufio.NewScanner(strings.NewReader(textCfg))
	cityMap := NewCityMap(config)

	foo := cityMap.cities["Foo"]
	bar := cityMap.cities["Bar"]
	strat := &mockStrategy{
		locations: map[int]*City{
			0: foo,
			1: bar,
		},
	}

	app := NewApplication(cityMap, strat, 2, 0)
	expectedPositions := map[*City][]int{
		foo: {0},
		bar: {1},
	}

	if !reflect.DeepEqual(app.aliensPositions, expectedPositions) {
		t.Errorf("expected positions %#v, got %#v", expectedPositions, app.aliensPositions)
	}
}

func TestApplication_Run_CityDestroyed(t *testing.T) {
	textCfg := "Foo east=Bar\nBar east=Baz"
	config := bufio.NewScanner(strings.NewReader(textCfg))
	cityMap := NewCityMap(config)

	foo := cityMap.cities["Foo"]
	bar := cityMap.cities["Bar"]

	// make 2 aliens meet at the most eastern city (Baz) and destroy it
	strat := &mockStrategy{
		locations: map[int]*City{
			0: foo,
			1: bar,
		},
		direction: direction("east"),
	}

	app := NewApplication(cityMap, strat, 2, 2)
	app.Run()

	// Baz is expected to be destroyed
	delete(bar.neighbourNames, "Baz")
	delete(bar.neighbours, "east")

	expectedCities := map[string]*City{
		"Foo": foo,
		"Bar": bar,
	}

	if len(app.cityMap.citiesList) != 2 {
		t.Errorf("citiesList size expected %d, got %d", 2, len(app.cityMap.citiesList))
	}

	if len(app.aliensPositions) != 0 {
		t.Errorf("aliensPositions size expected %d, got %d", 0, len(app.aliensPositions))
	}

	if !reflect.DeepEqual(expectedCities, app.cityMap.cities) {
		t.Errorf("cities expected %#v, got %#v", expectedCities, app.cityMap.cities)
	}
}

func TestApplication_Run_NoneIsDestroyed(t *testing.T) {
	textCfg := "Foo east=Bar\nZoo east=Gaz"
	config := bufio.NewScanner(strings.NewReader(textCfg))
	cityMap := NewCityMap(config)

	foo := cityMap.cities["Foo"]
	zoo := cityMap.cities["Zoo"]

	// make 2 aliens meet at the most eastern city (Baz) and destroy it
	strat := &mockStrategy{
		locations: map[int]*City{
			0: foo,
			1: zoo,
		},
		direction: direction("east"),
	}

	app := NewApplication(cityMap, strat, 2, 2)
	app.Run()

	expectedCities := map[string]*City{
		"Foo": cityMap.cities["Foo"],
		"Zoo": cityMap.cities["Zoo"],
		"Bar": cityMap.cities["Bar"],
		"Gaz": cityMap.cities["Gaz"],
	}

	if len(app.cityMap.citiesList) != 4 {
		t.Errorf("citiesList size expected %d, got %d", 4, len(app.cityMap.citiesList))
	}

	if len(app.aliensPositions) != 2 {
		t.Errorf("aliensPositions size expected %d, got %d", 2, len(app.aliensPositions))
	}

	if !reflect.DeepEqual(expectedCities, app.cityMap.cities) {
		t.Errorf("cities expected %#v, got %#v", expectedCities, app.cityMap.cities)
	}
}
