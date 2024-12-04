package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Api struct {
	year string
	day  string
}

func NewApi(day string, year string) *Api {
	return &Api{
		day:  day,
		year: year,
	}
}

func (api *Api) LoadData() {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", api.year, api.day), nil)

	if err != nil {
		log.Fatalln("Could not create request", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Cookie", os.Getenv("AOC_TOKEN"))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatalln(fmt.Errorf("could not get day %s %s %w", api.day, api.year, err))
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(fmt.Errorf("could not read response body %w", err))
	}

	err = os.WriteFile(fmt.Sprintf("./input/cache/data/%s/%s", api.year, api.day), body, 0644)

	if err != nil {
		log.Fatalln(fmt.Errorf("could not write to cache %w", err))
	}
}
