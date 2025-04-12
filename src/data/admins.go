package data

import (
	"fmt"
	"snapp/config"
	"snapp/loggers"
)

type admin struct {
	fullname string
	id       string
	password string
}

var logger = loggers.NewLogger(config.GetConfig())
var admins = map[string]admin{"arash2007mahdavi": {fullname: "arash mahdavi", id: "arash2007mahdavi", password: "@rash2007"}}

func AdminCheck(id string, password string) bool {
	for k, v := range admins {
		if id == k {
			if v.password == password {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func AddAdmin(fullname string, id string, password string) error {
	for k := range admins {
		if k == id {
			logger.Error(loggers.Admin, loggers.Add, "invalid ID", nil)
			return fmt.Errorf("there is an admin with this ID")
		}
	}
	new := admin{
		fullname: fullname,
		id: id,
		password: password,
	}
	admins[id] = new
	return nil
}