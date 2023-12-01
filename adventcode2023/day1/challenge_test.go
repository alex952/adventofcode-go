package day1

import (
	"reflect"
	"testing"
)

func TestReadCalibrationInputs(t *testing.T) {
	input_line := "1eightwo"
	input := ReadCalibrationLine(input_line)

	if !reflect.DeepEqual(input.calibrationValues, []byte(input_line)) {
		t.Fail()
	}
}

func TestFirstPartFirstInt(t *testing.T) {
	input_line := "1eightwo2"
	input := ReadCalibrationLine(input_line)

	if first_int := input.firstPartFirstInt(); first_int != '1' {
		t.Fail()
	}
}

func TestFirstPartLastInt(t *testing.T) {
	input_line := "1eightwo2"
	input := ReadCalibrationLine(input_line)

	if last_int := input.firstPartLastInt(); last_int != '2' {
		t.Fail()
	}
}

func TestPartFirstInt(t *testing.T) {
	input_line := "1eightwo2"
	input := ReadCalibrationLine(input_line)

	if first_int := input.firstInt(); first_int != '1' {
		t.Fail()
	}
}

func TestPartFirstIntLetters(t *testing.T) {
	input_line := "eightwo2"
	input := ReadCalibrationLine(input_line)

	if first_int := input.firstInt(); first_int != '8' {
		t.Fail()
	}
}

func TestPartLastInt(t *testing.T) {
	input_line := "1eightwo2"
	input := ReadCalibrationLine(input_line)

	if last_int := input.lastInt(); last_int != '2' {
		t.Fail()
	}
}

func TestPartLastIntLetters(t *testing.T) {
	input_line := "1eighttwo"
	input := ReadCalibrationLine(input_line)

	if last_int := input.lastInt(); last_int != '2' {
		t.Fail()
	}
}

func TestPartLastIntLettersOverlap(t *testing.T) {
	input_line := "1eightwo"
	input := ReadCalibrationLine(input_line)

	if last_int := input.lastInt(); last_int != '2' {
		t.Fail()
	}
}

func TestFirstCalibrationnumber(t *testing.T) {
	input_line := "1eigh2three"
	input := ReadCalibrationLine(input_line)

	if calibration := input.FirstPartCalibrationNumber(); calibration != 12 {
		t.Fail()
	}
}

func TestCalibrationnumber(t *testing.T) {
	input_line := "1eigh2three"
	input := ReadCalibrationLine(input_line)

	if calibration := input.CalibrationNumber(); calibration != 13 {
		t.Fail()
	}
}

func TestCalibrationnumberOverlap(t *testing.T) {
	input_line := "1sevenine"
	input := ReadCalibrationLine(input_line)

	if calibration := input.CalibrationNumber(); calibration != 19 {
		t.Fail()
	}
}

func TestCalibrationnumberOverlapAllText(t *testing.T) {
	input_line := "sevenine"
	input := ReadCalibrationLine(input_line)

	if calibration := input.CalibrationNumber(); calibration != 79 {
		t.Fail()
	}
}

func TestTotalCalibrationFirstPart(t *testing.T) {
	input_lines := []string{"sevenine", "1orsienteight"}
	var calibration_data []*CalibrationInput

	for _, l := range input_lines {
		input := ReadCalibrationLine(l)
		calibration_data = append(calibration_data, input)
	}

	if calibration := TotalCalibration(calibration_data); calibration != 79+18 {
		t.Fail()
	}
}
