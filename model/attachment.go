package model

type AttachmentReqArray struct {
	AttachmentReq []AttachmentReq `json:"attachments"`
}

type AttachmentReq struct {
	Id            string `json:"id" gorm:"primaryKey;autoIncrement"`
	IssueId       string `json:"issueId"`
	Title         string `json:"title"`
	Base64Content string `json:"base64content"`
}

type Attachment struct {
	Id       string
	IssueId  string
	Title    string
	FilePath string
}
