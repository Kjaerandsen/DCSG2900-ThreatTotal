package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	// Internal
	"dcsg2900-threattotal/api"
	"dcsg2900-threattotal/auth"
	"dcsg2900-threattotal/storage"
	"dcsg2900-threattotal/utils"

	// External
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Initialize global variables
func init() {
	var err error

	utils.Ctx = context.Background()

	utils.Config = oauth2.Config{
		ClientID:     os.Getenv("clientId"),
		ClientSecret: os.Getenv("clientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.dataporten.no/oauth/authorization",
			TokenURL: "https://auth.dataporten.no/oauth/token",
		},
		RedirectURL: os.Getenv("feideRedirectUrl"),
		Scopes:      []string{oidc.ScopeOpenID, "email"},
	}

	// Initializing authentication connection
	utils.Provider, err = oidc.NewProvider(utils.Ctx, "https://auth.dataporten.no")
	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{
		ClientID: utils.Config.ClientID,
	}

	utils.Verifier = utils.Provider.Verifier(oidcConfig)

	RedisPool := storage.InitPool()
	utils.Conn = RedisPool.Get()

	// Get api keys as environment variables here
	utils.APIKeyVirusTotal = os.Getenv("APIKeyVirusTotal")
	utils.APIKeyGoogle = os.Getenv("APIKeyGoogle")
	utils.APIKeyHybridAnalysis = os.Getenv("APIKeyHybridAnalysis")
	utils.APIKeyOTX = os.Getenv("APIKeyOTX")

	utils.UrlBlockList = make([]string, 3)
	utils.UrlBlockList[0] = "ntnu.no"
	utils.UrlBlockList[1] = "ntnu.edu"
	utils.UrlBlockList[2] = "testsafebrowsing.appspot.com/s/malware.html"
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	// Login function which takes a code, uses it to retrieve a token and return a hash of the token to the user for
	// authentication.
	r.GET("/login", func(c *gin.Context) {
		code := c.Query("code")
		authenticated, hash := auth.Authenticate(code, "")
		if authenticated {
			fmt.Println("hash is: ", hash)
			c.JSON(http.StatusOK, gin.H{"hash": hash})
		} else {
			http.Error(c.Writer, "Failed authenticating with the code.", http.StatusUnauthorized)
		}
	})

	// Logout function which deletes a user from the database and sends a logout request to Feide.
	r.DELETE("/login", func(c *gin.Context) {
		hash := c.Query("userAuth")
		err := auth.Logout(hash)

		if !err {
			http.Error(c.Writer, "Failed logging out the user. Either the user was not logged in, or the login was expired.", http.StatusUnauthorized)
		} else {
			c.JSON(http.StatusOK, gin.H{"Logoutstatus": "Successfull"})
		}
	})

	// Auth function which authenticates a user using a hash of the token.
	r.GET("/auth", func(c *gin.Context) {
		auth2 := c.Query("auth")
		authenticated, _ := auth.Authenticate("", auth2)
		if authenticated {
			c.JSON(http.StatusOK, gin.H{"yes": "You are authenticated"})
		} else {
			http.Error(c.Writer, "Authentication is invalid or expired, please try to login again.", http.StatusUnauthorized)
		}
	})

	// Url intelligence takes a url and returns data on the url from our third-party sources.
	r.GET("/url-intelligence", func(c *gin.Context) {
		hash := c.Query("userAuth")
		authenticated, _ := auth.Authenticate("", hash)
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated": "You are not authenticated. User login is invalid."})
		} else {
			api.UrlIntelligence(c)
		}
	})

	// Hash intelligence takes a filehash and returns data on the file from our third-party sources.
	r.GET("/hash-intelligence", func(c *gin.Context) {

		hash := c.Query("userAuth")

		authenticated, _ := auth.Authenticate("", hash)
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated": "You are not authenticated. User login is invalid."})
		} else {
			api.HashIntelligence(c)
		}
	})

	// inspiration from https://github.com/dutchcoders/go-virustotal/blob/24cc8e6fa329f020c70a3b32330b5743f1ba7971/virustotal.go#L305
	// Upload function which allows the user to upload a file to virustotal, returns the id of the file in virustotal.
	r.POST("/upload", func(c *gin.Context) {

		hash := c.Query("userAuth")

		authenticated, _ := auth.Authenticate("", hash)
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated": "You are not authenticated. User login is invalid."})
		} else {
			api.UploadFile(c)
		}
	})

	// Upload retrieve function which takes a file id and returns the information on the file retrieved from virustotal.
	r.GET("/upload", func(c *gin.Context) {

		hash := c.Query("userAuth")

		authenticated, _ := auth.Authenticate("", hash)
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated": "You are not authenticated. User login is invalid."})
		} else {
			api.UploadFileRetrieve(c)
		}
	})

	r.GET("/escalate", func(c *gin.Context) {
		token := c.Query("userAuth")
		url := c.Query("url")
		result := c.Query("result")

		//api.EscalateAnalysis(url, result, token)

		authenticated, _ := auth.Authenticate("", token)
		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"authenticated": "You are not authenticated. User login is invalid."})
		} else {
			api.EscalateAnalysis(url, result, token)
			c.JSON(http.StatusOK, gin.H{"Successfull": "yes"})
		}

	})

	log.Fatal(r.Run(":8081"))
	// These don't do anything, and can't be placed above the line above as they stop the connections prematurely then.
	/*
		conn.Close()      // Close the connection
		redisPool.Close() // Close the pool
	*/
}
