package commands

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/h2non/filetype"
	"github.com/jozsefsallai/joe-cli/config"
	"github.com/urfave/cli"
)

func uploadFile(s *session.Session, c config.AWSConfig, f string) error {
	fmt.Println("Uploading file \"", f, "\"...")
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	info, _ := file.Stat()
	size := info.Size()

	buf := make([]byte, size)
	file.Read(buf)

	kind, _ := filetype.Match(buf)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(c.S3.Bucket),
		Key:           aws.String(filepath.Base(f)),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buf),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(kind.MIME.Value),
	})

	return err
}

func uploadCommand(ctx *cli.Context) {
	if ctx.NArg() != 1 {
		log.Fatal("Please provide a file to upload.")
	}

	file := ctx.Args().Get(0)

	conf := config.GetConfig()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.AWS.S3.Region),
		Credentials: credentials.NewStaticCredentials(
			conf.AWS.Key,
			conf.AWS.Secret,
			"",
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	c := conf.AWS

	err = uploadFile(s, c, strings.TrimSpace(file))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded successfully!")
}

// UploadCommand uploads a file to an S3 bucket (i.sallai.me)
var UploadCommand = cli.Command{
	Name:    "upload",
	Aliases: []string{"up"},
	Usage:   "Upload a file to an S3 bucket",
	Action:  uploadCommand,
}
