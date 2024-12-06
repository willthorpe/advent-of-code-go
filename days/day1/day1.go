package day1

import (
	"advent-of-code-2024/days/common/utils"
	"advent-of-code-2024/input"
	"slices"
	"strings"
)

type Day struct {
	data      []string
	solution1 int
	solution2 int
}

func NewDay() Day {
	i := input.NewInput("1", "2024")

	data := i.GetData()

	return Day{
		data:      data,
		solution1: 0,
		solution2: 0,
	}
}

func (d *Day) Run() (int, int) {
	leftList, rightList := d.generateSortedLists()

	for index, value := range leftList {
		d.solution1 += utils.Abs(value, rightList[index])
	}

	for _, value := range leftList {
		occurrencesInRightList := 0

		for _, rightValue := range rightList {
			if value == rightValue {
				occurrencesInRightList++
			}
		}

		d.solution2 += value * occurrencesInRightList
	}

	return d.solution1, d.solution2
}

func (d *Day) generateSortedLists() ([]int, []int) {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, line := range d.data {
		splitLine := strings.SplitAfter(line, "   ")
		leftValue := utils.ConvertStringToInt(strings.TrimSpace(splitLine[0]))
		rightValue := utils.ConvertStringToInt(strings.TrimSpace(splitLine[1]))

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList
}
