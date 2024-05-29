package main

import (
	"io"
	"os"
	"strings"
	"testing"

)
func Test_updateMessage(t *testing.T){
	wg.Add(1)
	
	updateMessage("epsilon",&wg)
	wg.Wait()

	if msg != "epsilon" {

		t.Errorf("Expected to find epsilon, but it is not there,%v",msg)
	}



}


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
             t.Errorf("Expected to find epsilon, but it is not there %v",output)
	}



}
