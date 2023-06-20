package initializers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	awsRegion   string
	awsEndpoint string
	bucketName  string
	s3svc       *s3.Client
)

func init() {
	awsRegion = os.Getenv("AWS_REGION")
	awsEndpoint = os.Getenv("AWS_ENDPOINT")
	bucketName = os.Getenv("S3_BUCKET")

	// Definir as configurações padrão se as variáveis de ambiente não estiverem definidas
	if awsRegion == "" {
		awsRegion = "us-east-1"
	}
	if awsEndpoint == "" {
		awsEndpoint = "http://localhost:4566"
	}
	if bucketName == "" {
		bucketName = "app-files"
	}

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// Retornar um EndpointNotFoundError permitirá que o serviço use a resolução padrão
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações da AWS: %s", err)
	}

	s3svc = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}

func main() {
	err := CreateBucket(bucketName)
	if err != nil {
		log.Fatalf("Erro ao criar o bucket: %s", err)
	}

	_, err = ListBuckets()
	if err != nil {
		log.Fatalf("Erro ao listar os buckets: %s", err)
	}

	err = listObjects()
	if err != nil {
		log.Fatalf("Erro ao listar os objetos: %s", err)
	}
}

func CreateBucket(bucketName string) error {
	output, err := s3svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("erro ao criar o bucket: %w", err)
	}
	fmt.Println(output)

	return nil
}

func ListBuckets() ([]types.Bucket, error) {
	result, err := s3svc.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("erro ao listar os buckets: %w", err)
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}

	return result.Buckets, nil
}

func UploadFile(objectKey string, fileBytes []byte) error {
	bucketName := os.Getenv("S3_BUCKET")
	reader := bytes.NewReader(fileBytes)


	_, err := s3svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   reader,
	})

	if err != nil {
		log.Printf("Couldn't upload file to %v:%v. Here's why: %v\n", bucketName, objectKey, err)
	}

	return err
}


func listObjects() error {
	output, err := s3svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return fmt.Errorf("erro ao listar os objetos: %w", err)
	}

	fmt.Println("Lista de Objetos:")
	for _, object := range output.Contents {
		fmt.Printf("Chave: %s, Tamanho: %d\n", aws.ToString(object.Key), object.Size)
	}

	return nil
}
