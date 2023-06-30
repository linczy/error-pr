package main

import (
	"error-pr/pkg/stackerr"
	"errors"
	"fmt"
	"log"
	"os"
)
import "github.com/joomcode/errorx"

/*
CustomErrorType with additional field
*/
type CustomErrorType struct {
	message  string
	fdc3Code string
}

/*
Eevery struct is good for error just we need implement the error() function
*/
func (e *CustomErrorType) error() string {
	return fmt.Sprintf("Error: %s (code: %d)", e.message, e.fdc3Code)
}

var OneCustomError = CustomErrorType{
	message:  "Huge error",
	fdc3Code: "fdc3 error code",
}

/*
Dedicated error, you can use this to compare instead of checking the message itself.
*/
var DedicatedErrorInstance = fmt.Errorf("Level 2 error message")

var IllegalStateError = errorx.IllegalState.New("Level 2 error message")

func main() {
	main2()
}
func main1() {
	fmt.Println("Application starting")

	//This config prints outs the line where the logger was used:
	//log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("--------Simple error------------")
	err := throwError()
	if err != nil {
		log.Printf("Error: %+v", err)
	}
	log.Printf("Is DedicatedErrorInstance: %+v ", errors.Is(err, DedicatedErrorInstance))

	fmt.Println("--------WRAP error------------")
	err = throwErrorWrap()
	if err != nil {
		log.Printf("Error: %+v", err)
	}
	log.Printf("Is DedicatedErrorInstance: %+v ", errors.Is(err, DedicatedErrorInstance))

	fmt.Println("--------ErrorX------------")
	err = throwErrorX()
	if err != nil {
		log.Printf("Error: %+v", err)
	}
	log.Printf("Is xerror found: %+v", errors.Is(err, IllegalStateError))
	//or
	log.Printf("Is xerror IllegalState type: %+v", errorx.IsOfType(err, errorx.IllegalState))

}

func throwRootErrorX() error {
	return IllegalStateError
}

func throwErrorX() error {
	err := throwRootErrorX()
	if err != nil {
		return errorx.Decorate(err, "Level 1 error message")
	}
	return nil
}

func throwErrorWrap() error {
	err := throwDedicatedError()
	if err != nil {
		return fmt.Errorf("Level 1 error message %w", err)
	}
	return nil
}

func throwDedicatedError() error {
	return DedicatedErrorInstance
}

func throwError() error {
	err := throwDedicatedError()
	if err != nil {
		return fmt.Errorf("Level 1 error message %v", err)
	}
	return nil
}

func main2() {
	err := fmt.Errorf("original error")
	err = stackerr.NewStackTraceError(err, "something went wrong")
	fmt.Println(err.Error())
	fmt.Fprintln(os.Stderr, err.(interface{ StackTrace() string }).StackTrace())
}
