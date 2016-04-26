package main

import (
	"fmt"
	"github.com/bndr/gopencils"
)
#SJB
type respStruct struct {
	Name  string
	Completed bool
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