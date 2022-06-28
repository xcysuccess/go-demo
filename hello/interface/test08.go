package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func main() {
	var i interface{} = new(Student)
	/*此时会发生panic，因为new出来的是*Student类型*/
	//s := i.(Student)
	s := i.(*Student)

	fmt.Println(s)
}