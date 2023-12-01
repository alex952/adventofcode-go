package day1

import (
	"reflect"
	"testing"
)

func TestReadCalibrationInputs(t *testing.T) {
	input_line := "1eightwo"
	input := ReadCalibrationLine(input_line)

	if !reflect.DeepEqual(input.calibrationValues, []byte(input_line)) {
		t.Fatalf("Incorrect read of line %s", input_line)
	}
}
