package main

import (
	"fmt"
	"strconv"
)

func validateMap(strArr string, mapResult map[string]string, err *string) {
	*err = "Error: -"
	for i := 0; i < len(strArr); i++ {
		length, e := strconv.ParseInt(strArr[i:(i+2)], 10, 8)
		if e != nil {
			*err = "Error: El arreglo tiene el formato incorrecto. El largo es inválido."
		}

		var value int64
		typea := strArr[(i + 2):(i + 3)]
		typeb, e := strconv.ParseInt(strArr[(i+3):(i+5)], 10, 8)

		if (typea != "A" && typea != "N") || e != nil {
			*err = "Error: El arreglo tiene el formato incorrecto. El tipo es inválido."
		} else if int(typeb+length) > len(strArr) {
			*err = "Error: El arreglo tiene el formato incorrecto. El tipo sobrepasa el arreglo."
		} else {
			if typea == "N" {
				value, e = strconv.ParseInt((strArr[typeb:(typeb + length)]), 10, 64)
				if value == 0 && e != nil {
					*err = "Error: El arreglo tiene el formato incorrecto. El valor no es numérico."
				}
			}
		}
		mapResult["Largo"+strconv.Itoa(i)+":"] = strArr[i:(i + 2)]
		mapResult["Tipo"+strconv.Itoa(i)+" :"] = typea + strArr[(i+3):(i+5)]
		mapResult["Valor"+strconv.Itoa(i)+":"] = strArr[typeb:(typeb + length)]

		i = int(length) + int(typeb) - 1
	}
}

func generateMap(mapResult map[string]string, arrBytes []byte) (err string) {
	strArr := string(arrBytes[:])
	valArr, e := strconv.ParseInt(strArr, 10, 64)

	if len(strArr) > 0 && valArr > 0 && e == nil {
		err = "Error: El arreglo tiene el formato incorrecto. Es numérico."
	} else if strArr == "" {
		err = "Error: El arreglo tiene el formato incorrecto. Está vacío."
	} else {
		validateMap(strArr, mapResult, &err)
	}
	return err
}

func printResult(mapResult map[string]string, err string) {
	for key, value := range mapResult {
		fmt.Println(key, value)
	}
	fmt.Println("")
	fmt.Println(err)
}

func main() {
	mapResult := make(map[string]string)
	arrBytes := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 49, 48, 48}

	var err string = generateMap(mapResult, arrBytes)
	printResult(mapResult, err)
}
