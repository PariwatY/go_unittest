package services_test

import (
	"go_unittest/services"
	"testing"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	caseTest := []testCase{
		{name: "should_success_when_inputGrade80_then_returnA", score: 80, expected: "A"},
		{name: "should_success_when_inputGrade74_then_returnB", score: 74, expected: "B"},
		{name: "should_success_when_inputGrade65_then_returnC", score: 65, expected: "C"},
		{name: "should_success_when_inputGrade50_then_returnD", score: 54, expected: "D"},
		{name: "should_success_when_inputGradeLessthan50_then_returnF", score: 2, expected: "F"},
	}

	lengthCaseTest := len(caseTest)

	for i := 0; i < lengthCaseTest; i++ {
		t.Run(caseTest[i].name, func(t *testing.T) {
			grade := services.CheckGrade(caseTest[i].score)
			expected := caseTest[i].expected

			if grade != expected {
				t.Errorf("got %v expected %v", grade, expected)
			}
		})
	}

}
