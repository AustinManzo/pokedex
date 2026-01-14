package main

import (
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {

	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {

	}
	if err != nil {

	}
	fmt.Printf("%s", body)
	return nil
}
