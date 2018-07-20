package main

import (
	"math/rand"
	"time"
)

type alienStrategy interface {
	disembark(cities []*City, alienID int) *City
	move() direction
}

type randomStrategy struct{}

func NewRandomStrategy() *randomStrategy {
	rand.Seed(time.Now().UTC().UnixNano())
	return &randomStrategy{}
}

func (s *randomStrategy) disembark(cities []*City, alienID int) *City {
	return cities[rand.Intn(len(cities))]
}

func (s *randomStrategy) move() direction {
	return directions[rand.Intn(len(directions))]
}

type mockStrategy struct {
	direction direction
	locations map[int]*City
}

func (s *mockStrategy) disembark(_ []*City, alienID int) *City {
	return s.locations[alienID]
}

func (s *mockStrategy) move() direction {
	return s.direction
}
