package model

import (
	"context"
)

const (
	REFRESH_TYPE AuthType = "refreshToken"
	CLIENT_TYPE  AuthType = "client"
	DEVICE_TYPE  AuthType = "device"
	NONE_TYPE    AuthType = "none"
	IRIS         Stage    = "IRIS"
	NORTHPOLE    Stage    = "NORTHPOLE"
	ALSTROEMERIA Stage    = "ALSTROEMERIA"
	AWS_SNOWDROP Stage    = "AWS_SNOWDROP"
	MARIGOLD     Stage    = "MARIGOLD"
	LUPINUS      Stage    = "LUPINUS"
	VIOLA        Stage    = "VIOLA"
	PHYSALIS     Stage    = "PHYSALIS"
)

type (
	AuthType    = string
	AccessToken = string
	Scope       = string
	TokenType   = string
	Stage       = string
	ExpiresIn   = int
	RequestCtx  context.Context
)
