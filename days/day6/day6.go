package day6

import (
	"advent-of-code-2024/days/common/matrix"
	"advent-of-code-2024/days/common/utils"
	"advent-of-code-2024/input"
	"slices"
	"strconv"
	"strings"
)

type Day struct {
	data      []string
	solution1 int
	solution2 int
}

func NewDay() *Day {
	i := input.NewInput("6", "2024")

	data := i.GetData()

	return &Day{
		data:      data,
		solution1: 0,
		solution2: 0,
	}
}

func (d *Day) Run() (int, int) {
	data := matrix.Create(d.data)
	guardStartingRowPosition := 0
	guardStartingColumnPosition := 0

	for rowIndex, row := range data {
		indexOfGuard := slices.IndexFunc(row, func(i string) bool { return i == "^" })
		if indexOfGuard != -1 {
			guardStartingRowPosition = rowIndex
			guardStartingColumnPosition = indexOfGuard

			break
		}
	}

	coordinatesVisited := d.findUniqueCoordinatesVisited(data, guardStartingRowPosition, guardStartingColumnPosition)
	d.solution1 = len(coordinatesVisited)

	for _, coordinate := range coordinatesVisited {
		splitCoordinate := strings.Split(coordinate, ",")

		rowIndex := utils.ConvertStringToInt(splitCoordinate[0])
		colIndex := utils.ConvertStringToInt(splitCoordinate[1])

		testData := matrix.Create(d.data)
		if rowIndex == guardStartingRowPosition && colIndex == guardStartingColumnPosition {
			continue
		}

		testData[rowIndex][colIndex] = "O"
		d.findUniqueCoordinatesVisited(testData, guardStartingRowPosition, guardStartingColumnPosition)
	}

	return d.solution1, d.solution2
}

func (d *Day) findUniqueCoordinatesVisited(data [][]string, guardRowPosition int, guardColumnPosition int) []string {
	var allCoordinatesVisited []string

	finish := false

	for !finish {
		allCoordinatesVisited = append(allCoordinatesVisited, strconv.Itoa(guardRowPosition)+","+strconv.Itoa(guardColumnPosition))
		finish = d.checkPositionOnBoundary(guardRowPosition, guardColumnPosition, data)

		if !finish {
			finish = d.isInfiniteLoop(allCoordinatesVisited, finish)
		}

		if !finish {
			if data[guardRowPosition][guardColumnPosition] == "^" {
				transposedColumn := matrix.TransposeColumn(data, guardColumnPosition)

				if transposedColumn[guardRowPosition-1] == "#" || transposedColumn[guardRowPosition-1] == "O" {
					data[guardRowPosition][guardColumnPosition] = ">"
				} else {
					data[guardRowPosition-1][guardColumnPosition] = "^"
					guardRowPosition--
				}
			}

			if data[guardRowPosition][guardColumnPosition] == ">" {
				if data[guardRowPosition][guardColumnPosition+1] == "#" || data[guardRowPosition][guardColumnPosition+1] == "O" {
					data[guardRowPosition][guardColumnPosition] = "v"
				} else {
					data[guardRowPosition][guardColumnPosition+1] = ">"
					guardColumnPosition++
				}
			}

			if data[guardRowPosition][guardColumnPosition] == "v" {
				transposedColumn := matrix.TransposeColumn(data, guardColumnPosition)

				if transposedColumn[guardRowPosition+1] == "#" || transposedColumn[guardRowPosition+1] == "O" {
					data[guardRowPosition][guardColumnPosition] = "<"
				} else {
					data[guardRowPosition+1][guardColumnPosition] = "v"
					guardRowPosition++
				}
			}

			if data[guardRowPosition][guardColumnPosition] == "<" {
				if data[guardRowPosition][guardColumnPosition-1] == "#" || data[guardRowPosition][guardColumnPosition-1] == "O" {
					data[guardRowPosition][guardColumnPosition] = "^"
				} else {
					data[guardRowPosition][guardColumnPosition-1] = "<"
					guardColumnPosition--
				}
			}
		}
	}

	slices.Sort(allCoordinatesVisited)
	return slices.Compact(allCoordinatesVisited)
}

func (d *Day) isInfiniteLoop(allCoordinatesVisited []string, finish bool) bool {
	lastVisitedCoordinate := allCoordinatesVisited[len(allCoordinatesVisited)-1]
	countLastVisitedCoordinate := 0
	for _, coordinate := range allCoordinatesVisited {
		if lastVisitedCoordinate == coordinate {
			countLastVisitedCoordinate++
		}
	}

	if countLastVisitedCoordinate > 4 || len(allCoordinatesVisited) > (len(d.data)*len(d.data[0])) {
		d.solution2++
		finish = true
	}

	return finish
}

func (d *Day) checkPositionOnBoundary(rowIndex int, columnIndex int, data [][]string) bool {
	if rowIndex == 0 || columnIndex == 0 || rowIndex == len(data)-1 || columnIndex == len(data[rowIndex])-1 {
		return true
	}

	return false
}
