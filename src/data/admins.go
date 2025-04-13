package data

import (
	"snapp/config"
	"snapp/loggers"
)

type Admin struct {
	Fullname string
	Id       string
	Password string
}

var logger = loggers.NewLogger(config.GetConfig())
var Admins = map[string]Admin{"arash2007mahdavi": {Fullname: "arash mahdavi", Id: "arash2007mahdavi", Password: "@rash2007"}}

func AdminCheck(id string, password string) bool {
	for k, v := range Admins {
		if id == k {
			if v.Password == password {
				return true
			} else {
				return false
			}
		}
	}
	return false
}