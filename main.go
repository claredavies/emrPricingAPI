package main

import (
    "fmt"
    "log"

    "emrPricingAPI/constants"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    // Replace these values with your own AWS credentials and region
    awsAccessKeyID := constants.AwsAccessKeyID
    awsSecretAccessKey := constants.AwsSecretAccessKey
    awsSessionToken := constants.AwsSessionToken
    awsRegion := "us-east-1" // Replace with your desired AWS region

    // Create AWS credentials
    creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, awsSessionToken)

    // Create a new AWS session with custom configuration
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(awsRegion),
        Credentials: creds,
    })
    if err != nil {
        log.Fatalf("Error creating AWS session: %v", err)
    }

    // Create an S3 service client using the custom session
    svc := s3.New(sess)

    // List S3 buckets
    result, err := svc.ListBuckets(nil)
    if err != nil {
        log.Fatalf("Error listing S3 buckets: %v", err)
    }

    // Print the names of the S3 buckets
    fmt.Println("S3 Buckets:")
    for _, b := range result.Buckets {
        fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
    }
}