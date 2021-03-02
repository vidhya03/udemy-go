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

	if resp.Body != nil {

		lw := logWritter{}
		rd := textFileReader{}
		wtn, err := io.Copy(lw, rd)

		if err != nil {
			fmt.Println("Error :", err)
			os.Exit(-1)
		}
		fmt.Println(wtn)
	}

}

func (logWritter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("The total length is ", len(bs))
	return len(bs), nil
}

type textFileReader struct{}

func (textFileReader) Read(bs []byte) (int, error) {
	return 1, nil
}
