package main

import "fmt"
import "errors"

func f(a int) (int, error) {
	if a > 3 {
		return 0, errors.New("To big to comprehend!")
	}
	return 0, nil
}

func main() {
	res, err := f(4)
	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		fmt.Println("res=", res)
	}
}
