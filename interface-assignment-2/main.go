package main

import (
	"fmt"
	"io"
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
	openfile(os.Args[1])

}

func openfile(filename string) {
	f, err := os.Open(filename)
	check(err)

	io.Copy(os.Stdout, f)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
