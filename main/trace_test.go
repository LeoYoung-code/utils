package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// run the main function
	main()

	// open the trace file
	f, err := os.Open("trace.out")
	if err != nil {
		// fail if the file isn't created
		panic(err)
	}

	// read the file
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		// fail if the file isn't readable
		panic(err)
	}

	// fail if the file is empty
	if len(bytes) == 0 {
		panic("trace file is empty")
	}
}