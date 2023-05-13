package testlib1

import "fmt"

var globalMap map[string]string = make(map[string]string)

const prefix = "testlib1@v2.0.2 "

func Register(k, v string) {
	fmt.Println(prefix + "Register")
	globalMap[k] = v
}

func GetAll() {
	fmt.Println(prefix + "GetAll")
	fmt.Println(globalMap)
}
