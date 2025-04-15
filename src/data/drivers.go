package data

type Driver struct {
	Fullname string
	Id string
	Password string
	Wallet float64
}

var Drivers = make(map[string]Driver)

func AddToDataBase(new Driver) {
	Drivers[new.Id] = new
}