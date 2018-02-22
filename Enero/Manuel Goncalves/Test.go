package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
)

var matriz = [][] string{
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},
[]string{"_", "_", "_", "_", "_", "_","_", "_", "_"},}

var containCero = true

func main(){
	init := getTimeInNanoSecond()
	readFile()
	end := getTimeInNanoSecond()
	fmt.Println("Finalizó en (nanoS):")
	fmt.Println(init-end)

}

func getTimeInMilis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getTimeInNanoSecond() int64 {
	return time.Now().UnixNano()
}

func readFile(){
	b, err := ioutil.ReadFile("C:\\Users\\manue\\Desktop\\file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	var texto = strings.Split(str,",")
	var col,row = 0,0
    for i := 0; i < len(texto);i++{
    	if texto[i] != ""{
			matriz[row][col] = texto[i]
		}else {
			matriz[row][col] = "0"
		}
    	if col >= 8 {
    		row ++
    		col = 0
		}else {
			col++
		}
	}
	resolve()
}
func onlyOneRow(_rowInit int,_rowEnd int, _col int)int{
	var response = 0
	for i:=_rowInit;i<=_rowEnd;i++{
		tmp:=matriz[i][_col]
		if(tmp == "0"){
			response++
		}
	}
	return response
}


func onlyOneCol( _colInit int,_colEnd int,_row int)int{
	var response = 0
	for i:=_colInit;i<=_colEnd;i++{
		tmp:=matriz[_row][i]
		if(tmp == "0"){
			response++
		}
	}
	return response
}
func findSquare(_num int,_square int) bool{
	var initRow,finalRow,initCol,finalCol = 0,0,0,0
	switch(_square){
	case 1:
		initCol = 0
		initRow = 0
		finalCol = 2
		finalRow = 2
		break
	case 2:
		initCol = 3
		initRow = 0
		finalCol = 5
		finalRow = 2
		break
	case 3:
		initCol = 6
		initRow = 0
		finalCol = 8
		finalRow = 2
		break
	case 4:
		initCol = 0
		initRow = 3
		finalCol = 2
		finalRow = 5
		break
	case 5:
		initCol = 3
		initRow = 3
		finalCol = 5
		finalRow = 5
		break
	case 6:
		initCol = 6
		initRow = 3
		finalCol = 8
		finalRow = 5
		break
	case 7:
		initCol = 0
		initRow = 6
		finalCol = 2
		finalRow = 8
		break
	case 8:
		initCol = 3
		initRow = 6
		finalCol = 5
		finalRow = 8
		break
	case 9:
		initCol = 6
		initRow = 6
		finalCol = 8
		finalRow = 8
		break
	}

	tmpRow := initRow
	tmpCol := initCol

	rowFind := -1
	colFind := -1

	for {
		resp := matriz[tmpRow][tmpCol];
		numInt:=strconv.Itoa(_num)
		if resp == numInt{
			return false
		}
		//Buscamos por columna y Fila

		if resp == "0"{
			containCero = true
			if !findColumn(_num,tmpCol) && !findRow(_num,tmpRow){
				if(colFind == -1) && (rowFind == -1){
					colFind = tmpCol
					rowFind = tmpRow
				}else {
					return false
				}
			}
		}

		//Vemos si es la única opción en esa casilla
	/*	if( onlyOneRow(initRow,finalRow,tmpCol) ==1 ){

			matriz[tmpRow][tmpCol] = strconv.Itoa(_num)
			return true
		}
*/
		tmpRow++
		if(tmpRow > finalRow){
				tmpCol ++
				tmpRow = initRow
		}

		if(tmpCol > finalCol){
			break
		}

	}
	matriz[rowFind][colFind] = strconv.Itoa(_num)
	return true

}

func findRow(_number int,_row int) bool{
	for cont:= 0 ;cont<9; cont++{
		var tmp = matriz[_row][cont];
		if tmp == strconv.Itoa(_number){
			return true
		}
	}
	return false;
}

func findColumn(_number int,_col int) bool{
	for cont:= 0 ;cont<9; cont++{
		var tmp = matriz[cont][_col];
		if tmp == strconv.Itoa(_number){
			return true
		}
	}
	return false;
}

func resolve(){

	for containCero == true{
		containCero = false
		for number:=1 ; number<=9 ; number++{
			for square:=1 ; square<=9 ; square++ {
				findSquare(number,square)
			}
		}
	}

	fmt.Println("Programa terminado")
	fmt.Println(matriz)

}
