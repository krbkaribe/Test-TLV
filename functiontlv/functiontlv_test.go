package main

import "testing"

func TestFunctionTLV(t *testing.T) {
	mapResult := make(map[string]string)
	arrBytes := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48}

	err := generateMap(mapResult, arrBytes)

	if err == "" {
		t.Error("Error durante el Test")
		t.Fail()
	} else {
		t.Log("Test finalizado correctamente")
	}
}
