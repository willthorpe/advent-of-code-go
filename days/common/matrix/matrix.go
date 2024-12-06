package matrix

import (
	"advent-of-code-2024/days/common/utils"
	"math"
	"strings"
)

func Create(array []string) [][]string {
	var matrix [][]string

	for _, row := range array {
		splitRow := strings.Split(row, "")
		matrix = append(matrix, splitRow)
	}
	return matrix
}

func TransposeColumn(data [][]string, columnIndex int) []string {
	var columnAsRow []string
	for _, da := range data {
		columnAsRow = append(columnAsRow, da[columnIndex])
	}
	return columnAsRow
}

func FindDiagonalForCoordinates(data [][]string, rowIndex int, columnIndex int) []string {
	var diagonal []string
	startRow := int(math.Max(float64(rowIndex)-float64(columnIndex), 0))
	startColumn := int(math.Max(float64(columnIndex)-float64(rowIndex), 0))
	minimumBoundary := int(math.Min(float64(len(data)), float64(len(data[0]))))
	maximumStartingPosition := int(math.Max(float64(startRow), float64(startColumn)))
	diagonalRange := minimumBoundary - maximumStartingPosition

	for i := range diagonalRange {
		diagonal = append(diagonal, data[startRow+i][startColumn+i])
	}

	return diagonal
}

func FindAntidiagonalForCoordinates(data [][]string, rowIndex int, columnIndex int) []string {
	var diagonal []string
	c := rowIndex + columnIndex
	column := int(math.Min(float64(c), float64(len(data)-1)))
	row := c - column
	stop := false

	for column >= 0 && !stop {
		rowLength := len(data[row])
		diagonal = append(diagonal, data[row][column])

		column -= 1
		row += 1

		if row == rowLength {
			stop = true
		}
	}

	return diagonal
}

func SearchRight(data []string, startIndex int, regex string) (int, int) {
	dataAsString := strings.Join(data[startIndex+1:], "")

	return utils.FindPositionOfRegexInString(dataAsString, regex)
}

func SearchLeft(data []string, endIndex int, regex string) (int, int) {
	dataAsString := strings.Join(data[:endIndex], "")

	return utils.FindPositionOfRegexInString(dataAsString, regex)
}
