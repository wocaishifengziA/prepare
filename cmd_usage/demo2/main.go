package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
		output []byte
	)

	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo 1")

	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("123", string(output))
}
