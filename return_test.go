package meh

import (
	"errors"
	"github.com/simulatedsimian/assert"
	"testing"
)

func test1(val int) (err error) {
	defer SetOnError(&err)

	switch val {
	case 0:
		ReturnError(nil)
	case 1:
		ReturnError(errors.New("Error1"))
	case 2:
		ReturnError(errors.New("Error2"))
	case 3:
		panic("panic")
	}

	return
}

func TestReturn1(t *testing.T) {

	assert.Equal(t, test1(0), nil)
	assert.Equal(t, test1(1).Error(), "Error1")
	assert.Equal(t, test1(2).Error(), "Error2")

	assert.MustPanic(t, func(t *testing.T) {
		test1(3)
	})
}

func test2(val int) (err error) {
	defer OnError(func(e error) {
		err = e
		if val == 2 {
			panic("panic")
		}
	})

	switch val {
	case 0:
		ReturnError(nil)
	case 1:
		ReturnError(errors.New("Error1"))
	case 2:
		ReturnError(errors.New("Error2"))
	case 3:
		panic("panic")
	}

	return
}

func TestReturn2(t *testing.T) {

	assert.Equal(t, test2(0), nil)
	assert.Equal(t, test2(1).Error(), "Error1")

	assert.MustPanic(t, func(t *testing.T) {
		test2(3)
	})

	assert.MustPanic(t, func(t *testing.T) {
		assert.Equal(t, test2(2).Error(), "Error2")
	})
}
