package gradebook_test

import (
	"gradebook"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStudent(t *testing.T) {
	// TO DO: Add your code here.
	gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)

	want := gradebook.Student{ID: 1234, Name: "Jane Doe", HW1Grade: 95, HW2Grade: 85, HW3Grade: 92}
	got, err := gradebook.GetRecord(1234)

	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetRecordBadIDReturnsError(t *testing.T) {
	// TO DO: Add your code here.
	t.Parallel()

	gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)

	_, err := gradebook.GetRecord(9999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}
