package testlib1

import "fmt"

var globalMap map[string]string = make(map[string]string)

func Register(k, v string) {
	fmt.Println("testlib1@v0.0.1 Register")
	globalMap[k] = v
}

func GetAll() {
	fmt.Println("testlib1@v0.0.1 GetAll")
	fmt.Println(globalMap)
}
