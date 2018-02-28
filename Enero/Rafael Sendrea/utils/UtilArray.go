package Utils

import (
	"strconv"
	"strings"
)

//Convertir un arreglo de string en int
func ConvertStringArraytoInt(array []string) []int {
	var intValueArray []int
	var actual int

	for i := 0; i < len(array); i++ {
		actual, _ = strconv.Atoi(strings.Trim(array[i], "\r\n"))
		//fmt.Printf("Convirtiendo %s de string a %d en int. \n", strings.Trim(array[i], "\r\n"), actual)
		intValueArray = append(intValueArray, actual)

	}

	return intValueArray
}

//Recibe un arreglo de strings y lo devuelve con 0 donde hayan espacios vacios
func FillEmptySpaces(array []string) []string {
	for i := 0; i < len(array); i++ {
		if array[i] == "" {
			array[i] = "0"
		}
	}

	return array
}

//Funcion que valida si un valor se encuentra en un arreglo
func ValueInArray(value int, array [Width]int) bool {

	/* fmt.Printf("Validando el valor %d en arreglo \n", value)
	fmt.Println(array) */

	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return true
		}
	}

	return false
}

//Funcion que valida si un valor se encuentra en un arreglo
func ValueInArray2(value int, array []int) bool {

	/* fmt.Printf("Validando el valor %d en arreglo \n", value)
	fmt.Println(array) */

	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return true
		}
	}

	return false
}

func RemoveDuplicates(elements [Width]int) []int {

	result := []int{}

	for i := 0; i < len(elements); i++ {
		if !ValueInArray2(elements[i], result) || elements[i] == 0 {
			result = append(result, elements[i])
		}
	}

	return result
}
