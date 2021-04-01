package main

import (
	"strings"
	"testing"
)

//testPrinting checks whether distragin strings works well
func TestPrinting(tt *testing.T) {
	testingString := "its been a very long and dynamic day"
	tempStringcontArray := strings.Fields(testingString)
	if len(tempStringcontArray) != 8 {
		tt.Error("Diffrent number of words has been calculated")
	}
}
