package auth

import (
	"dcsg2900-threattotal/utils"
	"fmt"
)

// Authenticate function, takes a code or a token,
// returns a bool, and if the input is a valid code a hash is also returned.
func Authenticate(code string, token string) (authenticated bool, hash string) {
	if code != "" {
		hash, err := addUser(code)
		if !err {
			return false, ""
		}
		return true, hash
	} else if token != "" {
		return checkAuth(token), ""
	}
	return false, ""
}

// Func which adds a user to the database and returns a code
func addUser(code string) (hash string, auth bool) {
	tokenResponse, auth := codeToToken(code)
	if !auth {
		return "", false
	}
	hash = tokenToHash(tokenResponse)
	// Add the hash to the database with tokenResponse as the value

	return hash, true
}

// Func which takes a code and returns an authentication token.
func codeToToken(code string) (token string, authenticated bool) {
	// Temporary feide testing
	oauth2Token, err := utils.Config.Exchange(utils.Ctx, code)
	if err != nil {
		fmt.Println("Failed to exchange token: " + err.Error())
		return
	}
	fmt.Println(oauth2Token)

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
