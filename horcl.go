package horcl

import (
	"errors"
	"fmt"
	"log"
)

var (
	ORCL_ERROR = errors.New("99200:orcl->%s")
)

func OrclPanic(err error) {
	if err != nil {
		log.Panic(errors.New(fmt.Sprintf(ORCL_ERROR.Error(), err.Error())))
	}
}

func Insert() {

}

func Update() {

}

func Delete() {

}

func FindAll() {

}

func FindOne() {

}

func FindLimit() {

}

func Exists() {

}
