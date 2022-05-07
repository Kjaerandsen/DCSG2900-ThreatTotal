package main

import (
	"context"
	"dcsg2900-threattotal/api"
	"dcsg2900-threattotal/auth"
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/storage"
	"dcsg2900-threattotal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	// External
	//webrisk "cloud.google.com/go/webrisk/apiv1"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	//"google.golang.org/grpc/status"
	//"google.golang.org/api/option"
	//webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1"
	//"google.golang.org/api/webrisk/v1"
	//"google.golang.org/api/option"
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
		RedirectURL: "http://localhost:3000",
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

		var wg sync.WaitGroup
		hash := c.Query("hash")

		var responseData [2]utils.FrontendResponse2

		var hybridApointer, AlienVaultpointer *utils.FrontendResponse2

		hybridApointer = &responseData[0]
		AlienVaultpointer = &responseData[1]

		wg.Add(2)
		go api.CallHybridAnalysisHash(hash, hybridApointer, &wg)
		go api.CallAlienVaultHash(hash, AlienVaultpointer, &wg)
		wg.Wait()

		var resultResponse utils.ResultFrontendResponse

		resultResponse.FrontendResponse = responseData[:]
		var resultPointer = &resultResponse

		utils.SetResultHash(resultPointer, len(responseData))

		Hashint, err := json.Marshal(resultResponse)
		if err != nil {
			fmt.Println(err)
			logging.Logerror(err)
			//c.Data(http.StatusInternalServerError, "application/json", )
		}

		//fmt.Println("WHERE IS MY CONTENT", responseData)

		c.Data(http.StatusOK, "application/json", Hashint)

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

	log.Fatal(r.Run(":8081"))
	// These don't do anything, and can't be placed above the line above as they stop the connections prematurely then.
	/*
		conn.Close()      // Close the connection
		redisPool.Close() // Close the pool
	*/
}
