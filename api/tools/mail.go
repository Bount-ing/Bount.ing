package tools

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-mail/mail"
)

func SendEmail(to, subject, content string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("NOREPLY_MAIL_ADDRESS"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	log.Printf("Sending mail to %s from %s", to, os.Getenv("NOREPLY_MAIL_ADDRESS"))
	m.SetBody("text/html", fmt.Sprintf(MailTemplate, content, time.Now().Year()))

	d := mail.NewDialer("smtp.gmail.com",
		465,
		os.Getenv("NOREPLY_MAIL_ADDRESS"),
		os.Getenv("NOREPLY_MAIL_PASSWD"),
	)
	return d.DialAndSend(m)
}

var MailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Template</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #000;
            color: #fff;
        }
        .mail-body {
            max-width: 600px;
            margin: auto;
            background-color: #1c1c1c;
            padding: 20px;
            border-radius: 8px;
        }
        .header {
            text-align: center;
            padding: 20px 0;
        }
        .header h1 {
            font-size: 26px;
            margin: 0;
            color: #008888;
        }
        .header p {
            font-size: 14px;
            color: #c7c7c7;
            margin-top: 5px;
        }
        .hero {
            text-align: center;
            padding: 30px;
        }
        .hero h2 {
            font-size: 28px;
            margin-bottom: 10px;
            color: #fff;
        }
		.hero h3 {
			font-size: 22px;
			margin-bottom: 10px;
			color: #fff;
		}
        .hero p {
            font-size: 18px;
            color: #e5e5e5;
            margin-bottom: 20px;
        }
        .cta-button {
            display: inline-block;
            background-color: #008888;
            color: #fff !important;
            padding: 12px 25px;
            font-size: 16px;
            border-radius: 5px;
            text-decoration: none;
            font-weight: bold;
        }
        .content {
            padding: 20px;
            text-align: left;
        }
        .content h3 {
            font-size: 22px;
            margin-bottom: 10px;
            color: #fff;
        }
        .content p {
            font-size: 16px;
            line-height: 1.6;
            color: #c7c7c7;
        }
        .footer {
            text-align: center;
            padding: 20px;
            font-size: 14px;
            color: #7f7f7f;
        }
        .footer a {
            color: #008888;
            text-decoration: none;
        }
        li strong {
            color: #fff;
        }
        li {
            color: #fff;
        }
    </style>
</head>
<body>
    <div class="mail-body">
        <div class="header">
            <h1>Bount.ing</h1>
            <p>OSS Bounties</p>
        </div>
        <div class="hero">%s</div>
        <div class="footer">
            <p>&copy; %d Bount.ing. All rights reserved.</p>
            <p>
                <a href="#">Unsubscribe</a> | 
                <a href="#">Privacy Policy</a>
            </p>
        </div>
    </div>
</body>
</html>
`
