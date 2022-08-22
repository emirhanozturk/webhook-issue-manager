package service

import (
	"github.com/webhook-issue-manager/model"
	attachmentrepository "github.com/webhook-issue-manager/storage/attachment-repository"
)

var (
	attachmentrepo attachmentrepository.AttachmentRepository = attachmentrepository.NewAttachmentRepository()
)

type AttachmentService interface {
	CreateAttachment(attachmentReq *model.AttachmentReq) error
}

type attachmentservice struct{}

func NewAttachmentService() AttachmentService {
	return &attachmentservice{}
}

func (*attachmentservice) CreateAttachment(attachmentReq *model.AttachmentReq) error {
	err := attachmentrepo.AddAttachment(attachmentReq)
	if err != nil {
		return err
	}
	return nil
}
