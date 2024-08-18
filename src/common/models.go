package common

import (
	"errors"
	"fmt"
	"net/url"
	"path"
)

type CodeParams struct {
	Path         string `json:"path"`
	DevContainer bool   `json:"devContainer"`
	Remote       bool   `json:"remote"`
}

func (c *CodeParams) Validate() error {
	if c.Path == "" {
		return errors.New("path must be defined")
	}

	cleanedPath := path.Clean(c.Path)

	if !path.IsAbs(cleanedPath) || cleanedPath == "/" {
		return errors.New("path must be absolute non-root path")
	}
	return nil
}

type BrowserParams struct {
	Url string `json:"url"`
}

func (b *BrowserParams) Validate() error {
	if b.Url == "" {
		return errors.New("url must be defined")
	}

	u, err := url.Parse(b.Url)
	if err != nil {
		return fmt.Errorf("failed to parse the request body as URL: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("only HTTP or HTTPS url is supported")
	}

	if u.Host == "" {
		return errors.New("URL does not have a host")
	}

	if u.Opaque != "" {
		return errors.New("invalid URL format")
	}

	return nil
}
