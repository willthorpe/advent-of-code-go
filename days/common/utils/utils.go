package utils

import (
	"log"
	"math"
	"regexp"
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

func FindPositionOfRegexInString(data string, regex string) (int, int) {
	r := regexp.MustCompile(regex)
	i := r.FindStringIndex(data)

	if len(i) > 0 {
		return i[0], i[1]
	}

	return -1, -1
}
