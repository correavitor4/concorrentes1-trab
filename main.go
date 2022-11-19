package main

import (
	_ "fmt"
	"github.com/correavitor4/concorrentes1-trab/pkg/processingManager"
)

func main() {
	//1. Create the matrixes
	var matrix1 [2500][2500]int
	var matrix2 [2500][2500]int

	// 2. Fill matrixes with the 2 int at all positions
	polutateMatrix(&matrix1)
	polutateMatrix(&matrix2)

	//3. Create the finalMatrix
	var finalMatrix [2500][2500]int

	// 4. Start processing
	processingManager.StartProcessing(&matrix1, &matrix2, &finalMatrix)
}

func polutateMatrix(matrix *[2500][2500]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 2
		}
	}
}