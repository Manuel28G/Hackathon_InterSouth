package main

import (
	"strings"
	"../programasudoku/Controller"
	"io/ioutil"
	"os"
	"fmt"
	"time"

)

var values, result, err = Controller.SquarePossibilities{}, Controller.SquarePossibilities{}, error(nil)

func main() {

	file, err := os.Create("View/Salida.txt")
	if err != nil {
		fmt.Print("Error al crear archivo")
		return
	}
	t := time.Now()
	asciiSubstring2 := t.Format(time.RFC3339Nano)[11:27]


	b, err := ioutil.ReadFile("View/Entrada.txt") // just pass the file name
	if err != nil {
		file.WriteString("Error el archivo no existe")
		os.Exit(3)
	}

	defer file.Close()
	var	 hardest = strings.Split(string(b),"\n");
	for _, input := range hardest {

		file.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
		file.WriteString("Entrada: \n")
		file.WriteString(input+"\n")
		if values, err = Controller.ParseGrid(input); err == nil {

			file.WriteString("Grid is valid.\n")
		} else {

			file.WriteString("Parsed Grid: Illegal.\n")

		}
		file.WriteString("Resuelto:\n")
		result, err = Controller.Solve(input+"\n")
		if err != nil {

			file.WriteString("Ocurrio un error al intertar resolver la matriz.\n")

		} else {

			file.WriteString(Controller.Display(result)+"\n")
		}

		file.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	}
	t2 := time.Now()
	file.WriteString("Tiempo Inicial\n")
	file.WriteString(asciiSubstring2+"\n")
	asciiSubstring := t2.Format(time.RFC3339Nano)[11:27]
	file.WriteString("Tiempo Final\n")
	file.WriteString(asciiSubstring+"\n")
	file.WriteString("Diferencia\n")
	file.WriteString(t2.Sub(t).String())

}
