package processingManager

import (
	"fmt"
	"sync"
	"time"
)

func StartProcessing(matriz *[2500][2500]int, matriz2 *[2500][2500]int, matrizFinal *[2500][2500]int) {
	fmt.Println("Processing started")
	
	var start1 = time.Now()
	normalProcess(matriz, matriz2, matrizFinal)
	var elapsed1 = time.Since(start1)

	var start2 = time.Now()
	multiThreadProcess(matriz, matriz2, matrizFinal)
	var elapsed2 = time.Since(start2)

	fmt.Println("Normal process time of execution: ", elapsed1)
	fmt.Println("Multi thread process time of execution: ", elapsed2)
	
}

// func checkFinalMatrixValue(m *[2500][2500]int, valueToCheck int) (bool) {
// 	var check bool = true
// 	for i := 0; i < len(m); i++ {
// 		for j := 0; j < len(m[i]); j++ {
// 			if m[i][j] != valueToCheck {
// 				check = false
// 				break
// 			}
// 		}
// 	}
// 	return check
// }

func normalProcess(matriz *[2500][2500]int, matriz2 *[2500][2500]int, matrizFinal *[2500][2500]int){
	//const routinesNumber int = 9
	var linesToCalc = make([]int, 0)
	for i := 0; i < len(matriz); i++ {
		linesToCalc = append(linesToCalc, i)
	}
	//var linesThatEachRoutineWillBeProcessing = returnLinesThatEachRoutineWillProcess(matriz, routinesNumber)
	
	processLines(linesToCalc, matriz, matriz2, matrizFinal)
	
}


func multiThreadProcess(matriz *[2500][2500]int, matriz2 *[2500][2500]int, matrizFinal *[2500][2500]int){
	const routinesNumber int = 9
	var linesThatEachRoutineWillBeProcessing = returnLinesThatEachRoutineWillProcess(matriz, routinesNumber)
	var wg sync.WaitGroup
	wg.Add(routinesNumber)

	for i := 0; i < routinesNumber; i++ {
		go func(i int) {
			processLines(linesThatEachRoutineWillBeProcessing[i].linesToCalc, matriz, matriz2, matrizFinal)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func returnLinesThatEachRoutineWillProcess(m *[2500][2500]int, routinesNumber int) []splitedWork {
	//1. Create a slice of the lines that each routine will process (structs slice)
	var linesToAtributterWork = make([]splitedWork, routinesNumber)

	//2. init each splitedWork struct in linesToAtributterWork
	for i := 0; i < routinesNumber; i++ {
		linesToAtributterWork[i].initSplitedWorkStruct()
	}

	var currentRoutineIndex int = 0
	for i := 0; i < len(m); i++ {
		linesToAtributterWork[currentRoutineIndex].linesToCalc = append(linesToAtributterWork[currentRoutineIndex].linesToCalc, i)

		if currentRoutineIndex == routinesNumber-1 {
			currentRoutineIndex = 0
		} else {
			currentRoutineIndex++
		}
	}

	return linesToAtributterWork
}

type splitedWork struct {
	linesToCalc []int
}

func (s splitedWork) initSplitedWorkStruct() {
	s.linesToCalc = make([]int, 0)
}
