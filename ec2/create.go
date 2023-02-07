package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Create an AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create a new EC2 client
	svc := ec2.New(sess)

	// Define the EC2 instance parameters
	instanceType := "t2.micro"
	imageID := "ami-0c55b159cbfafe1f0"
	minCount := int64(1)
	maxCount := int64(1)

	// Create the EC2 instance
	result, err := svc.RunInstances(&ec2.RunInstancesInput{
		InstanceType: &instanceType,
		ImageId:      &imageID,
		MinCount:     &minCount,
		MaxCount:     &maxCount,
	})

	if err != nil {
		println("Error creating EC2 instance: ", err)
		return
	}

	println("Successfully created EC2 instance: ", *result.Instances[0].InstanceId)
}
