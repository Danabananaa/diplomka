package model

type AuthSerivice interface {
	SignUp()
	LogOut()
	LogIn()
	Refresh()
}

type Analysis interface{}
