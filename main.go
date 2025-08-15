package main

import (
	"fmt"
	"os/exec"
)

func main() {
	runTest(getUnitTestCmnd())
	runTest(getIntegrationTestCmnd())
}

func getUnitTestCmnd() *exec.Cmd {
	return exec.Command("flutter", "test", "test/unit_test/unit_test.dart")
}

func getIntegrationTestCmnd() *exec.Cmd {
	return exec.Command("flutter", "test", "integration_test/integration_test.dart")
}

func runTest(cmd *exec.Cmd) {
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
