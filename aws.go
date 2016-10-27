package main

import (
    "bytes"
    "io/ioutil"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

var bucketName string
var projectData []byte
var output *s3.PutObjectOutput
var err error

func (u *Uploader)UploadData()(out string, e error) {
    if *development {
        out = "upload to s3 is disabled in development mode"
        return
    }

    if projectData, err = ioutil.ReadFile(u.tmpFile.Name()); err != nil {
        e = err
        return
    }

    if bucketName, err = u.BucketName(); err != nil {
        e = err
        return
    }

    svc := s3.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})
    params := &s3.PutObjectInput{
        Bucket:             aws.String( bucketName ),
        Key:                aws.String( u.Project.Key() ),
        Body:               bytes.NewReader( projectData ),
    }

    output, e = svc.PutObject(params)
    out = output.String()

    return
}
