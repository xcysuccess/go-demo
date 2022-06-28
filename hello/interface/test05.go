package main

import "fmt"

type MyError struct {}

func (i MyError) Error() string {
	return "MyError"
}

func main() {
	err := Process()
	fmt.Println(err)

	fmt.Println(err == nil)
}

func Process() error {
	var err *MyError = nil
	return err
}