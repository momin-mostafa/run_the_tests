package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go runTest(getUnitTestCmnd(), &wg)
	go runTest(getIntegrationTestCmnd(), &wg)

	wg.Wait()
}

func getUnitTestCmnd() *exec.Cmd {
	return exec.Command("flutter", "test", "test/unit_test/unit_test.dart")
}

func getIntegrationTestCmnd() *exec.Cmd {
	return exec.Command("flutter", "test", "integration_test/integration_test.dart")
}

func runTest(cmd *exec.Cmd, wg *sync.WaitGroup) {
	defer wg.Done()

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running %v: %v\n", cmd.Args, err)
	}

	fmt.Printf("Output from %v:\n%s\n", cmd.Args, stdout)
}
