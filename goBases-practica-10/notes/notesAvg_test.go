package notes

import (
	"fmt"
	"testing"
)

func TestNotesAvg(t *testing.T) {
	// arrange
	student := []int{10, 4, 7, 8, 10}
	// student1 := []int{3, 5, -3, 3, 10}
	resultStudent := 0.0
	resultStudent1 := -1.0
	sum := 0

	for _, grade := range student {
		sum += grade
	}

	resultStudent = float64(sum) / float64(len(student))

	// act
	notesAvgStudent := notesAvg(10, 4, 7, 8, 10)
	notesAvgStudent1 := notesAvg(3, 5, -3, 3, 10)

	// assert
	if notesAvgStudent != resultStudent {
		t.Errorf("Expected %f, got %f", resultStudent, notesAvgStudent)

	} else {
		fmt.Println("OK!")
	}

	if notesAvgStudent1 != resultStudent1 {
		t.Errorf("Expected %f, got %f", resultStudent1, notesAvgStudent1)
	} else {
		fmt.Println("Ok!")
	}

}
