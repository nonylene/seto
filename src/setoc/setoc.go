package setoc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"path/filepath"

	"github.com/nonylene/seto/src/common"
)

func request(cfg *Config, url string, requestBody []byte) error {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				// https://pkg.go.dev/net
				var d net.Dialer
				return d.DialContext(ctx, "unix", cfg.SocketPath)
			},
		},
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to post JSON: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code %d: %s", resp.StatusCode, responseBody)
	}

	return nil
}

func Browser(cfg *Config, url string) error {
	params := common.BrowserParams{
		Url: url,
	}
	log.Printf("Request: %+v\n", params)

	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return request(cfg, "http://localhost/run/browser", paramsBytes)
}

// path can be relative
func Code(cfg *Config, path string, devContainer bool) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve the path: %w", err)
	}

	params := common.CodeParams{
		Path:         abs,
		DevContainer: devContainer,
		Remote:       true,
	}
	log.Printf("Request: %+v\n", params)

	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return request(cfg, "http://localhost/run/code", paramsBytes)
}
