package data

type LineTravel struct {
	CustomerID string
	Genesis City
	Target City
	Distance float64
	Price float64
}

var counterLineTravels = 1
var LineTravels = make(map[int]LineTravel)

func AddToLineTravels(new LineTravel) {
	LineTravels[counterLineTravels] = new
	counterLineTravels++
}