package meh

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic2(t *testing.T) {

	test := func() (err error) {
		defer SetOnError(&err)

		ReturnError(nil)
		ReturnError(errors.New("this is an error1"))
		ReturnError(errors.New("this is an error2"))

		return
	}

	fmt.Println(test())
}

func TestPanic3(t *testing.T) {

	test := func() (i int, err error) {

		defer OnError(func(e error) {
			i, err = 0, e
		})

		ReturnError(nil)
		ReturnError(errors.New("this is an error1"))
		ReturnError(errors.New("this is an error2"))

		return
	}

	fmt.Println(test())
}
