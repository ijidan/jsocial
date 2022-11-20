package repository

import (
	"fmt"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/mail"
)

type EmailRepository struct {
	manager *config.Manager
	conf    *config.Email
	email   *mail.Mail
}

func (e *EmailRepository) SendEmail(receiver map[string]string, cc map[string]string, subject string, content string) error {
	return e.email.Send(receiver, cc, subject, content)
}

func (e *EmailRepository) SendNewUserEmail(userId uint64, extra string) error {
	subject := "新用户注册"
	content := fmt.Sprintf("<p>新用户注册成功，用户Id：%d,附件信息：%s</p>", userId, extra)
	receiver := make(map[string]string)
	return e.email.Send(receiver, nil, subject, content)
}

func (e *EmailRepository) SendErrorEmail(err error) error {
	subject := "程序错误信息"
	content := fmt.Sprintf("<p>错误信息：%+v</p>", err)
	return e.email.Send(e.manager.Email, nil, subject, content)
}

func NewEmailRepository(manager *config.Manager, conf *config.Email) *EmailRepository {
	email := mail.NewMail(conf.Smtp, conf.Port, conf.Ssl, conf.Account, conf.Password)
	repo := &EmailRepository{
		manager: manager,
		conf:    conf,
		email:   email,
	}
	return repo
}
