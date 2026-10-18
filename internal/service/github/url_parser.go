package github

import (
	"net/url"
	"strings"

	custom_errors "github.com/slinkeres/open-source-project-analyzer/internal/errors"
)

func ParseUrlToOwnerAndName(rawUrl string) (owner, name, normalizedUrl string, err error) {
	parse, err := url.Parse(rawUrl)
	if err != nil {
		return "", "", "", err
	}
	if parse.Host != "github.com" {
		return "", "", "", custom_errors.ErrNotSupported
	}
	segments := []string{}
	for _, s := range strings.Split(parse.Path, "/") {
		if s != "" {
			segments = append(segments, s)
		}
	}
	if len(segments) < 2 {
		return "", "", "", custom_errors.ErrInvalidUrl
	}
	owner = segments[0]
	name = strings.TrimSuffix(segments[1], ".git")
	normalizedUrl = "https://github.com/" + owner + "/" + name
	return
}
