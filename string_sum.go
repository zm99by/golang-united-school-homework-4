package string_sum

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestFailWhenEmptyInputStringSum(t *testing.T) {
	s := "   "
	expected := ""

	actual, err := StringSum(s)

	if !reflect.DeepEqual(actual, expected) || err == nil || err.Error() != "e1: input is empty" {
		fmt.Println(expected, len(expected))
		fmt.Println(actual, len(actual))
		t.Errorf("error, man")
	}
}

func TestFailWhenLessTwoOperands(t *testing.T) {
	s := "555"
	expected := ""

	actual, err := StringSum(s)

	errorMessage := "e2: expecting two operands, but received more or less"
	if !reflect.DeepEqual(actual, expected) || err == nil || errors.Unwrap(err).Error() == errorMessage {
		fmt.Println(expected, len(expected))
		fmt.Println(actual, len(actual))
		t.Errorf(err.Error())
	}
}

func TestFailWhenTooMuchOperands(t *testing.T) {
	s := "-1 - 2 + 88"
	expected := ""

	actual, err := StringSum(s)

	errorMessage := "e3: expecting two operands, but received more or less"
	if !reflect.DeepEqual(actual, expected) || err == nil || errors.Unwrap(err).Error() == errorMessage {
		fmt.Println(expected, len(expected))
		fmt.Println(actual, len(actual))
		t.Errorf(err.Error())
	}
}

func TestFailWhenInputIncorrect(t *testing.T) {
	s := "24c+55"
	expected := ""
	expectedErr := &strconv.NumError{Func: "Atoi", Num: "24c", Err: strconv.ErrSyntax}

	actual, err := StringSum(s)

	if !reflect.DeepEqual(actual, expected) || err == nil || !reflect.DeepEqual(errors.Unwrap(err), expectedErr) {
		fmt.Println(expected, len(expected))
		fmt.Println(actual, len(actual))
		t.Errorf(err.Error())
	}
}

func TestSumOfPositiveOperands(t *testing.T) {
	s := "11 + 33"
	expected := "44"

	actual, err := StringSum(s)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}

func TestSumOfDifferentOperands(t *testing.T) {
	s := "-50 + 1"
	expected := "-49"

	actual, err := StringSum(s)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}

func TestSumOfNegativeOperands(t *testing.T) {
	s := "- 48 + -22"
	expected := "-70"

	actual, err := StringSum(s)

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(expected, actual) {
		fmt.Println(expected)
		fmt.Println(actual)
		t.Errorf("error, man")
	}
}
