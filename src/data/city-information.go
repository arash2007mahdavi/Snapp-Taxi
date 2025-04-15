package data

import (
	"math"
)

type City string

const (
	Tehran     City = "Tehran"
	Tehran_lat int  = 10
	Tehran_lon int  = 17

	Tabriz     City = "Tabriz"
	Tabriz_lat int  = 3
	Tabriz_lon int  = 19

	Esfehan     City = "Esfehan"
	Esfehan_lat int  = 10
	Esfehan_lon int  = 10

	Shiraz     City = "Shiraz"
	Shiraz_lat int  = 17
	Shiraz_lon int  = 8

	Mashhad     City = "Mashhad"
	Mashahd_lat int  = 19
	Mashahd_lon int  = 18

	Boshehr     City = "Boshehr"
	Boshehr_lat int  = 5
	Boshehr_lon int  = 1

	Rasht     City = "Rasht"
	Rasht_lat int  = 8
	Rasht_lon int  = 18

	Zahedan     City = "Zahedan"
	Zahedan_lat int  = 18
	Zahedan_lon int  = 3

	Kerman     City = "Kerman"
	Kerman_lat int  = 6
	Kerman_lon int  = 7

	Ardabil     City = "Ardabil"
	Ardabil_lat int  = 4
	Ardabil_lon int  = 20

	Karaj     City = "Karaj"
	Karaj_lat int  = 9
	Karaj_lon int  = 17
)

var Citys = []City{Tehran,Tabriz,Esfehan,Shiraz,Mashhad,Boshehr,Rasht,Zahedan,Kerman,Ardabil,Karaj}

var Lats = map[City]int{
	Tehran:  Tehran_lat,
	Tabriz:  Tabriz_lat,
	Esfehan: Esfehan_lat,
	Shiraz:  Shiraz_lat,
	Mashhad: Mashahd_lat,
	Boshehr: Boshehr_lat,
	Rasht:   Rasht_lat,
	Zahedan: Zahedan_lat,
	Kerman:  Kerman_lat,
	Ardabil: Ardabil_lat,
	Karaj:   Karaj_lat,
}

var Lons = map[City]int{
	Tehran:  Tehran_lon,
	Tabriz:  Tabriz_lon,
	Esfehan: Esfehan_lon,
	Shiraz:  Shiraz_lon,
	Mashhad: Mashahd_lon,
	Boshehr: Boshehr_lon,
	Rasht:   Rasht_lon,
	Zahedan: Zahedan_lon,
	Kerman:  Kerman_lon,
	Ardabil: Ardabil_lon,
	Karaj:   Karaj_lon,
}

func GetDistance(city1 City, city2 City) (distance float64) {
	x := Lats[city1] - Lats[city2]
	y := Lons[city1] - Lons[city2]
	distance = math.Sqrt(float64((x * x) + (y * y)))
	distance = math.Round(distance*100) / 100
	return
}

func GetPrice(distance float64) (price float64) {
	price = distance * 50
	return
}

func ValidateCity(name string) bool {
	for _, city := range Citys {
		if city == City(name) {
			return true
		}
	}
	return false
}