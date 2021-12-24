package email

import (
	"gopkg.in/gomail.v2"
	_ "gopkg.in/gomail.v2"
	"ockham-api/config"
)

func SendEmail(mailTo []string, subject string, body string) error {
	emailConfig := config.GetConfig().Email

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(emailConfig.User, emailConfig.Sign))
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.User, emailConfig.Pass)

	err := d.DialAndSend(m)
	return err
}
