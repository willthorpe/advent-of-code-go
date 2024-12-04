package main

import (
	"advent-of-code-2024/days/day1"
	"advent-of-code-2024/days/day2"
	"advent-of-code-2024/days/day3"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d1 := day1.NewDay()
	d2 := day2.NewDay()
	d3 := day3.NewDay()

	log.Println(d1.Run())
	log.Println(d2.Run())
	log.Println(d3.Run())
}
