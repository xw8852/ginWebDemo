package util

import (
	"github.com/pkg/errors"
	"fmt"
)


func Convert(err error) bool{
	if err!=nil{
		e := errors.New(err.Error())
		fmt.Printf("%+v", e)
		return true
	}
	return false
}