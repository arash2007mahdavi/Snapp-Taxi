package data

type User struct {
	fullname string
	id       string
	password string
	wallet   float64
}

type SimpleUser struct {
	Fullname string `json:"fullname" binding:"required,min=5,max=25"`
	ID       string `json:"id" binding:"required,id"`
	Password string `json:"password" binding:"required,password"`
}

var users = make(map[string]User)

func AddUser(user SimpleUser) bool {
	for k := range users {
		if k == user.ID {
			return false
		}
	}
	new := User{
		fullname: user.Fullname,
		id:       user.ID,
		password: user.Password,
	}
	new.wallet = 0
	users[new.id] = new
	return true
}
