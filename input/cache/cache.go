package cache

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cache struct {
	year string
	day  string
}

func NewCache(day string, year string) *Cache {
	return &Cache{
		day:  day,
		year: year,
	}
}

func (c *Cache) CacheExists() bool {
	exists, err := os.Stat(fmt.Sprintf("./input/cache/data/%s/%s", c.year, c.day))

	if err != nil {
		log.Println(err)
		return false
	}

	if exists != nil {
		return true
	}

	return false
}

func (c *Cache) GetData() ([]string, error) {
	var data []string

	file, err := os.Open(fmt.Sprintf("./input/cache/data/%s/%s", c.year, c.day))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return data, nil
}
