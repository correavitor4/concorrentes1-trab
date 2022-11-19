package processingManager

import "fmt"

// import "fmt"

// "fmt"

type multiplierLinesPerMatrix2Strucuture struct {
	matrix1Line *[2500]int        // The line of the matrix1 that will be multiplied
	matrix2     *[2500][2500]int // The matrix2 that will be used to multiply
	matrix3Line *[2500]int        // The line of the matrix3 that will be filled after calc
}

func execCalc(mp *multiplierLinesPerMatrix2Strucuture) {
	var m3LineIndex int = 0
	for columnIndex := 0; columnIndex < len(mp.matrix2); columnIndex++ {
		var columnToCalc = getColumnByIndex(mp.matrix2, columnIndex)
		mp.matrix3Line[m3LineIndex] = multiPlyLinePerColumn(mp.matrix1Line, columnToCalc)
		m3LineIndex++
	}
}

func getColumnByIndex(m *[2500][2500]int, index int) [2500]*int {
	var column [2500]*int
	for i := 0; i < len(m); i++ {
		column[i] = &m[i][index]
	}
	return column
}

func multiPlyLinePerColumn(line *[2500]int, column [2500]*int) int {
	var result int
	for i := 0; i < len(line); i++ {
		// fmt.Println("line[i]: ", *column[i])
		result += line[i] * *column[i]
	}
	return result
}

func processLines(linesToCalc []int, matrix1 *[2500][2500]int, matrix2 *[2500][2500]int, matrix3 *[2500][2500]int) {
	for i := 0; i < len(linesToCalc); i++ {
		var lineToCalc = linesToCalc[i]
		var matrix1Line = matrix1[lineToCalc]
		var matrix3Line = matrix3[lineToCalc]
		var mp = multiplierLinesPerMatrix2Strucuture{
			matrix1Line: &matrix1Line,
			matrix2:     matrix2,
			matrix3Line: &matrix3Line,
		}
		execCalc(&mp)
		fmt.Println("Line ", lineToCalc, " calculated")
	}
}
