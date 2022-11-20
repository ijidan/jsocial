package repository

import (
	"fmt"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/mail"
)

func SendEmail(receiver map[string]string, cc map[string]string, subject string, content string) error {
	conf := global.GR.Conf.Email
	email := mail.NewMail(conf.Smtp, conf.Port, conf.Ssl, conf.Account, conf.Password)
	return email.Send(receiver, cc, subject, content)
}

func SendNewUserEmail(userId uint64, extra string) error {
	email := mail.NewMail(global.GR.Conf.Email.Smtp, global.GR.Conf.Email.Port, global.GR.Conf.Email.Ssl, global.GR.Conf.Email.Account, global.GR.Conf.Email.Password)
	subject := "新用户注册"
	content := fmt.Sprintf("<p>新用户注册成功，用户Id：%d,附件信息：%s</p>", userId, extra)
	return email.Send(global.GR.Conf.Manager.Email, nil, subject, content)
}

func SendErrorEmail(err error) error {
	conf := global.GR.Conf.Email
	email := mail.NewMail(conf.Smtp, conf.Port, conf.Ssl, conf.Account, conf.Password)
	subject := "程序错误信息"
	content := fmt.Sprintf("<p>错误信息：%+v</p>", err)
	return email.Send(global.GR.Conf.Manager.Email, nil, subject, content)
}
