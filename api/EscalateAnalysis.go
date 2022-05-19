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

//Function linked to the escalation to manual analysis button in the frontend. Function sends email to user whom requested manual analysis.
//Function utlizes the gomail package.
//This function has been created with inspiration from https://www.loginradius.com/blog/engineering/sending-emails-with-golang/.
func EscalateAnalysis(url string, result string, token string, hash string) {

	email_pwd := os.Getenv("email_pwd") //Get service password from ENV.

	from := "threattotalv2@gmail.com" //Address to send email from.

	to := getUserEmail(token) //Gets the email of the user.

	m := gomail.NewMessage() //Create a new message.

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", "Analysis sucessfully escalated")

	var email_body string
	// Set E-Mail body. - The IF/Else checks if the email is an escalation of URL og File hash search
	if hash == "" {
		email_body = fmt.Sprintf("Your email has been escalated to manual analysis\n Details:\n URL: %s\n RequestLink: %s\n Do not reply to this email\n\n Further contact will be made from this email address", url, result)
	} else {
		email_body = fmt.Sprintf("Your email has been escalated to manual analysis\n Details:\n File hash: %s\n RequestLink: %s\n Do not reply to this email\n\n Further contact will be made from this email address", hash, result)
	}

	m.SetBody("text/plain", email_body) //Set body to type text.

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, from, email_pwd)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		logging.Logerror(err, "Error sending email - EscalateManualAnalysis.")
	}
}

//This function retrieves the user email from the redis caching solution.
func getUserEmail(hash string) (email string) {

	//fmt.Println("Hash for Redis req:", hash)

	value, err := utils.Conn.Do("GET", "user:"+hash) //Connect to the cache and query.
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
			logging.Logerror(err, "Error in cache lookup - Url-intelligence")

		}
	}
	responseBytes, err := json.Marshal(value) //Marshal data
	if err != nil {
		fmt.Println(err)
		logging.Logerror(err, "Error marshalling data")
	}

	var test []byte
	var JWTdata utils.IdAndJwt

	err = json.Unmarshal(responseBytes, &test) //Unmarshal data
	json.Unmarshal(test, &JWTdata)

	email = fmt.Sprintf("%s", JWTdata.Claims["email"]) //Set the email
	return email                                       //Return the email.
}
