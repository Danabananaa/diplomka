package model

import "net/http"

type AuthHandlers interface {
	// LogIn Exiting the session
	LogIn(http.ResponseWriter, *http.Request)
	// LogOut Entering the session
	LogOut(http.ResponseWriter, *http.Request)
	// Refresh Updating a token
	Refresh(http.ResponseWriter, *http.Request)
	// SigUp User registration
	SignUp(http.ResponseWriter, *http.Request)
}
