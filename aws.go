package main

import (
    "io"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

var bucketName string

func (u *Uploader)UploadData(projectData io.Reader)(string, error) {
    svc := s3.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

    if bucketName, err := u.BucketName(); err != nil {
        return "", err
    }

    params := &s3.PutObjectInput{
        Bucket:             aws.String( bucketName ),
        Key:                aws.String( u.Project.Key() ),
        Body:               projectData,
    }

    output, err := svc.PutObject(params)
    return output.String(), err
}
