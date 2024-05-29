package main

import (
	"io"
	"os"
	"strings"
	"testing"

)

func Test_printMessage(t *testing.T) {

	stdOut := os.Stdout
	r,w,_ := os.Pipe()
	os.Stdout = w
	msg = "epsilon"

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output,"epsilon") {
             t.Error("Expected to find epsilon, but it is not there")
	}



}
