package models

type EmailDto struct {
	To      string
	Subject string
	Body    string
}

type S3Dto struct {
	FileName string
	Key      string
}
