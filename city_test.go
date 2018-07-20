package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestNewCityMap(t *testing.T) {
	textCfg := "Foo north=Bar"
	config := bufio.NewScanner(strings.NewReader(textCfg))
	cityMap := NewCityMap(config)

	if len(cityMap.cities) != 2 {
		t.Errorf("expected %d cities, got %d", 2, len(cityMap.cities))
	}

	if !reflect.DeepEqual(cityMap.cities["Foo"].neighbourNames, map[string]bool{"Bar": true}) {
		t.Errorf("expected neighbour Bar")
	}

	if cityMap.cities["Foo"].neighbours[directionNorth].name != "Bar" {
		t.Errorf("expected neighbour Bar, got %s", cityMap.cities["Foo"].neighbours[directionNorth].name)
	}

	if !reflect.DeepEqual(cityMap.cities["Bar"].neighbourNames, map[string]bool{"Foo": true}) {
		t.Errorf("expected neighbour Bar")
	}

	if cityMap.cities["Bar"].neighbours[directionSouth].name != "Foo" {
		t.Errorf("expected neighbour Foo, got %s", cityMap.cities["Bar"].neighbours[directionSouth].name)
	}
}

func TestCityMap_destroyCity(t *testing.T) {
	textCfg := "Foo north=Bar"
	config := bufio.NewScanner(strings.NewReader(textCfg))
	cityMap := NewCityMap(config)
	cityToDestroy := cityMap.cities["Foo"]
	cityMap.destroyCity(cityToDestroy)

	leftover := &City{
		name:           "Bar",
		neighbours:     map[direction]*City{},
		neighbourNames: map[string]bool{},
	}

	if len(cityMap.cities) != 1 {
		t.Errorf("expected %d cities, got %d", 1, len(cityMap.cities))
	}

	if !reflect.DeepEqual(leftover, cityMap.cities["Bar"]) {
		t.Errorf("expected leftover %#v, got %#v", leftover, cityMap.cities["Bar"])
	}
}

func TestCityMap_ExportCityMap(t *testing.T) {
	type fields struct {
		citiesList []*City
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{}, ""},
		{
			"valid",
			fields{
				citiesList: []*City{
					{name: "Foo", neighbours: map[direction]*City{directionWest: {name: "Bar"}}},
					{name: "Baz", neighbours: map[direction]*City{directionEast: {name: "Faz"}}},
				},
			},
			"Foo west=Bar\nBaz east=Faz\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &CityMap{
				citiesList: tt.fields.citiesList,
			}
			if got := cm.ExportCityMap(); got != tt.want {
				t.Errorf("CityMap.ExportCityMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
