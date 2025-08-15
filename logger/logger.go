package logger

import (
	"fmt"
	"os"
	"os/exec"
)

// it should log both success and failed cases.
// success case should print Test Passed !

func FailedLog(cmd *exec.Cmd, err error) {
	file, e := os.Create("testFailure.log")
	if e != nil {
		fmt.Printf("Error creating log file %v", e)
	}

	file.WriteString(err.Error())

	fmt.Printf("Test Failed %v: \n Log Created : testFailure.log", cmd.Args)

	file.Close()
}

func SuccessLog(cmd *exec.Cmd, stdout []byte) {
	fmt.Printf("Test Passed %v: \n", cmd.Args)
}
