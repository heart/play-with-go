package main

import (
	"fmt"

	"github.com/heart/play-with-go/request"
	"github.com/heart/play-with-go/services"
)

func main() {
	bnkService := services.NewBNKService()
	mems, e := bnkService.GetBNKMembers(request.NewRequester())

	if e == nil {
		for _, m := range mems {
			fmt.Println(m.ThaiFirstName, m.ThaiLastName)
		}
	} else {
		fmt.Println(e)
	}
}
