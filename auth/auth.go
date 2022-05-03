package auth

import (
	"dcsg2900-threattotal/utils"
	"fmt"

	"golang.org/x/oauth2"
)

// Authenticate function, takes a code or a token,
// returns a bool, and if the input is a valid code a hash is also returned.
func Authenticate(code string, token string) (authenticated bool, hash string) {
	authenticated = false
	var err bool
	if code != "" {
		hash, err = addUser(code)
		if !err {
			return
		}
		authenticated = true
		return
	} else if token != "" {
		authenticated = checkAuth(token)
		return
	}
	return
}

// Func which adds a user to the database and returns a code
func addUser(code string) (hash string, auth bool) {
	tokenResponse, auth := CodeToToken(code)
	if !auth {
		return "", false
	}
	hash = tokenToHash(tokenResponse)
	// Add the hash to the database with tokenResponse as the value

	return hash, true
}

// Func which takes a code and returns an authentication token.
// Inspiration from https://github.com/coreos/go-oidc/blob/v3/example/userinfo/app.go
func CodeToToken(code string) (token string, authenticated bool) {
	// Temporary feide testing
	oauth2Token, err := utils.Config.Exchange(utils.Ctx, code)
	if err != nil {
		fmt.Println("Failed to exchange token: " + err.Error())
		return "", false
	}
	fmt.Println(oauth2Token)

	// Uses the old token to get the userinfo again if expired
	userInfo, err := utils.Provider.UserInfo(utils.Ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		fmt.Println("Failed to get userinfo: " + err.Error())
		return "", false
	}

	// Still need a function to check expiration of the access-token somehow

	fmt.Println(oauth2.StaticTokenSource(oauth2Token).Token())

	fmt.Println(userInfo)

	//

	// Make the token a FeideToken struct

	// Return it

	return "", true
}

// Checks if a token is valid, returns a bool
func checkAuth(token string) (authenticated bool) {

	return true
}

// Func which takes an authentication token and returns a hash.
func tokenToHash(code string) (hash string) {

	return ""
}
