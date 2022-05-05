package auth

import (
	"crypto/sha256"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Authenticate function, takes a code or a token,
// returns a bool, and if the input is a valid code a hash is also returned.
func Authenticate(code string, token string) (authenticated bool, hash string) {
	authenticated = false
	var err bool
	if code != "" {
		fmt.Println("Hash is not empty")
		hash, err = addUser(code)
		fmt.Println("hash is: ", hash)
		if !err {
			return
		}
		authenticated = true
		fmt.Println("Returning: ", authenticated, hash)
		return authenticated, hash
	} else if token != "" {
		_, authenticated = getAuth(token)
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
	//hash = tokenToHash(tokenResponse)
	// Add the hash to the database with tokenResponse as the value

	return tokenResponse, true
}

// Func which takes a code and returns an authentication token.
// Inspiration from the go-oidc examples: https://github.com/coreos/go-oidc/blob/v3/example/userinfo/app.go
// and https://github.com/coreos/go-oidc/blob/v3/example/idtoken/app.go
func CodeToToken(code string) (token string, authenticated bool) {
	// Get the token
	oauth2Token, err := utils.Config.Exchange(utils.Ctx, code)
	if err != nil {
		fmt.Println("Failed to exchange token: " + err.Error())
		return "", false
	}
	// Extra fields contain: scope, token_type and id_token

	// Get the jwt
	rawIDToken, error := oauth2Token.Extra("id_token").(string)
	if !error {
		fmt.Println("No jwt returned.")
		return "", false
	}

	// Verify the jwt
	idToken, err := utils.Verifier.Verify(utils.Ctx, rawIDToken)
	if err != nil {
		fmt.Println("Failed to validate the jwt.")
		return
	}

	var dataTest map[string]interface{}
	// Parse the userdata in the jwt
	idToken.Claims(&dataTest)
	fmt.Println("JWT claims: ", dataTest)

	var data utils.IdAndJwt
	data.Oauth2Token = *oauth2Token
	data.Jwt = *idToken
	data.Claims = dataTest

	marshalledTokens, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling tokens in CodeToToken")
		return "", false
	}

	hash := tokenToHash(oauth2Token.AccessToken)

	// Add to the database
	_, err = utils.Conn.Do("SETEX", hash, (oauth2Token.Expiry.Unix() - time.Now().Unix()), marshalledTokens)
	if err != nil {
		fmt.Println("Error adding data to redis:" + err.Error())
		return "", false
	}
	//fmt.Println(response)

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

	// If everything is successfull return true and the authentication code for the frontend user.
	return hash, true
}

func getAuth(hash string) (token string, err bool) {
	value, error := utils.Conn.Do("GET", hash)
	if value == nil {
		if error != nil {
			fmt.Println("Error:" + error.Error())
		}
		fmt.Println("No Cache hit")
		return "", false
	} else {

		var responseData utils.IdAndJwt
		error := json.Unmarshal(value.([]byte), &responseData)
		if error != nil {
			fmt.Println("Error unmarshalling")
			// If there is an error delete the key
			value, error := utils.Conn.Do("DEL", token)
			if error != nil {
				fmt.Println("Failed deleting key in redis: ", err)
			}
			fmt.Println("Redis delete response: ", value)
			return "", false
		}

		//fmt.Println("marhaslled data: ", responseData)

		// If email is needed a helper which checks the expiration of the jwt
		// and requests a new one is needed.

		return responseData.Oauth2Token.AccessToken, true
	}
}

// Func which takes an authentication token and returns a hash.
func tokenToHash(code string) (hash string) {
	// Create the sha256 hash
	fileHash := sha256.New()
	fileHash.Write([]byte(code))

	return fmt.Sprintf("%x", fileHash.Sum(nil))
}

func Logout(hash string) bool {
	// Get the login details from the database
	userToken, err := getAuth(hash)
	if !err {
		return false
	}

	// Delete the database item
	_, error := utils.Conn.Do("DEL", hash)
	if error != nil {
		fmt.Println("Error removing data from redis:" + error.Error())
		return false
	}

	// Send a logout request to feide
	resp, error := http.Get("https://auth.dataporten.no/openid/endsession?id_token_hint=" + userToken)
	if error != nil {
		fmt.Println("Error logging out from feide:" + error.Error() + ". HTTP status code: " + resp.Status)
		return false
	}

	// Return true if successfull
	return true
}
