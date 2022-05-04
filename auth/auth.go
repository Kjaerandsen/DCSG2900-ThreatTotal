package auth

import (
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"time"

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
	var oauth2Token *oauth2.Token
	// Temporary feide testing
	oauth2Token, err := utils.Config.Exchange(utils.Ctx, code)
	if err != nil {
		fmt.Println("Failed to exchange token: " + err.Error())
		return "", false
	}
	fmt.Println(oauth2Token)

	var test map[string]interface{}

	test2, _ := json.Marshal(oauth2Token)

	_ = json.Unmarshal(test2, &test)

	fmt.Println(test)

	fmt.Println("Extra: ", oauth2Token.Extra("id_token"))
	fmt.Println("Extra: ", oauth2Token.WithExtra("id_token"))
	fmt.Println("Extra: ", oauth2Token.WithExtra("userInfo"))

	marshalledToken, err := json.Marshal(oauth2Token)
	if err != nil {
		fmt.Println("Error marshalling oauth2 token in CodeToToken")
		return "", false
	}

	fmt.Println("Expiry time unformatted: ", oauth2Token.Expiry)
	fmt.Println("Time now: ", time.Now())
	fmt.Println("Time diff: ", oauth2Token.Expiry.Unix()-time.Now().Unix())

	// Add to the database
	response, err := utils.Conn.Do("SETEX", oauth2Token.AccessToken, (oauth2Token.Expiry.Unix() - time.Now().Unix()), marshalledToken)
	if err != nil {
		fmt.Println("Error adding data to redis:" + err.Error())
		return "", false
	}

	fmt.Println(response)

	// Uses the old token to get the userinfo again if expired
	/*
		userInfo, err := utils.Provider.UserInfo(utils.Ctx, oauth2.StaticTokenSource(oauth2Token))
		if err != nil {
			fmt.Println("Failed to get userinfo: " + err.Error())
			return "", false
		}

		// Still need a function to check expiration of the access-token somehow

		fmt.Println(oauth2.StaticTokenSource(oauth2Token).Token())

		fmt.Println(userInfo)
	*/
	//

	// Make the token a FeideToken struct

	// Return it

	return "", true
}

// Checks if a token is valid, returns a bool
func checkAuth(token string) (authenticated bool) {
	value, err := utils.Conn.Do("GET", token)
	if value == nil {
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
		fmt.Println("No Cache hit")
		return false
	} else {
		// Check if the key is valid first
		fmt.Println(value)
		fmt.Println(fmt.Sprintf("%v", value))
		fmt.Println(string([]byte(fmt.Sprintf("%v", value))))
		// Then if the jwt is valid

		// If not make a new jwt request to feide and replace the old jwt

		return true
	}
}

// Func which takes an authentication token and returns a hash.
func tokenToHash(code string) (hash string) {

	return ""
}
