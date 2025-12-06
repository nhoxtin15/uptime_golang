package controller

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
}

func logout(w http.ResponseWriter, r *http.Request) {
}

func forgotPassword(w http.ResponseWriter, r *http.Request) {
}

func resetPassword(w http.ResponseWriter, r *http.Request) {
}

func verifyEmail(w http.ResponseWriter, r *http.Request) {

}


func authentication_controller(){
	http.HandleFunc("POST /login", login)
	http.HandleFunc("POST /register", register)
	http.HandleFunc("POST /logout", logout)
	http.HandleFunc("POST /forgotPassword", forgotPassword)
	http.HandleFunc("POST /resetPassword", resetPassword)
	http.HandleFunc("POST /verifyEmail", verifyEmail)
}