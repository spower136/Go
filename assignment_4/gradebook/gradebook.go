package gradebook

import "fmt"

type Student struct {
	ID       int
	Name     string
	HW1Grade float64
	HW2Grade float64
	HW3Grade float64
}

var grBook = make(map[int]Student)

// Create a new Student object with the given id, name and homework grades
// If there idea is already in the gradebook return an appropriate error.
// Otherwise, add the record to the gradebook.
func AddStudent(id int, name string, hw1gr, hw2gr, hw3gr float64) error {
	// TO DO: Add your code here

	_, found := grBook[id]
	if found {
		return fmt.Errorf("A student with ID %d is already in the gradebook", id)
	}

	newStudent := Student{
		ID:       id,
		Name:     name,
		HW1Grade: hw1gr,
		HW2Grade: hw2gr,
		HW3Grade: hw3gr}
	grBook[id] = newStudent

	return nil
}

// https://github.com/bitfield/ftl-code/blob/main/17.1/bookstore.go

// Returns the student record for the giving id if it exist, an error otherwise.
func GetRecord(id int) (Student, error) {
	// TO DO: Add your code here
	s, ok := grBook[id]
	if !ok {
		return Student{}, fmt.Errorf("ID %d doesn't exist", id)
	}
	return s, nil
}

// Modify the record with the given id in the gradebook.
// First check if there is already a record with that id.
// Modify the record if it exists and return nil error.
// If the record doesn't not exist return an appropriate error
func ModifyRecord(id int, hw1gr, hw2gr, hw3gr float64) error {
	student, found := grBook[id]
	// TO DO: Add your code here
	if found {
		student.HW1Grade = hw1gr
		student.HW2Grade = hw2gr
		student.HW3Grade = hw3gr
		grBook[id] = student
		return nil
	}
	return fmt.Errorf("ID %d does not exist", id)
}

// Display each students name and average of 3 HW grades to 2 decimal places
func DisplayAverage() {
	// TO DO: Add your code here
	fmt.Println("\nStudent Homework Grade Averages:")
	for _, v := range grBook {
		name := v.Name
		avg := (v.HW1Grade + v.HW2Grade + v.HW3Grade) / 3
		fmt.Printf("%s %.2f\n", name, avg)
	}
}
