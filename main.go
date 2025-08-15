package main

import (
	"os/exec"
	"sync"

	"github.com/momin-mostafa/run_my_tests/logger"
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
		logger.FailedLog(cmd, err)
	}

	logger.SuccessLog(cmd, stdout)
}
