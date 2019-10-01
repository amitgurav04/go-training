package main

import "fmt"

type contactInfo struct {
	city  string
	phone string
	email string
}

type Student struct {
	firstName string
	lastName  string
	std       int
	aggregate float64
	contactInfo
}

func (s Student) print() {
	fmt.Printf("Student: %+v", s)
}


func (s *Student) updateFirstName(fName string) {

	s.firstName = fName
}


















