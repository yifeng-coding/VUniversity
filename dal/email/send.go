package email

import (
	"context"
	"github.com/yifeng-coding/VUniversity/config"
	"gopkg.in/gomail.v2"
	"os"
)

// Send 发送邮件
func Send(ctx context.Context, to, subject, body string) error {
	emailConf := config.GetConfig().Email
	if len(emailConf.Username) == 0 {
		emailConf.Username = os.Getenv("Vuniversity_email_username")
		if len(emailConf.Username) == 0 {
			panic("Error to get email username")
		}
	}
	if len(emailConf.Password) == 0 {
		emailConf.Password = os.Getenv("Vuniversity_email_password") //
		if len(emailConf.Password) == 0 {
			panic("Error to get email password")
		}
	}

	// 构造邮件
	msg := gomail.NewMessage()
	msg.SetHeaders(map[string][]string{
		"From":    {emailConf.Username},
		"To":      {to},
		"Subject": {subject},
	})
	msg.SetBody("text/plain", body)
	// 发送邮件
	dialer := gomail.NewDialer(emailConf.Host, emailConf.Port, emailConf.Username, emailConf.Password)
	return dialer.DialAndSend(msg)
}
