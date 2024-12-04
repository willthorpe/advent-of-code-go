package common

import (
	"log"
	"math"
	"strconv"
)

func Abs(x int, y int) int {
	return int(math.Abs(float64(x - y)))
}

func ConvertStringToInt(s string) int {
	intValue, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return intValue
}
