package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	output []byte
}

func main() {
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
		res *result
	)

	resultChan = make(chan *result, 100)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err error
			cmd *exec.Cmd
		)

		cmd = exec.CommandContext(ctx,"C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2; echo hello")
		output, err = cmd.CombinedOutput()
		resultChan <- &result{
			err: err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)
	cancelFunc()
	res = <- resultChan
	fmt.Println(res.err, string(res.output))
}
