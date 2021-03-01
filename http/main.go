package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("https://vidhya03.labkit.in")

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(-1)
	}

	lw := logWritter{}
	wtn, err := io.Copy(lw, resp.Body)

	fmt.Println(wtn)

}

func (logWritter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("The total length is ", len(bs))
	return len(bs), nil
}
