package meh

import (
	"reflect"
)

// try/catch/finally impl.
// dont do this. its wrong. maybe
// Here Be Dragons

func Throw(e interface{}) {
	panic(e)
}

type Tryblock struct {
	try      func()
	catchers []interface{}
}

type UncaughtException struct {
}

func Try(f func()) *Tryblock {
	return &Tryblock{try: f}
}

func (t *Tryblock) Catch(f interface{}) *Tryblock {
	t.catchers = append(t.catchers, f)
	return t
}

func convertTo(i interface{}, to reflect.Type) (interface{}, bool) {
	from := reflect.TypeOf(i)
	if from.ConvertibleTo(to) {
		return reflect.ValueOf(i).Convert(to).Interface(), true
	}
	return nil, false
}

func callFunction(f interface{}, arg interface{}) bool {
	fval := reflect.ValueOf(f)

	if fval.Type().NumIn() != 1 {
		return false
	}

	if arg, ok := convertTo(arg, fval.Type().In(0)); ok {
		argVals := []reflect.Value{reflect.ValueOf(arg)}
		fval.Call(argVals)
		return true
	}

	return false
}

func (t *Tryblock) callCatcher(e interface{}) (called bool) {
	for i := 0; i < len(t.catchers) && !called; i++ {
		called = callFunction(t.catchers[i], e)
	}
	return
}

func (t *Tryblock) Finally(finally func()) {

	inFinally := false
	defer func() {
		if e := recover(); e != nil {
			called := false
			if !inFinally {
				called = t.callCatcher(e)
				finally()
			}

			if !called {
				panic(e)
			}
		}
	}()

	t.try()
	inFinally = true
	finally()
}
