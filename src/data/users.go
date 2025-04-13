package data

type User struct {
	Fullname string
	Id       string
	Password string
	Wallet   float64
}

var Users = make(map[string]User)