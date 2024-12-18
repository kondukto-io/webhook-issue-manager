package model

type AttachmentReqArray struct {
	AttachmentReq []AttachmentReq `json:"attachments"`
}

type AttachmentReq struct {
	UUID          string `json:"id" gorm:"primaryKey;autoIncrement"`
	IssueID       string `json:"issueID"`
	Title         string `json:"title"`
	Base64Content string `json:"base64_content"`
}

type Attachment struct {
	ID       string
	IssueID  string
	Title    string
	FilePath string
}
