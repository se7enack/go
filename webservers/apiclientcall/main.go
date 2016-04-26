package main

import (
	"fmt"
	"github.com/bndr/gopencils"
)

type respStruct struct {
	Id        int
	Name      string
	Completed bool
	Due       string
}

func main() {
	api := gopencils.Api("http://localhost:9000")
	// Users Resource
	jobs := api.Res("jobs")

	id := []string{"1"}

	for _, username := range id {
		// Create a new pointer to response struct
		r := new(respStruct)
		// Get user with id i into the newly created response struct
		_, err := jobs.Id(username, r).Get()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
	}
}
