package day2

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

func NewDay() *Day {
	i := input.NewInput("2", "2024")

	data := i.GetData()

	return &Day{
		data:      data,
		solution1: 0,
		solution2: 0,
	}
}

func (d *Day) Run() (int, int) {
	for _, line := range d.data {
		values := strings.Split(line, " ")
		intValues := make([]int, 0)

		for _, value := range values {
			intValues = append(intValues, utils.ConvertStringToInt(value))
		}

		gaps, ordering := d.findGapsAndOrdering(intValues)
		errorCount := d.findErrorCount(gaps, ordering)

		if errorCount == 0 {
			d.solution1 += 1
		}

		if errorCount > 0 {
			for index, _ := range intValues {
				intValuesTest := make([]int, len(values))
				copy(intValuesTest, intValues)
				intValuesTest = slices.Delete(intValuesTest, index, index+1)

				gaps, ordering = d.findGapsAndOrdering(intValuesTest)
				errorCountTest := d.findErrorCount(gaps, ordering)

				if errorCountTest == 0 {
					d.solution2 += 1

					break
				}
			}
		}
	}
	d.solution2 += d.solution1

	return d.solution1, d.solution2
}

func (d *Day) findGapsAndOrdering(intValues []int) ([]int, []string) {
	gaps := make([]int, 0)
	ordering := make([]string, 0)

	for index, value := range intValues {
		if index > 0 {
			gaps = append(gaps, utils.Abs(value, intValues[index-1]))

			if value > intValues[index-1] {
				ordering = append(ordering, "ASC")
			} else if value < intValues[index-1] {
				ordering = append(ordering, "DESC")
			} else {
				ordering = append(ordering, "SAME")
			}
		}
	}
	return gaps, ordering
}

func (d *Day) findErrorCount(gaps []int, ordering []string) int {
	countGapErrors := 0
	for _, gap := range gaps {
		if gap == 0 || gap > 3 {
			countGapErrors++
		}
	}

	countOrderingErrors := 0
	firstOrdering := ordering[0]
	for _, value := range ordering {
		if value != firstOrdering {
			countOrderingErrors++
		}
	}

	errorCount := countGapErrors + countOrderingErrors

	return errorCount
}
