package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

var (
	err error
	out []byte
)

func main() {
	// prepare a buffer, to which the PID will be written
	var stdout bytes.Buffer

	// prepare the command
	sleepCmd := exec.Command("midprocrunner", "-cmd='sleep'", "-args='30'")
	sleepCmd.Stdout = &stdout

	// run the command
	err := sleepCmd.Run()
	if nil != err {
		panic(err)
	}

	// convert the PID string to a valid integer
	pidInt, err := strconv.ParseInt(stdout.String(), 10, 64)
	if nil != err {
		panic(err)
	}
	pid := int(pidInt)
	fmt.Printf("Created a detached process with an ID of %d!\n", pid)

	for {

	}
}
