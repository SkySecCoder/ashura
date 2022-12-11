package awsScanner

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func AwsScannerHandler() {
	whoami()
}

func whoami() {
	mysession := createSession("us-east-1")
	stsClient := sts.New(mysession)

	result, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}
	arn := *result.Arn

	user := strings.Split(arn, "/")[len(strings.Split(arn, "/"))-1]
	account := strings.Split(arn, ":")[4]
	fmt.Println("\t- User\t\t: " + user)
	fmt.Println("\t- Account\t: " + string(account))
	fmt.Println("\t- ARN\t\t: " + string(arn))
}

func createSession(region string) *session.Session {
	creds := credentials.NewEnvCredentials()

	mysession, _ := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})
	return mysession
}
