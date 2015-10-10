package meh

func SetOnError(errp *error) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*errp = err
		} else {
			panic(r)
		}
	}
}

func OnError(f func(err error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			f(err)
		} else {
			panic(r)
		}
	}
}

func ReturnError(err error) {
	if err != nil {
		panic(err)
	}
}
