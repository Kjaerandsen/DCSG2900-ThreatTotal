package utils

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/oauth2"
)

var Conn redis.Conn
var Provider *oidc.Provider
var Ctx context.Context

// https://pkg.go.dev/golang.org/x/oauth2@v0.0.0-20200107190931-bf48bf16ab8d?utm_source=gopls#example-Config
// https://github.com/coreos/go-oidc/blob/v3/example/userinfo/app.go
var Config oauth2.Config
