package util

import "net/url"

type QueryMap = map[string]*string

func Query(_url *url.URL, args QueryMap) string {
	q := _url.Query()
	for k, v := range args {
		if v != nil && *v != "" {
			q.Set(k, *v)
		}
	}
	return q.Encode()
}
