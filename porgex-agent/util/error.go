package util

import "fmt"

func ErrorCheck(e error) {
	if e != nil {
		fmt.Errorf("App Get Memory Error", e)
		panic(e)
	}
}
