package main

import (
	"fmt"

	"github.com/heart/play-with-go/request"
)

func main() {
	mems, e := NewBNKService().getBNKMembers(request.NewRequester())

	if e == nil {
		for _, m := range mems {
			fmt.Println(m.ThaiFirstName, m.ThaiLastName)
		}
	} else {
		fmt.Println(e)
	}
}
