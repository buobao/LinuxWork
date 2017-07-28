package main

import (
	"errors"
	"fmt"
	"os"
)

/**
 * 自定义Error类型，实现内建Error接口
 * type Error interface {
 *      Error() string
 * }
 */
type MyError struct {
	arg int
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

/**
 * Failed[*errors.errorString]: Bad Arguments - negative!
 * Success:  16
 * Failed[*main.MyError]: 1000 - Bad Arguments - too large!
 *
 * Recovered! Panic message[Cannot find specific file]
 * 4 3 2 1 0
 */
func main() {
	// 1.Test error
	args := []int{-1, 4, 1000}
	for _, i := range args {
		if r, e := testError(i); e != nil {
			fmt.Printf("Failed[%T]: %v\n", e, e)
		} else {
			fmt.Println("Success: ", r)
		}
	}

	// 2.Test defer
	src, err := os.Open("control.go")
	if err != nil {
		fmt.Printf("打开文件错误[%v]\n", err)
		return
	}
	defer src.Close()
	// use src...

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	// 3.Test panic/recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered! Panic message[%s]\n", r)
		}
	}()

	_, err2 := os.Open("test.go")
	if err2 != nil {
		panic("Cannot find specific file")
	}
}

func testError(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Bad Arguments - negative!")
	} else if arg > 256 {
		return -1, &MyError{arg, "Bad Arguments - too large!"}
	} else {
		return arg * arg, nil
	}
}
