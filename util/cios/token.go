package ciosutil

import "strings"

func ParseAccessToken(accessToken string) string {
	if !strings.Contains(accessToken, "Bearer ") {
		return "Bearer " + accessToken
	} else {
		return accessToken
	}
}
