package meh

// messing with various aspects of The Go language
// Here Be Dragons

import (
	//	"errors"
	"fmt"
	"github.com/simulatedsimian/assert"
	"reflect"
	"testing"
)

// try/catch/finally impl.
// dont do this. its wrong. maybe

func TestThrow(t *testing.T) {

	Try(func() {
		Throw("dddd")
	}).Catch(func(e int) {
		fmt.Println("caught:", reflect.TypeOf(e), e)
	}).Catch(func(e string) {
		fmt.Println("caught:", reflect.TypeOf(e), e)
	}).Catch(func(e error) {
		fmt.Println("caught:", reflect.TypeOf(e), e)
	}).Finally(func() {
		fmt.Println("finally")
	})
}

func TestThrow2(t *testing.T) {

	assert.MustPanic(t, func(t *testing.T) {

		Try(func() {
		}).Catch(func(e int) {
			fmt.Println("caught:", reflect.TypeOf(e), e)
		}).Catch(func(e string) {
			fmt.Println("caught:", reflect.TypeOf(e), e)
		}).Catch(func(e error) {
			fmt.Println("caught:", reflect.TypeOf(e), e)
		}).Finally(func() {
			Throw("asdf")
		})

	})

}
