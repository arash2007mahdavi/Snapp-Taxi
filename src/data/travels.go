package data

type Travel struct {
	TravelID string
	CustomerID string
	DriverID string
	Genesis City
	Target City
	Distance float64
	Price float64
}

var Travels = make(map[string]Travel)
var counterTravels = 1