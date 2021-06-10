# Authorization

## How to set Token

```go
ctx := ciosctx.WithToken(context.Background(), "Bearer Hogehoge")
```

## How to Auto Refresh Auth

```go
client := ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    AutoRefresh:     true,
    AuthConfig:      ciossdk.RefreshTokenAuth("client_id", "client_secret", "refresh_token", "scope"),
})
```

## How to Auto Client Auth

```go
client := ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    AutoRefresh:     true,
    AuthConfig:      ciossdk.ClientAuthConf("client_id", "client_secret, "scope"),
})
```

## How to Auto Device Auth

```go
client := ciossdk.NewCiosClient(ciossdk.CiosClientConfig{
    AutoRefresh:     true,
    AuthConfig:      ciossdk.DeviceAuthConf("client_id", "client_secret", "assertion", "scope"),
})
```

## Get Token

### interface

```
GetAccessTokenByRefreshToken() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error)
GetAccessTokenOnClient() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error)
GetAccessTokenOnDevice() (sdkmodel.AccessToken, sdkmodel.Scope, sdkmodel.TokenType, sdkmodel.ExpiresIn, error)
SetClientSecret(clientSecret string)
SetClientId(clientId string)
SetRef(ref string)
SetAssertion(assertion string)
SetDebug(debug bool)
SetScope(scope string)
```

### Usage

#### Get an AccessToken by refresh token

※ Fix in the future
```go
token, scope, tokenType, expiresIn, err := client.Auth().GetAccessTokenByRefreshToken()
```

#### Get an AccessToken by client

※ Fix in the future
```go
token, scope, tokenType, expiresIn, err := client.Auth().GetAccessTokenOnClient()
```

#### Get an AccessToken by device 

※ Fix in the future
```go
token, scope, tokenType, expiresIn, err := client.Auth().GetAccessTokenOnDevice()
```
