package day3

import (
	"advent-of-code-2024/days/common/utils"
	"advent-of-code-2024/input"
	"regexp"
	"slices"
	"strings"
)

type Day struct {
	data      []string
	solution1 int
	solution2 int
}

func NewDay() *Day {
	i := input.NewInput("3", "2024")

	data := i.GetData()

	return &Day{
		data:      data,
		solution1: 0,
		solution2: 0,
	}
}

func (d *Day) Run() (int, int) {
	data := strings.Join(d.data, "")

	mulR := regexp.MustCompile(`(?:mul\((\d+),(\d+)\)){1}`)
	mulMatches := mulR.FindAllStringSubmatch(data, -1)
	mulIndexes := mulR.FindAllStringIndex(data, -1)

	doR := regexp.MustCompile(`(?:do\(\)){1}`)
	doIndexes := doR.FindAllStringIndex(data, -1)

	dontR := regexp.MustCompile(`(?:don't\(\)){1}`)
	dontIndexes := dontR.FindAllStringIndex(data, -1)

	for _, m := range mulMatches {
		d.solution1 += utils.ConvertStringToInt(m[1]) * utils.ConvertStringToInt(m[2])
	}

	enableMul := true

	for index, _ := range data {
		doExistsForIndex := slices.ContainsFunc(doIndexes, func(do []int) bool { return do[0] == index })

		if doExistsForIndex {
			enableMul = true
		}

		dontExistsForIndex := slices.ContainsFunc(dontIndexes, func(dont []int) bool { return dont[0] == index })

		if dontExistsForIndex {
			enableMul = false
		}

		mulIndex := slices.IndexFunc(mulIndexes, func(mul []int) bool { return mul[0] == index })

		if enableMul && mulIndex != -1 {
			d.solution2 += utils.ConvertStringToInt(mulMatches[mulIndex][1]) * utils.ConvertStringToInt(mulMatches[mulIndex][2])
		}
	}

	return d.solution1, d.solution2
}
