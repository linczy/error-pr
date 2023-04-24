package main

import (
	"errors"
	"fmt"
	"log"
)
import "github.com/joomcode/errorx"

func main() {
	fmt.Println("Application starting")

	//This config print outs the line where the logger was used:
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("--------Simple error------------")
	err := throwError1()
	if err != nil {
		log.Printf("Error: %+v", err)
	}

	fmt.Println("--------WRAP error------------")
	err = throwError1Wrap()
	if err != nil {
		log.Printf("Error: %+v", err)
	}

	fmt.Println("--------ErrorX------------")
	err = throwError1X()
	if err != nil {
		log.Printf("Error: %+v", err)
	}
}

func throwError2X() error {
	return errorx.IllegalState.New("Level 2 error message")
}

func throwError1X() error {
	err := throwError2X()
	if err != nil {
		return errorx.Decorate(err, "Level 1 error message")
	}
	return nil
}

func throwError2Wrap() error {
	return errors.New("Level 2 error message")
}

func throwError1Wrap() error {
	err := throwError2Wrap()
	if err != nil {
		return fmt.Errorf("Level 1 error message %w", err)
	}
	return nil
}

func throwError2() error {
	return errors.New("Level 2 error message")
}

func throwError1() error {
	err := throwError2()
	if err != nil {
		return fmt.Errorf("Level 1 error message %v", err)
	}
	return nil
}
