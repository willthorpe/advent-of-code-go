package day4

import (
	"advent-of-code-2024/input"
	"math"
	"regexp"
	"strings"
)

type Day struct {
	data      []string
	solution1 int
	solution2 int
}

func NewDay() *Day {
	i := input.NewInput("4", "2024")

	data := i.GetData()

	return &Day{
		data:      data,
		solution1: 0,
		solution2: 0,
	}
}

func (d *Day) Run() (int, int) {
	var data [][]string

	for _, row := range d.data {
		splitRow := strings.Split(row, "")
		data = append(data, splitRow)
	}

	for rowIndex, row := range data {
		for columnIndex, _ := range row {
			if data[rowIndex][columnIndex] == "X" {
				indexOfX := columnIndex
				d.searchMasLeftToRight(row, indexOfX)
				d.searchMasRightToLeft(row, indexOfX)

				var column []string
				for _, da := range data {
					column = append(column, da[columnIndex])
				}
				d.searchMasLeftToRight(column, rowIndex)
				d.searchMasRightToLeft(column, rowIndex)

				diagonalLR := d.createDiagonalLeftToRight(rowIndex, columnIndex, data)
				indexOfXInDiagonal := int(math.Min(float64(rowIndex), float64(columnIndex)))
				d.searchMasLeftToRight(diagonalLR, indexOfXInDiagonal)
				d.searchMasRightToLeft(diagonalLR, indexOfXInDiagonal)

				diagonalRL := d.createDiagonalRightToLeft(rowIndex, columnIndex, data)
				indexOfXInDiagonal = int(min(float64(len(data)-1), float64(rowIndex+columnIndex))) - columnIndex
				d.searchMasLeftToRight(diagonalRL, indexOfXInDiagonal)
				d.searchMasRightToLeft(diagonalRL, indexOfXInDiagonal)
			}

			if data[rowIndex][columnIndex] == "A" {
				diagonalLR := d.createDiagonalLeftToRight(rowIndex, columnIndex, data)
				indexOfAInDiagonal := int(math.Min(float64(rowIndex), float64(columnIndex)))

				if indexOfAInDiagonal == 0 || indexOfAInDiagonal == len(diagonalLR)-1 {
					continue
				}

				if (diagonalLR[indexOfAInDiagonal-1] == "M" && diagonalLR[indexOfAInDiagonal+1] == "S") || (diagonalLR[indexOfAInDiagonal+1] == "M" && diagonalLR[indexOfAInDiagonal-1] == "S") {
					diagonalRL := d.createDiagonalRightToLeft(rowIndex, columnIndex, data)
					indexOfAInDiagonal = int(min(float64(len(data)-1), float64(rowIndex+columnIndex))) - columnIndex

					if (diagonalRL[indexOfAInDiagonal-1] == "M" && diagonalRL[indexOfAInDiagonal+1] == "S") || (diagonalRL[indexOfAInDiagonal+1] == "M" && diagonalRL[indexOfAInDiagonal-1] == "S") {
						d.solution2 += 1
					}
				}
			}
		}
	}

	return d.solution1, d.solution2
}

func (d *Day) createDiagonalLeftToRight(rowIndex int, columnIndex int, data [][]string) []string {
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

func (d *Day) createDiagonalRightToLeft(rowIndex int, columnIndex int, data [][]string) []string {
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

func (d *Day) searchMasLeftToRight(data []string, startIndex int) {
	dataAsString := strings.Join(data[startIndex+1:], "")
	xmasR := regexp.MustCompile(`^MAS`)
	xmasIndex := xmasR.FindStringIndex(dataAsString)

	if len(xmasIndex) > 0 && xmasIndex[0] == 0 {
		d.solution1 += 1
	}
}

func (d *Day) searchMasRightToLeft(data []string, endIndex int) {
	dataAsString := strings.Join(data[:endIndex], "")
	samxR := regexp.MustCompile(`SAM$`)
	samxIndex := samxR.FindStringIndex(dataAsString)

	if len(samxIndex) > 0 && samxIndex[1] == len(dataAsString) {
		d.solution1 += 1
	}
}
