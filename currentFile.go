package cloedy

import (
    "runtime"
    "errors"
)

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}

func CurrentDir() string {
    return path.Dir(CurrentFile())
}
