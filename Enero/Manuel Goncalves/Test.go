package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
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

func main(){
	init := getTimeInMilis()
	readFile()
	end := getTimeInMilis()
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
	fmt.Println(texto)
	var col,fil = 0,0
    for i := 0; i < len(texto);i++{
    	matriz[fil][col] = texto[i]
    	if col >= 8 {
    		fil ++
    		col = 0
		}else {
			col++
		}
	}
	fmt.Println(matriz) // print the content as a 'string'
}

func evaluateFill(_fill int) int{
	return 0
}

func evaluateColumn(_colum int) int{
	return 0
}

func evaluateSquare(_square int) int{
	return 0
}

func columnContain(_colum int,_value int) bool{
	return false
}
func fillContain(_fill int,_value int) bool{
	/*var tmp=""
	for i:=0;i<9 ;i++{
		tmp = matriz[_fill][i]
		i,e :=strconv.Atoi(_fill)
		if( tmp == i){
			return true
		}
	}*/
	return false
}
func squareContain(_square int,_value int) bool{
	return false
}