package controllers

import (
	"net/http"
)

// RegisterControllers is a frontend
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
