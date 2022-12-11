package scanner

import (
	log "github.com/sirupsen/logrus"

	"github.com/SkySecCoder/ashura/ashura/pkg/scanner/gitScanner"
	"github.com/SkySecCoder/ashura/ashura/pkg/scanner/awsScanner"
)

func ScannerHandler(scannerType string) {
	if scannerType == "aws" {
		awsScanner.AwsScannerHandler()
	} else if scannerType == "git" {
		gitScanner.GitScannerHandler()
	} else {
		log.Info("Incorrect flag for scanner was supplied\nFollowing are the supported scanner types:")
		log.Info("1. aws\n2. git")
	}
}
