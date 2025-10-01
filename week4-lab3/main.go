package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.50
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0-4")
	}
	return nil
}

func main() {
	//var st Student = Student{ID:"1", Name:"ruthaichanok", Email:"reungthanoo_r@silpakorn.edu", Year:6, GPA:2.62}

	//st := Student = Student{ID:"1", Name:"ruthaichanok", Email:"reungthanoo_r@silpakorn.edu", Year:6, GPA:2.62}

	Students := []Student{
		{ID: "1", Name: "reuthaichanok", Email: "reungthanoo_r@silpakorn.edu", Year: 6, GPA: 2.62},
		{ID: "2", Name: "alice", Email: "alice@silpakorn.edu", Year: 4, GPA: 3.21},
	}
	newStudent := Student{ID: "3", Name: "trudy", Email: "trudy@silpakorn.edu", Year: 2, GPA: 3.92}
	students = append(students, newStudent)

	for i, student := range students {
		fmt.Printf("%id Honor = %v\n", i, student.IsHonor())
		fmt.Printf("%id Validation = %v\n", i, student.Validate())
	}

	fmt.Printf("Honor %v", st.IsHonor())
	fmt.Printf("Validation = %v\n", st.Validate())
}
