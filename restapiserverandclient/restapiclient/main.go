package main

import (
	"fmt"
	"github.com/bndr/gopencils"
)

type fieldStructure struct {
	Id        int
	Name      string
	Completed bool
	Due       string
}

func main() {
	api := gopencils.Api("http://localhost:9000")
	jobs := api.Res("jobs")
	id := []string{"1"}

	for _, x := range id {
		r := new(fieldStructure)
		_, err := jobs.Id(x, r).Get()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
	}
}
