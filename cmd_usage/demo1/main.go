package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	var cmd *exec.Cmd
	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo 1")
	err := cmd.Run()
	fmt.Println(err)
}
