package day6

import (
	"advent-of-code-2024/days/common/matrix"
	"advent-of-code-2024/days/common/utils"
	"advent-of-code-2024/input"
	"fmt"
	"slices"
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
	coordinatesVisited := d.findUniqueCoordinatesVisited(data)
	d.solution1 = len(coordinatesVisited)

	var coordinatesToPutBlocker []string

	for i, coordinate := range coordinatesVisited {
		if i != 0 {
			coordinatesToPutBlocker = d.addToCoordinatesIfItDoesNotExist(coordinatesToPutBlocker, coordinate)

			splitCoordinate := strings.Split(coordinate, ",")

			if utils.ConvertStringToInt(splitCoordinate[0]) > 0 {
				coordinatesToPutBlocker = d.addToCoordinatesIfItDoesNotExist(coordinatesToPutBlocker, fmt.Sprintf("%d,%s", utils.ConvertStringToInt(splitCoordinate[0])-1, splitCoordinate[1]))
			}

			if utils.ConvertStringToInt(splitCoordinate[1]) > 0 {
				coordinatesToPutBlocker = d.addToCoordinatesIfItDoesNotExist(coordinatesToPutBlocker, fmt.Sprintf("%d,%d", utils.ConvertStringToInt(splitCoordinate[0]), utils.ConvertStringToInt(splitCoordinate[1])-1))
			}

			if utils.ConvertStringToInt(splitCoordinate[0]) < len(d.data)-1 {
				coordinatesToPutBlocker = d.addToCoordinatesIfItDoesNotExist(coordinatesToPutBlocker, fmt.Sprintf("%d,%s", utils.ConvertStringToInt(splitCoordinate[0])+1, splitCoordinate[1]))
			}

			if utils.ConvertStringToInt(splitCoordinate[1]) < len(d.data)-1 {
				coordinatesToPutBlocker = d.addToCoordinatesIfItDoesNotExist(coordinatesToPutBlocker, fmt.Sprintf("%d,%d", utils.ConvertStringToInt(splitCoordinate[0]), utils.ConvertStringToInt(splitCoordinate[1])+1))
			}
		}
	}

	for _, coordinate := range coordinatesToPutBlocker {
		splitCoordinate := strings.Split(coordinate, ",")

		rowIndex := utils.ConvertStringToInt(splitCoordinate[0])
		colIndex := utils.ConvertStringToInt(splitCoordinate[1])

		if data[rowIndex][colIndex] != "^" {
			data = matrix.Create(d.data)
			testData := append([][]string(nil), matrix.Create(d.data)...)
			testData[rowIndex][colIndex] = "O"

			d.findUniqueCoordinatesVisited(testData)
		}
	}

	return d.solution1, d.solution2
}

func (d *Day) addToCoordinatesIfItDoesNotExist(coordinatesToPutBlocker []string, coordinate string) []string {
	if !slices.Contains(coordinatesToPutBlocker, coordinate) {
		coordinatesToPutBlocker = append(coordinatesToPutBlocker, coordinate)
	}
	return coordinatesToPutBlocker
}

func (d *Day) findUniqueCoordinatesVisited(data [][]string) []string {
	var uniqueCoordinatesVisited []string
	var allCoordinatesVisited []string

	finish := false

	for !finish {
		guardRowPosition := 0
		guardColumnPosition := 0
		for rowIndex, row := range data {
			indexOfGuard := slices.IndexFunc(row, func(i string) bool { return i == "^" || i == ">" || i == "v" || i == "<" })
			if indexOfGuard != -1 {
				guardRowPosition = rowIndex
				guardColumnPosition = indexOfGuard
			}
		}

		if data[guardRowPosition][guardColumnPosition] == "^" {
			uniqueCoordinatesVisited, allCoordinatesVisited = d.addPositionToUniqueCoordinates(guardRowPosition, guardColumnPosition, uniqueCoordinatesVisited, allCoordinatesVisited)
			finish = d.checkPositionOnBoundary(guardRowPosition, guardColumnPosition, data)

			if !finish {
				finish = d.isInfiniteLoop(allCoordinatesVisited, finish)
			}

			if !finish {
				transposedColumn := matrix.TransposeColumn(data, guardColumnPosition)

				if transposedColumn[guardRowPosition-1] == "#" || transposedColumn[guardRowPosition-1] == "O" {
					data[guardRowPosition][guardColumnPosition] = ">"
				} else {
					data[guardRowPosition][guardColumnPosition] = "X"
					data[guardRowPosition-1][guardColumnPosition] = "^"
				}
			}
		}

		if data[guardRowPosition][guardColumnPosition] == ">" {
			uniqueCoordinatesVisited, allCoordinatesVisited = d.addPositionToUniqueCoordinates(guardRowPosition, guardColumnPosition, uniqueCoordinatesVisited, allCoordinatesVisited)
			finish = d.checkPositionOnBoundary(guardRowPosition, guardColumnPosition, data)

			if !finish {
				finish = d.isInfiniteLoop(allCoordinatesVisited, finish)
			}

			if !finish {
				if data[guardRowPosition][guardColumnPosition+1] == "#" || data[guardRowPosition][guardColumnPosition+1] == "O" {
					data[guardRowPosition][guardColumnPosition] = "v"
				} else {
					data[guardRowPosition][guardColumnPosition] = "X"
					data[guardRowPosition][guardColumnPosition+1] = ">"
				}
			}
		}

		if data[guardRowPosition][guardColumnPosition] == "v" {
			uniqueCoordinatesVisited, allCoordinatesVisited = d.addPositionToUniqueCoordinates(guardRowPosition, guardColumnPosition, uniqueCoordinatesVisited, allCoordinatesVisited)
			finish = d.checkPositionOnBoundary(guardRowPosition, guardColumnPosition, data)
			transposedColumn := matrix.TransposeColumn(data, guardColumnPosition)

			if !finish {
				finish = d.isInfiniteLoop(allCoordinatesVisited, finish)
			}

			if !finish {
				if transposedColumn[guardRowPosition+1] == "#" || transposedColumn[guardRowPosition+1] == "O" {
					data[guardRowPosition][guardColumnPosition] = "<"
				} else {
					data[guardRowPosition][guardColumnPosition] = "X"
					data[guardRowPosition+1][guardColumnPosition] = "v"
				}
			}
		}

		if data[guardRowPosition][guardColumnPosition] == "<" {
			uniqueCoordinatesVisited, allCoordinatesVisited = d.addPositionToUniqueCoordinates(guardRowPosition, guardColumnPosition, uniqueCoordinatesVisited, allCoordinatesVisited)
			finish = d.checkPositionOnBoundary(guardRowPosition, guardColumnPosition, data)

			if !finish {
				finish = d.isInfiniteLoop(allCoordinatesVisited, finish)
			}

			if !finish {
				if data[guardRowPosition][guardColumnPosition-1] == "#" || data[guardRowPosition][guardColumnPosition-1] == "O" {
					data[guardRowPosition][guardColumnPosition] = "^"
				} else {
					data[guardRowPosition][guardColumnPosition] = "X"
					data[guardRowPosition][guardColumnPosition-1] = "<"
				}
			}
		}
	}
	return uniqueCoordinatesVisited
}

func (d *Day) isInfiniteLoop(allCoordinatesVisited []string, finish bool) bool {
	if len(allCoordinatesVisited) > (len(d.data) * len(d.data[0])) {
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

func (d *Day) addPositionToUniqueCoordinates(rowIndex int, columnIndex int, uniqueCoordinatesVisited []string, allCoordinatesVisited []string) ([]string, []string) {
	position := fmt.Sprintf("%d,%d", rowIndex, columnIndex)
	if !slices.Contains(uniqueCoordinatesVisited, position) {
		uniqueCoordinatesVisited = append(uniqueCoordinatesVisited, position)
	}

	allCoordinatesVisited = append(allCoordinatesVisited, position)

	return uniqueCoordinatesVisited, allCoordinatesVisited
}
