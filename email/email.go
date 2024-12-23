package email

import "github.com/DusanKasan/parsemail"
import "net/mail"
import "time"
import "io"
import "fmt"

type Email struct {
	Header mail.Header

	Subject    string
	Sender     *mail.Address
	From       []*mail.Address
	ReplyTo    []*mail.Address
	To         []*mail.Address
	Cc         []*mail.Address
	Bcc        []*mail.Address
	Date       time.Time
	MessageID  string
	InReplyTo  []string
	References []string

	ResentFrom      []*mail.Address
	ResentSender    *mail.Address
	ResentTo        []*mail.Address
	ResentDate      time.Time
	ResentCc        []*mail.Address
	ResentBcc       []*mail.Address
	ResentMessageID string

	ContentType string
	Content     io.Reader

	HTMLBody string
	TextBody string

	Attachments   []Attachment
	EmbeddedFiles []EmbeddedFile
}

type Attachment struct {
	Filename    string
	ContentType string
	Data        io.Reader
}

type EmbeddedFile struct {
	CID         string
	ContentType string
	Data        io.Reader
}

func (e Email) String() string {
	return fmt.Sprintf("Subject: %v \nSender: %s  \nFrom: %s \nTo: %s", e.Subject, e.Sender, e.From, e.To)
}

func Parse(reader io.Reader) (email Email, err error) {

	e, err := parsemail.Parse(reader)
	if err != nil {
		return
	}

	email.Subject = e.Subject
	email.Sender = e.Sender
	email.From = e.From
	email.To = e.To
	email.HTMLBody = e.HTMLBody
	email.TextBody = e.TextBody
	return
}
