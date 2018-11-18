package main

import (
	"fmt"

	"github.com/heart/play-with-go/members"
)

func main() {
	str := ""
	str = fmt.Sprintf("%s, %s!", "Hello", "World")
	fmt.Println(str)

	m := &members.Member{
		ID:      "0",
		Name:    "Narongrit Kanhanoi",
		Email:   "nkheart4@gmail.com",
		Enabled: true,
	}
	str = fmt.Sprintf("member : %v", m)
	fmt.Println(str)

	m2 := m.Copy()
	m2.Name = "Heart"
	str = fmt.Sprintf("m = %v\nm2 = %v", m, m2)
	fmt.Println(str)
}
