package day4

import (
	"advent-of-code-2024/days/common/matrix"
	"advent-of-code-2024/input"
	"math"
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
	data := matrix.Create(d.data)

	for rowIndex, row := range data {
		for columnIndex, _ := range row {
			if data[rowIndex][columnIndex] == "X" {
				indexOfX := columnIndex
				d.searchMasLeft(row, indexOfX)
				d.searchMasRight(row, indexOfX)

				transposedColumn := matrix.TransposeColumn(data, columnIndex)
				d.searchMasLeft(transposedColumn, rowIndex)
				d.searchMasRight(transposedColumn, rowIndex)

				diagonalLR := matrix.FindDiagonalForCoordinates(data, rowIndex, columnIndex)
				indexOfXInDiagonal := int(math.Min(float64(rowIndex), float64(columnIndex)))
				d.searchMasLeft(diagonalLR, indexOfXInDiagonal)
				d.searchMasRight(diagonalLR, indexOfXInDiagonal)

				diagonalRL := matrix.FindAntidiagonalForCoordinates(data, rowIndex, columnIndex)
				indexOfXInDiagonal = int(min(float64(len(data)-1), float64(rowIndex+columnIndex))) - columnIndex
				d.searchMasLeft(diagonalRL, indexOfXInDiagonal)
				d.searchMasRight(diagonalRL, indexOfXInDiagonal)
			}

			if data[rowIndex][columnIndex] == "A" {
				diagonalLR := matrix.FindDiagonalForCoordinates(data, rowIndex, columnIndex)
				indexOfAInDiagonal := int(math.Min(float64(rowIndex), float64(columnIndex)))

				if indexOfAInDiagonal == 0 || indexOfAInDiagonal == len(diagonalLR)-1 {
					continue
				}

				if (diagonalLR[indexOfAInDiagonal-1] == "M" && diagonalLR[indexOfAInDiagonal+1] == "S") || (diagonalLR[indexOfAInDiagonal+1] == "M" && diagonalLR[indexOfAInDiagonal-1] == "S") {
					diagonalRL := matrix.FindAntidiagonalForCoordinates(data, rowIndex, columnIndex)
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

func (d *Day) searchMasLeft(data []string, endIndex int) {
	_, foundIndex := matrix.SearchLeft(data, endIndex, "SAM$")
	dataAsString := strings.Join(data[:endIndex], "")

	if foundIndex == len(dataAsString) {
		d.solution1 += 1
	}
}

func (d *Day) searchMasRight(data []string, startIndex int) {
	foundIndex, _ := matrix.SearchRight(data, startIndex, "^MAS")

	if foundIndex == 0 {
		d.solution1 += 1
	}
}
