package main

import (
	"fmt"
	"gradebook"
	"strings"
)

func main() {
	// gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)
	// err := gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)
	// fmt.Println(err)
	// gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)
	// gradebook.ModifyRecord(1234, 100, 100, 100)
	// r, _ := gradebook.GetRecord(1234)
	// fmt.Println(r)

	// gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)
	// gradebook.ModifyRecord(1234, 100, 100, 100)
	// r, _ := gradebook.GetRecord(1234)
	// fmt.Println(r)

	// fmt.Println(err)

	// gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)
	// record, _ := gradebook.GetRecord(1234)
	// fmt.Println(record)
	// _, err := gradebook.GetRecord(1235)
	// fmt.Println(err)
	// gradebook.AddStudent(1234, "Jane Doe", 95, 85, 92)
	// gradebook.AddStudent(1235, "John Smith", 72, 45, 88)
	// gradebook.DisplayAverage()

	// err := gradebook.ModifyRecord(1235, 100, 100, 100)
	// fmt.Println(err)
	// gradebook.AddStudent(1235, "John Smith", 72, 45, 88)
	// gradebook.DisplayAverage()
	// gradebook.ModifyRecord(1234, 100, 100, 100)

	// gradebook.DisplayAverage()

	// _, e := gradebook.GetRecord(1235)
	// fmt.Println(e)

	// s, e := gradebook.GetRecord(1234)
	// fmt.Println(s)

Loop:
	for {
		fmt.Println(`Available options:
						a to add a student
					        g to get a record
					        m to modify a record
					        d display average grades
						q to quit`)
		fmt.Print("Choose an option: ")
		// 		// TO DO: Add your code here

		var choice string
		fmt.Scan(&choice)
		fmt.Printf("%T", choice)
		choice = strings.ToLower(choice)
		choice = strings.TrimSpace(choice)

		switch choice {
		case "a":
			var id int
			fmt.Printf("\nInput the 4 digit id: ")
			fmt.Scan(&id)

			student, error := gradebook.GetRecord(id)
			if error != nil {
				fmt.Printf("%v\n\n", student)

				var fname string
				fmt.Print("Input the student's first name: ")
				fmt.Scan(&fname)
				var lname string
				fmt.Print("Input the student's last name: ")
				fmt.Scan(&lname)

				name := fname + " " + lname

				var gr1 float64
				fmt.Print("Input the student's first HW grade: ")
				fmt.Scan(&gr1)

				if gr1 < 0 || gr1 > 100 {
					fmt.Println("Invalid grade")
					continue Loop
				}

				var gr2 float64
				fmt.Print("Input the student's second HW grade: ")
				fmt.Scan(&gr2)

				if gr2 < 0 || gr2 > 100 {
					fmt.Println("Invalid grade")
					continue Loop
				}

				var gr3 float64
				fmt.Print("Input the student's third HW grade: ")
				fmt.Scan(&gr3)

				if gr3 < 0 || gr3 > 100 {
					fmt.Println("Invalid grade")
					continue Loop
				}

				gradebook.AddStudent(id, name, gr1, gr2, gr3)
				fmt.Printf("Student Added\n\n")

				err := gradebook.AddStudent(id, name, gr1, gr2, gr3)
				if err != nil {
					fmt.Printf("%v\n\n", err)
				} else {
					fmt.Printf("Record Added.\n\n")
				}

			} else {
				fmt.Printf("%v, already exists\n\n", student)
			}

		case "g":
			var id int
			fmt.Printf("\nInput the 4 digit id for the record to retrieve: ")
			fmt.Scan(&id)
			s, err := gradebook.GetRecord(id)
			if err != nil {
				fmt.Printf("%v\n\n", err)
			} else {
				fmt.Printf("%v\n\n", s)
			}

		case "m":
			var id int
			fmt.Print("Input the 4 digit id: ")
			fmt.Scan(&id)
			student, error := gradebook.GetRecord(id)
			// fmt.Println(error)
			if error != nil {
				fmt.Printf("%v\n\n", error)
				// break Loop
			} else {
				fmt.Printf("%v\n\n", student)

				var gr1 float64
				fmt.Print("Input the student's first HW grade: ")
				fmt.Scan(&gr1)

				var gr2 float64
				fmt.Print("Input the student's second HW grade: ")
				fmt.Scan(&gr2)

				var gr3 float64
				fmt.Print("Input the student's third HW grade: ")
				fmt.Scan(&gr3)

				err := gradebook.ModifyRecord(id, gr1, gr2, gr3)
				if err != nil {
					fmt.Printf("%v\n\n", err)

				} else {
					fmt.Printf("Record Modified.\n\n")
				}
			}

		case "d":
			gradebook.DisplayAverage()
			fmt.Printf("\n")
		case "q":
			break Loop
		}

	}
}
