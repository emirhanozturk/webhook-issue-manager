package model

type Attachment struct {
	Id            string `json:"id" gorm:"primaryKey;autoIncrement"`
	IssueId       string `json:"issueId"`
	Title         string `json:"title"`
	Base64Content string `json:"base64_content"`
}
