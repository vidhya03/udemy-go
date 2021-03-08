package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)
	fmt.Print(string(dat))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
