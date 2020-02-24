package main

import (
	"fmt"
	"strings"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	fmt.Println("[+] Running AwsDestroyer as :")
	creds := credentials.NewEnvCredentials()

	mysession, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: creds,
	})
	stsClient := sts.New(mysession)
	result, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		panic(err)
	}

	Whoami(string(*result.Arn))
	GetInstances(mysession)
}

func Whoami(arn string) {
	user := strings.Split(arn, "/")[len(strings.Split(arn, "/"))-1]
	account := strings.Split(arn, ":")[4]
	fmt.Println("\t- User\t\t: "+user)
	fmt.Println("\t- Account\t: "+string(account))
	fmt.Println("\t- ARN\t\t: "+string(arn))
}

func GetInstances(mysession *session.Session) {
	var instanceList []string
	ec2Client := ec2.New(mysession)
	result, err := ec2Client.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		panic(err)
	}

	for instance := range result.Reservations {
		blob := result.Reservations[instance].Instances
		instanceList = append(instanceList, *blob[0].InstanceId)
	}
	data, _ := json.MarshalIndent(instanceList, "", "    ")
	fmt.Println(string(data))
}
