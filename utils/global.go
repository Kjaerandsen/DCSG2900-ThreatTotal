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
var Config oauth2.Config
