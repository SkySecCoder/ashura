package awsDestroyer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
	"strings"
)

func AwsDestroyer() {
	masterInstance := map[string][]string{}
	ch := make(chan string)
	fmt.Println("[+] Running AwsDestroyer as :")

	whoami()

	for index := range Regions {
		go getInstances(Regions[index], &masterInstance, ch)
	}
	for _ = range Regions {
		_ = (<-ch)
	}

	data, _ := json.MarshalIndent(masterInstance, "", "    ")
	fmt.Println("[-] Instances to be destroyed")
	fmt.Println(string(data))
	input := bufio.NewReader(os.Stdin)
	fmt.Print("[?] Are you sure you want to destroy these instances?(y/n): ")

	choice, _ := input.ReadString('\n')
	choice = strings.ReplaceAll(choice, "\n", "")

	if (choice == "y") || (choice == "Y") {
		fmt.Println("")
		for region, instanceList := range masterInstance {
			instanceDestroyer(region, instanceList)
		}
	}
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

func getInstances(region string, masterInstance *map[string][]string, ch chan<- string) {
	var instanceList []string

	mysession := createSession(region)
	ec2Client := ec2.New(mysession)

	result, err := ec2Client.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		panic(err)
	}

	for instance := range result.Reservations {
		blob := result.Reservations[instance].Instances
		instanceList = append(instanceList, *blob[0].InstanceId)
	}
	if instanceList != nil {
		(*masterInstance)[region] = instanceList
	}
	ch <- fmt.Sprintf("completed")
}

func instanceDestroyer(region string, instanceList []string) {
	mysession := createSession(region)
	ec2Client := ec2.New(mysession)
	data, err := ec2Client.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: aws.StringSlice(instanceList),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for index := range data.TerminatingInstances {
		fmt.Println("-> " + *data.TerminatingInstances[index].InstanceId + " is in the " + *data.TerminatingInstances[index].CurrentState.Name + " state")
	}

}

func createSession(region string) *session.Session {
	creds := credentials.NewEnvCredentials()

	mysession, _ := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})
	return mysession
}
