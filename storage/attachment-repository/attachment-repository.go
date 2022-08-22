package attachmentrepository

import (
	"context"
	"encoding/base64"
	"errors"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/webhook-issue-manager/config"
	"github.com/webhook-issue-manager/model"
)

var (
	bucketName = "mymusic"
	location   = "us-east-1"
)

type AttachmentRepository interface {
	AddAttachment(attachment *model.Attachment) error
}

type attachmentrepository struct{}

func NewAttachmentRepository() AttachmentRepository {
	return &attachmentrepository{}
}

// AddAttachment implements AttachmentRepository
func (*attachmentrepository) AddAttachment(attachment *model.Attachment) error {

	ctx := context.Background()
	minioClient, err := config.MinioConnection()
	if err != nil {
		return errors.New("Error: " + err.Error())
	}

	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(attachment.Base64Content)))
	base64.StdEncoding.Decode(base64Text, []byte(attachment.Base64Content))

	// Make a new bucket called mymusic.
	bucketName := "attachments"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	if err != nil {
		return err
	}
	return nil
}
