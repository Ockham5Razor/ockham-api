package main

import (
	"gol-c/config"
	"gol-c/email"
	"testing"
)

func TestSendMail(t *testing.T) {
	envConfEmail := config.GetTestEnvConfig()["email_test"]
	sendTo := envConfEmail["sendTo"]
	subject := envConfEmail["subject"]
	body := envConfEmail["body"]
	err := email.SendEmail([]string{sendTo}, subject, body)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Skip()
	}
}
