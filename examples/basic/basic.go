package main

import (
	"errors"
	"fmt"

	"github.com/snwfdhmp/rport"
)

var (
	r = rport.NewReporter("http://localhost:9900/")
)

func main() {
	someBigFunc()
}

func someFunc() {
	fmt.Println("Hello world !")
	return
}

func someBigFunc() {
	someWeirdFunc()
	someFunc()

	data := map[string]interface{}{
		"userId":    "someIdInfo",
		"requestId": 124,
		"browser":   "Chrome",
	}

	err := someNastyFunc()
	if err != nil {
		r.Report(err, "nasty-func-bug", data)
		return
	}

	// r.Track(someNastyFunc(), "nasty-func-bug", data)
}

func someWeirdFunc() {
	return
}

func someNastyFunc() error {
	return errors.New("I'm nasty.")
}
