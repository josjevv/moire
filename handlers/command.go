package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func executeRaw(cmd *exec.Cmd) error {
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)

	return err
}
