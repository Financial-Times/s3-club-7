package main

import (
	"errors"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader struct {
	uuid string
	path string
	mode string
}

func (d *Downloader) BucketName() (b string, error error) {
	b = fmt.Sprintf("mio-project-%s", *clusterid)
	return
}

func (d *Downloader) DownloadProject () (tempFileName string, numBytes int64, error error) {
	tmpfile, error := ioutil.TempFile("", "s3downloading")
	if error != nil {
		log.Fatal(err)
		return
	}
	tempFileName = tmpfile.Name()
	defer tmpfile.Close()

	if bucketName, err = d.BucketName(); err != nil {
		error = err
		return
	}

	if key, err = d.Key(); err != nil {
		error = err
		return
	}

	downloader := s3manager.NewDownloader(session.New(&aws.Config{Region: aws.String("eu-west-1")}))
	numBytes, error = downloader.Download(tmpfile,
		&s3.GetObjectInput{
			Bucket: aws.String( bucketName ),
			Key:    aws.String( key ),
		})

	return
}


func (d *Downloader) Key() (key string, error error) {
	switch d.mode {
	case "ingest":
		key = fmt.Sprintf("%s/%s%s", d.uuid, d.path, ".plproj")
	case "publish":
		key = fmt.Sprintf("%s/%s%s", d.uuid, d.path, ".prproj")
	default:
		err = errors.New( fmt.Sprintf("mode: %s is not recognised", d.mode) )
		return
	}

	log.Printf( "Using key %s", key)
	return
}

func (d *Downloader) Filename() (key string, error error) {
	switch d.mode {
	case "ingest":
		key = fmt.Sprintf("%s%s", d.path, ".plproj")
	case "publish":
		key = fmt.Sprintf("%s%s", d.path, ".prproj")
	default:
		error = errors.New( fmt.Sprintf("mode: %s is not recognised", d.mode) )
		return
	}

	log.Printf( "Using key %s", key)
	return
}