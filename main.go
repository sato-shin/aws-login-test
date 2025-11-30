package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	out, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("failed to list buckets: %v", err)
	}

	if len(out.Buckets) == 0 {
		fmt.Println("No buckets found for the current credentials.")
		return
	}

	fmt.Println("Buckets:")
	for _, b := range out.Buckets {
		name := aws.ToString(b.Name)
		created := "unknown"
		if b.CreationDate != nil {
			created = b.CreationDate.In(time.Local).Format(time.RFC3339)
		}
		fmt.Printf("- %s (created %s)\n", name, created)
	}
}
