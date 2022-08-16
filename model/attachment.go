package model

type Attachment struct {
	Title         string `json:"title"`
	Base64Content string `json:"base64content"`
}
