package helper

import (
	"golang-crud-gin/library/constants"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func IsErrorCodeEqualsSuccess(code string) bool {
	return code == constants.SUCCESS.Code
}
