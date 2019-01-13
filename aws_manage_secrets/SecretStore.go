package main

import (
	"os"
	"io/ioutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
)

func encryptSecret(name string, value string) (string, error) {

	kmsKeyARN := "arn:aws:kms:us-east-1:794373491471:key/93d16842-a8b8-4626-a5e5-180a220e3770"
	kmsClient := kms.New(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	params := &kms.EncryptInput{
		KeyId:     aws.String(kmsKeyARN),
		Plaintext: []byte(value),
	}

	resp, err := kmsClient.Encrypt(params)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(name, resp.CiphertextBlob, 0644)
	if err != nil {
		return "", err
	}

	return name, nil
}

func uploadSecret(secretFileName string) error {

	s3Uploader := s3manager.NewUploader(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	reader, err := os.Open(secretFileName)
	if err != nil {
		return err
	}
	defer reader.Close()

	input := &s3manager.UploadInput{
		Bucket:           aws.String("cp-soft-secrets"),
		Key:              aws.String(secretFileName),
		Body:             reader,
	}
	_, err = s3Uploader.Upload(input)
	if err != nil {
		return err
	}

	err = os.Remove(secretFileName)
	if err != nil {
		return err
	}

	return nil
}

func downloadSecret(secretFileName string) (string, error) {

	s3Downloader := s3manager.NewDownloader(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	f, err := os.Create(secretFileName)
	if err != nil {
		return "", err
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String("cp-soft-secrets"),
		Key:    aws.String(secretFileName),
	}
	_, err = s3Downloader.Download(f, input)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func decryptSecretFile(secretFile string) (string, error) {

	secretBytes, err := ioutil.ReadFile(secretFile)
	if err != nil {
		return "", err
	}

	kmsClient := kms.New(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	params := &kms.DecryptInput{
		CiphertextBlob: secretBytes,
	}

	resp, err := kmsClient.Decrypt(params)
	if err != nil {
		return "", err
	}

	err = os.Remove(secretFile)
	if err != nil {
		return "", err
	}

	return string(resp.Plaintext), nil
}

func main() {
	secretName := "INFONAME"
	secret := "My Secret Information"

	secretFilename, err := encryptSecret(secretName, secret)
	if err != nil {
		log.Fatal("Failed to Encrypt data", err.Error())
	}

	err = uploadSecret(secretFilename)
	if err != nil {
		log.Fatal("Failed to Upload secret file to S3")
	}

	secretFilename2, err := downloadSecret(secretFilename)
	if err != nil {
		log.Fatal("Failed to Download secret file to S3")
	}

	plainText, err := decryptSecretFile(secretFilename2)
	if err != nil {
		log.Fatal("Failed to Decrypt secret")
	}
	log.Println("Secret is:", plainText)
}
