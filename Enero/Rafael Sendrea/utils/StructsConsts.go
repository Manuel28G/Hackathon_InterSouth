package Utils

import (
	"strconv"
	"time"
)

//Constantes necesarias
const (
	arrayLength int    = 81
	Width       int    = 9
	Height      int    = 9
	PathTofile  string = "entry.txt"
)

var (
	//Filas del tablero (posiciones)
	Row1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	Row2 = [...]int{9, 10, 11, 12, 13, 14, 15, 16, 17}
	Row3 = [...]int{18, 19, 20, 21, 22, 23, 24, 25, 26}
	Row4 = [...]int{27, 28, 29, 30, 31, 32, 33, 34, 35}
	Row5 = [...]int{36, 37, 38, 39, 40, 41, 42, 43, 44}
	Row6 = [...]int{45, 46, 47, 48, 49, 50, 51, 52, 53}
	Row7 = [...]int{54, 55, 56, 57, 58, 59, 60, 61, 62}
	Row8 = [...]int{63, 64, 65, 66, 67, 68, 69, 70, 71}
	Row9 = [...]int{72, 73, 74, 75, 76, 77, 78, 79, 80}

	//Columnas del tablero (posiciones)
	Col1 = [...]int{0, 9, 18, 27, 36, 45, 54, 63, 72}
	Col2 = [...]int{1, 10, 19, 28, 37, 46, 55, 64, 73}
	Col3 = [...]int{2, 11, 20, 29, 38, 47, 56, 65, 74}
	Col4 = [...]int{3, 12, 21, 30, 39, 48, 57, 66, 75}
	Col5 = [...]int{4, 13, 22, 31, 40, 49, 58, 67, 76}
	Col6 = [...]int{5, 14, 23, 32, 41, 50, 59, 68, 77}
	Col7 = [...]int{6, 15, 24, 33, 42, 51, 60, 69, 78}
	Col8 = [...]int{7, 16, 25, 34, 43, 52, 61, 70, 79}
	Col9 = [...]int{8, 17, 26, 35, 44, 53, 62, 71, 80}

	//Bloques del tablero (posiciones)
	Block1 = [...]int{0, 1, 2, 9, 10, 11, 18, 19, 20}
	Block2 = [...]int{3, 4, 5, 12, 13, 14, 21, 22, 23}
	Block3 = [...]int{6, 7, 8, 15, 16, 17, 24, 25, 26}
	Block4 = [...]int{27, 28, 29, 36, 37, 38, 45, 46, 47}
	Block5 = [...]int{30, 31, 32, 39, 40, 41, 48, 49, 50}
	Block6 = [...]int{33, 34, 35, 42, 43, 44, 51, 52, 53}
	Block7 = [...]int{54, 55, 56, 63, 64, 65, 72, 73, 74}
	Block8 = [...]int{57, 58, 59, 66, 67, 68, 75, 76, 77}
	Block9 = [...]int{60, 61, 62, 69, 70, 71, 78, 79, 80}
)

//Esctructura de Casilla
type Box struct {
	Value         int
	PosibleValues []int
}

//Estructura de tablero
type Board struct {
	BoardLength    int
	AllValues      []int
	InnerBox       [Width][Height]Box
	EmptyPositions []int
}

//Funciones de la estructura Board

//Llenar el tablero con los valores de su arreglo AllValues
func (board *Board) FillInnerBox() {

	k := 0

	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			board.InnerBox[i][j].Value = board.AllValues[k]
			k++
		}
	}
}

//Llenar arreglo de posiciones vacias
func (board *Board) FillEmptyPositions() {
	for i := 0; i < len(board.AllValues); i++ {
		if board.AllValues[i] == 0 {
			board.EmptyPositions = append(board.EmptyPositions, i)
		}
	}
}

//Imprimir tablero en pantalla
func (board *Board) BoardToString() string {
	var toPrint string

	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {

			if (j+1)%3 == 0 && j > 0 && j < (Height-1) {
				toPrint += strconv.Itoa(board.InnerBox[i][j].Value) + " | "
			} else {
				toPrint += strconv.Itoa(board.InnerBox[i][j].Value) + " "
			}
		}

		toPrint += "\n"

		if (i+1)%3 == 0 && i > 0 && i < (Width-1) {
			toPrint += "------+-------+------ \n"
		}
	}

	return toPrint

}

func (board *Board) SolveBacktrack(position int) bool {
	if position > len(board.AllValues) { // Ya todas las casillas vacias estan llenas, abortar
		return true
	}

	for x := 1; x < 10; x++ {
		board.AllValues[position] = x

		if board.IsValidNumber(position) { // Check for collisions
			if board.SolveBacktrack(board.NextPosition(position)) { // Moverse a siguiente casilla
				return true // Ya todas las casillas vacias estan llenas, abortar
			}
		}
	}

	board.AllValues[position] = 0 // Vaciar casilla
	return false                  //Solucion no encontrada, Backtrack
}

//Obtener siguiente posicion vacia
func (board *Board) NextPosition(position int) int {
	//defer timeTrack(time.Now(), "NextPosition")
	toReturn := 0

	for i := 0; i < len(board.EmptyPositions); i++ {
		if board.EmptyPositions[i] == position {
			if i < len(board.EmptyPositions)-1 {
				toReturn = board.EmptyPositions[i+1]
			} else {
				toReturn = 999999
			}
			break
		}
	}

	return toReturn
}

//Verificar si el numero colocado es valido
func (board *Board) IsValidNumber(position int) bool {
	//defer timeTrack(time.Now(), "IsValidNumber")
	if len(board.GetRow(position)) != len(RemoveDuplicates(board.GetRow(position))) || len(board.GetCol(position)) != len(RemoveDuplicates(board.GetCol(position))) || len(board.GetBlock(position)) != len(RemoveDuplicates(board.GetBlock(position))) {
		return false
	}

	return true
}

//Obtener arreglo de valores de la fila en la que se encuentra la posicion otorgada
func (board *Board) GetRow(position int) [Width]int {
	//defer timeTrack(time.Now(), "GetRow")
	var toReturn [Width]int
	var positions [Width]int

	if ValueInArray(position, Row1) {
		positions = Row1
	} else if ValueInArray(position, Row2) {
		positions = Row2
	} else if ValueInArray(position, Row3) {
		positions = Row3
	} else if ValueInArray(position, Row4) {
		positions = Row4
	} else if ValueInArray(position, Row5) {
		positions = Row5
	} else if ValueInArray(position, Row6) {
		positions = Row6
	} else if ValueInArray(position, Row7) {
		positions = Row7
	} else if ValueInArray(position, Row8) {
		positions = Row8
	} else if ValueInArray(position, Row9) {
		positions = Row9
	}

	for i := 0; i < len(positions); i++ {
		toReturn[i] = board.AllValues[positions[i]]
	}

	return toReturn
}

//Obtener arreglo de valores de la columna en la que se encuentra la posicion otorgada
func (board *Board) GetCol(position int) [Width]int {
	//defer timeTrack(time.Now(), "GetCol")
	var toReturn [Width]int
	var positions [Width]int

	if ValueInArray(position, Col1) {
		positions = Col1
	} else if ValueInArray(position, Col2) {
		positions = Col2
	} else if ValueInArray(position, Col3) {
		positions = Col3
	} else if ValueInArray(position, Col4) {
		positions = Col4
	} else if ValueInArray(position, Col5) {
		positions = Col5
	} else if ValueInArray(position, Col6) {
		positions = Col6
	} else if ValueInArray(position, Col7) {
		positions = Col7
	} else if ValueInArray(position, Col8) {
		positions = Col8
	} else if ValueInArray(position, Col9) {
		positions = Col9
	}

	for i := 0; i < len(positions); i++ {
		toReturn[i] = board.AllValues[positions[i]]
	}

	return toReturn
}

//Obtener arreglo de valores del bloque en la que se encuentra la posicion otorgada
func (board *Board) GetBlock(position int) [Width]int {
	//defer timeTrack(time.Now(), "GetBlock")

	var toReturn [Width]int
	var positions [Width]int

	if ValueInArray(position, Block1) {
		positions = Block1
	} else if ValueInArray(position, Block2) {
		positions = Block2
	} else if ValueInArray(position, Block3) {
		positions = Block3
	} else if ValueInArray(position, Block4) {
		positions = Block4
	} else if ValueInArray(position, Block5) {
		positions = Block5
	} else if ValueInArray(position, Block6) {
		positions = Block6
	} else if ValueInArray(position, Block7) {
		positions = Block7
	} else if ValueInArray(position, Block8) {
		positions = Block8
	} else if ValueInArray(position, Block9) {
		positions = Block9
	}

	for i := 0; i < len(positions); i++ {
		toReturn[i] = board.AllValues[positions[i]]
	}

	return toReturn
}

func (board *Board) timeTrack(start time.Time, name string) {
	elapsed := time.Since(start).String()
	elapsed = "La solucion tardo " + elapsed
	content := "Sudoku resuelto: \n" + board.BoardToString() + "\n" + elapsed
	WriteFile(content, "Rafael_Sendrea_HTIS_Enero.txt")
}

func (board *Board) Solve() {
	defer board.timeTrack(time.Now(), "Solve")
	board.SolveBacktrack(board.EmptyPositions[0])
	board.FillInnerBox()
}
