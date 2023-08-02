package oss

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"io"
	"log"
	"path/filepath"
	"time"
)

type S3 struct {
	URL    string
	Bucket string
	client *s3.Client
	ctx    context.Context
}

func NewS3(URL, Bucket, AwsAccessKeyId, AwsSecretAccessKe string) (*S3, error) {
	return NewS3Ctx(context.TODO(), URL, Bucket, AwsAccessKeyId, AwsSecretAccessKe)
}
func NewS3Ctx(ctx context.Context, URL, Bucket, AwsAccessKeyId, AwsSecretAccessKe string) (*S3, error) {
	var res S3
	res.ctx = ctx
	res.Bucket = Bucket
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:       "",
			URL:               URL,
			SigningRegion:     "",
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(res.ctx,
		config.WithRegion(""),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(AwsAccessKeyId, AwsSecretAccessKe, ""),
		),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalln("error:", err)
	}
	res.client = s3.NewFromConfig(cfg)
	_, err = res.bucketExists()
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Upload 上传文件, 可覆盖
func (s *S3) Upload(file io.Reader, key ...string) (*s3.PutObjectOutput, error) {
	result, err := s.client.PutObject(s.ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filepath.Join(key...)),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDownloadLink 获取文件下载链接
func (s *S3) GetDownloadLink(key ...string) (string, error) {
	presignClient := s3.NewPresignClient(s.client)
	presignedUrl, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(filepath.Join(key...)),
		},
		s3.WithPresignExpires(time.Minute*15))
	if err != nil {
		return "", err
	}
	return presignedUrl.URL, nil
}

// Delete 删除文件
func (s *S3) Delete(key ...string) (*s3.DeleteObjectOutput, error) {
	res, err := s.client.DeleteObject(s.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filepath.Join(key...)),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bucketExists 检查是否存在桶
func (s *S3) bucketExists() (bool, error) {
	_, err := s.client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(s.Bucket),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			var notFound *types.NotFound
			switch {
			case errors.As(apiError, &notFound):
				log.Printf("Bucket %v is available.\n", s.Bucket)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", s.Bucket, err)
			}
		}
	} else {
		log.Printf("Bucket %v exists and you already own it.", s.Bucket)
	}

	return exists, err
}

// FileExists 检查是否存在文件
func (s *S3) FileExists(key ...string) (bool, error) {
	_, err := s.client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filepath.Join(key...)),
	})

	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			var notFound *types.NotFound
			switch {
			case errors.As(apiError, &notFound):
				log.Printf("Bucket %v is available.\n", s.Bucket)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", s.Bucket, err)
			}
		}
	} else {
		log.Printf("Bucket %v exists and you already own it.", s.Bucket)
	}

	return exists, err
}