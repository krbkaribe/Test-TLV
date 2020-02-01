package main

import "testing"

func TestFunctionTLV(t *testing.T) {
	mapResult := make(map[string]string)
	arrBytes := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 49, 48, 48}
	strArr := string(arrBytes[:])
	var err string

	err = generateMap(mapResult, arrBytes)
	validateMap(strArr, mapResult, &err)
	printResult(mapResult, err)

	if err == "" {
		t.Error("Error durante el Test")
		t.Fail()
	} else {
		t.Log("Test finalizado correctamente")
	}
}
