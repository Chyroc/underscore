package httputils

import (
	"net/url"
)

func JoinURL(base, uri string) (string, error) {
	a, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	b, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	return a.ResolveReference(b).String(), nil
}
