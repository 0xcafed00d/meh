package meh

import (
	"errors"
	"github.com/simulatedsimian/assert"
	"reflect"
	"testing"
)

type empty struct {
}

var errorType reflect.Type

func init() {
	errorType = reflect.TypeOf(errors.New("").(error))
}

func testEx1(f func()) (extype interface{}, finallyCalled bool) {
	Try(func() {
		f()
	}).Catch(func(e int) {
		extype = reflect.TypeOf(e)
	}).Catch(func(e string) {
		extype = reflect.TypeOf(e)
	}).Catch(func(e error) {
		extype = reflect.TypeOf(e)
	}).Finally(func() {
		finallyCalled = true
	})
	return
}

func TestThrow(t *testing.T) {
	extype, finallyCalled := testEx1(func() { Throw(5) })
	assert.Equal(t, extype, reflect.TypeOf(5))
	assert.True(t, finallyCalled)

	extype, finallyCalled = testEx1(func() { Throw("bang") })
	assert.Equal(t, extype, reflect.TypeOf("bang"))
	assert.True(t, finallyCalled)

	extype, finallyCalled = testEx1(func() { Throw(errors.New("bang")) })
	assert.Equal(t, extype, errorType)
	assert.True(t, finallyCalled)

	extype, finallyCalled = testEx1(func() {})
	assert.Nil(t, extype)
	assert.True(t, finallyCalled)

	assert.MustPanic(t, func(t *testing.T) {
		testEx1(func() {
			Throw(empty{})
		})
	})

	assert.MustPanic(t, func(t *testing.T) {
		Try(func() {
		}).Catch(func(e int) {
		}).Finally(func() {
			Throw("asdf")
		})
	})

	Try(func() {
		Throw(5)
	}).Catch(func(e int, e2 int) {
		assert.True(t, false) // should not get here
	}).Catch(func(e int) {
	}).Finally(func() {
	})
}
