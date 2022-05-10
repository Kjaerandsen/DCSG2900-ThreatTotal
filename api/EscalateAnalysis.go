package api

import (
	"crypto/tls"
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

func EscalateAnalysis(url string, result string, token string) {

	email_pwd := os.Getenv("email_pwd")

	from := "threattotalv2@gmail.com"

	to := "pederas@stud.ntnu.no"

	coolstuff := getUserEmail(token)

	fmt.Println("After return", coolstuff)

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", "Analysis sucessfully escalated")

	// Set E-Mail body. You can set plain text or html with text/html
	email_body := fmt.Sprintf("Your email has been escalated to manual analysis\n Details:\n URL: %s\n RequestLink: %s\n Do not reply to this email\n\n Further contact will be made from this email address", url, result)

	m.SetBody("text/plain", email_body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, from, email_pwd)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func getUserEmail(hash string) (email string) {

	fmt.Println("Hash for Redis req:", hash)

	value, err := utils.Conn.Do("GET", "user:"+hash)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
			logging.Logerror(err, "Error in cache lookup - Url-intelligence")

		}
	}
	responseBytes, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}

	var test []byte
	var JWTdata utils.IdAndJwt

	fmt.Println(string(responseBytes))
	err = json.Unmarshal(responseBytes, &test)
	json.Unmarshal(test, &JWTdata)

	fmt.Println(test)
	fmt.Println(string(test))

	fmt.Println(JWTdata)
	fmt.Println(JWTdata.Claims["email"])

	email = fmt.Sprintf("%s", JWTdata.Claims["email"])
	return email
}
