package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Hello, world")
	SetOperationFailed("this is original format, msg1: %s, msg2: %s.", "orig1", "orig2")
	
	SetOperationFailed("this is another format without verbs.")
}

func SetOperationFailed(format string, args ...interface{}) {
     	// revise format and append more verbs if we need to append a new error msg
	err := errors.New("another error")
	format = format + " We've got a new error, err: %v"
	args = append(args, err)
	
	PrintFunc(format, args...)
	
}
func PrintFunc(format string, args ...interface{}){
        msg := fmt.Sprintf(format, args...)
	
	fmt.Println(msg)

}


