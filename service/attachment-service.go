package service

import (
	"github.com/webhook-issue-manager/model"
	attachmentrepository "github.com/webhook-issue-manager/storage/attachment-repository"
)

var (
	attachmentrepo attachmentrepository.AttachmentRepository = attachmentrepository.NewAttachmentRepository()
)

type AttachmentService interface {
	CreateAttachment(attachment *model.Attachment) error
}

type attachmentservice struct{}

func NewAttachmentService() AttachmentService {
	return &attachmentservice{}
}

func (*attachmentservice) CreateAttachment(attachment *model.Attachment) error {
	err := attachmentrepo.AddAttachment(attachment)
	if err != nil {
		return err
	}
	return nil
}
