package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

func (p *Student) debug() {
	fmt.Printf("I am debuging language pointer\n")
}

func main() {
	var s Student = Student{
		Name: "qcrao",
		Age: 18,
	}
	s.debug()
	fmt.Println(s)
}