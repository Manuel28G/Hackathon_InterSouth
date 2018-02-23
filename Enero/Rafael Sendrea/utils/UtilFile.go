package Utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Funcion que lee el archivo .txt y devuelve un objeto de tipo tablero para ser resuelto
func ReadFile(pathToFile string) *Board {

	var stringContent string
	boardToReturn := new(Board)

	file, err := os.Open(pathToFile)

	if err != nil {
		fmt.Printf("No puede encontrarse el archivo %s", pathToFile)
		fmt.Println(err.Error())
		return boardToReturn
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	for {
		fileContentString, readerError := fileReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}

		stringContent += fileContentString

		//Esto es para asegurarnos de leer solo la primera linea
		break
	}

	//fmt.Println(strings.Split(stringContent, ","))

	//Obtenemos el arreglo del archivo y vamos a colorar "0" en los espacios vacios
	valuesArray := FillEmptySpaces(strings.Split(stringContent, ","))

	//fmt.Println(valuesArray)

	//convertir el arreglo de strings a int para asignarlo a board
	intValueArray := ConvertStringArraytoInt(valuesArray)

	//Por ultimo le asignamos los valores al tablero
	boardToReturn.BoardLength = len(valuesArray)
	boardToReturn.AllValues = intValueArray

	return boardToReturn
}

func WriteFile(content string, fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, content)
}
