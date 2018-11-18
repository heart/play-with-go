package main

import (
	"fmt"

	"github.com/heart/play-with-go/members"
)

func login(member members.ILogin) bool {
	uname := member.Email()
	pass := member.Password()
	if uname == "nkheart4@gmail.com" && pass == "password" {
		return true
	}
	return false
}

func main() {
	m := &members.Member{
		ID:           "0",
		Name:         "Narongrit Kanhanoi",
		UserEmail:    "nkheart4@gmail.com",
		UserPassword: "password",
		Enabled:      true,
	}
	pass := login(m)
	if pass {
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Fail")
	}
}
