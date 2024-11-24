package tools

import (
	"fmt"
	"os"
	"time"

	"github.com/go-mail/mail"
)

func SendEmail(to, subject, content string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("NOREPLY_MAIL_ADDRESS"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", fmt.Sprintf(MailTemplate, content, time.Now().Year()))

	d := mail.NewDialer("smtp.gmail.com",
		465,
		os.Getenv("NOREPLY_MAIL_ADDRESS"),
		os.Getenv("NOREPLY_MAIL_PASSWD"),
	)
	return d.DialAndSend(m)
}

var MailTemplate = `
<!doctype html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta
			name="viewport"
			content="width=device-width, initial-scale=1.0"
		/>
		<style>
			.mail-body {
				font-family: Arial, sans-serif;
				margin: auto;
				padding: auto;
				color: #fff;
				background-color: #000;
				max-width: 70%%;
			}
			.header {
				padding: 20px;
				text-align: center;
			}
			.header h1 {
				color: #fff;
				font-size: 24px;
				margin: 0;
			}
			.header p {
				color: #c7c7c7;
				font-size: 14px;
				margin-top: 5px;
			}
			.hero {
				padding: 40px;
				text-align: center;
				background-color: #000;
				
			}
			.hero h2 {
				font-size: 32px;
				color: #fff;
				margin-bottom: 10px;
			}
			.hero p {
				font-size: 18px;
				color: #e5e5e5;
				margin-bottom: 20px;
			}

			.cta-button {
				background-color: #d8232a;
				color: #fff !important;
				padding: 10px 30px;
				text-decoration: none;
				font-size: 14px;
				border-radius: 8px;
				display: inline-block;
			}
			.content {
				padding: 20px;
				background-color: #1c1c1c;
			}
			.content h3 {
				color: #fff;
				font-size: 22px;
				margin-top: 0;
			}
			.content p {
				color: #c7c7c7;
				font-size: 16px;
				line-height: 1.6;
			}
			.footer {
				background-color: #1c1c1c;
				padding: 20px;
				text-align: center;
				font-size: 14px;
				color: #7f7f7f;
			}
			.footer a {
				color: #d8232a;
				text-decoration: none;
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
