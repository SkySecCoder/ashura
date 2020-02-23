package main

import (
	"fmt"
	//"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {
	fmt.Println("[+] Running AwsDestroyer...")
	creds := credentials.NewEnvCredentials()

	mysession, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: creds,
	})
	stsClient := sts.New(mysession)
	result, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
