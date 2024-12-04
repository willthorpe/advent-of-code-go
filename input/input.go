package input

import (
	"advent-of-code-2024/input/api"
	"advent-of-code-2024/input/cache"
	"log"
)

type Input struct {
	day   string
	year  string
	api   *api.Api
	cache *cache.Cache
}

func NewInput(day string, year string) *Input {
	return &Input{
		api:   api.NewApi(day, year),
		cache: cache.NewCache(day, year),
	}
}

func (i *Input) GetData() []string {
	cacheExists := i.cache.CacheExists()

	if cacheExists {
		data, err := i.cache.GetData()

		if err != nil {
			log.Println(err)
		}

		return data
	}

	i.api.LoadData()

	data, err := i.cache.GetData()

	if err != nil {
		log.Println(err)
	}
	
	return data
}
