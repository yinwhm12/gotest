package main

import (
	"net/smtp"
	"strings"
	"fmt"
)

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := "xxxx@163.com"
	password := "xxxxxxx"
	host := "smtp.163.com:25"
	to := "xxxx@qq.com"

	subject := "Test send email by golang"

	//body := `
    //<html>
    //<body>
    //<h3>
    //"这是GO语言写的测试邮件。"
    //</h3>
    //</body>
    //</html>
    //`

	body := "hello my name is"
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}

}
