package attachmentrepository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/webhook-issue-manager/config"
	"github.com/webhook-issue-manager/model"
	"github.com/webhook-issue-manager/storage/postgres"
)

type AttachmentRepository interface {
	AddAttachment(attachmentReq *model.AttachmentReq) error
}

type attachmentrepository struct{}

func NewAttachmentRepository() AttachmentRepository {
	return &attachmentrepository{}
}

// AddAttachment implements AttachmentRepository
func (*attachmentrepository) AddAttachment(attachmentReq *model.AttachmentReq) error {

	ctx := context.Background()
	minioClient, err := config.MinioConnection()
	if err != nil {
		return errors.New("Error: " + err.Error())
	}

	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(attachmentReq.Base64Content)))
	base64.StdEncoding.Decode(base64Text, []byte(attachmentReq.Base64Content))

	permissions := 0644
	fileP := "C:/code/minio"
	file := fmt.Sprintf("%s/%s_%s.jpeg", fileP, attachmentReq.Title, attachmentReq.IssueId)
	err = ioutil.WriteFile(file, base64Text, fs.FileMode(permissions))
	if err != nil {
		return err
	}

	// Make a new bucket called mymusic.
	bucketName := "attachments"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
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

	objectName := attachmentReq.Title + "_" + attachmentReq.IssueId + ".jpeg"
	filePath := file
	contentType := "image/jpeg"

	_, err = minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	attachmentId, _ := uuid.NewRandom()

	attachment := model.Attachment{
		Id:       attachmentId.String(),
		IssueId:  attachmentReq.IssueId,
		Title:    attachmentReq.Title,
		FilePath: filePath,
	}

	db := postgres.Inıt()
	db.Create(&attachment)

	return nil
}
