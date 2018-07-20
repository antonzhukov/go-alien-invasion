# Description

This is the simple game called Alien Invasion which is implemented in Go.

## Assumptions

1. If in the file there's a city A which has a road to city B, 
I assume that the city B also has the road to city A. 
2. There's only 1 road from city A to city B. For this I implemented neighbours uniqueness map.
3. Instruction states that if two aliens end up in the same city, the city is being destroyed.
I assume that if there are more than 1 alien in the city, they destroy it and die themselves.

## Trade-Offs

1. There's a slice of cities in the CityMap. It is required only for aliens to disembark
at a random city from this list. The bad thing about this slice is that we need to
maintain the actual list of cities, which causes O(n) run everytime we destroy a city.
Honestly, the other way to randomly pick a city is start a `for` cycle and it will give 
us the random city. But with `rand` package I wanted to make it more explicit.