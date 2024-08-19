package seto

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/nonylene/seto/src/common"
)

func execCode(cfg *Config, params *common.CodeParams) error {
	command := []string{"code"}
	cleanedPath := path.Clean(params.Path)

	if params.DevContainer {
		// https://github.com/microsoft/vscode-remote-release/issues/2133
		encoded := hex.EncodeToString([]byte(cleanedPath))
		remote := "dev-container+" + encoded
		if params.Remote {
			remote += "@" + cfg.CodeRemoteArgument
		}
		workspacePath := path.Join("/workspaces", path.Base(cleanedPath))
		// code {container path} --remote {contianer remote args}
		command = append(command, workspacePath, "--remote", remote)
	} else {
		// code {path} --remote {remote args}
		command = append(command, cleanedPath)
		if params.Remote {
			command = append(command, "--remote", cfg.CodeRemoteArgument)
		}
	}

	log.Printf("exec-code: %s", strings.Join(command, " "))
	out, err := exec.Command(command[0], command[1:]...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute code: %w\n%s", err, out)
	}

	return nil
}

func execBrowser(cfg *Config, params *common.BrowserParams) error {
	u, _ := url.Parse(params.Url)
	url := u.String()

	command := append(cfg.BrowserCommand, url)

	log.Printf("exec-browser: %s", strings.Join(command, " "))
	out, err := exec.Command(command[0], command[1:]...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute code: %w\n%s", err, out)
	}

	return nil
}

func Serve(cfg *Config) error {
	http.HandleFunc("POST /run/code", func(w http.ResponseWriter, r *http.Request) {
		var params common.CodeParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Printf("Faild to parse the request body: %+v", err)
			http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
			return
		}

		err = params.Validate()
		if err != nil {
			log.Printf("Faild to validate the request body: %+v", err)
			http.Error(w, "Failed to validate the request body", http.StatusBadRequest)
			return
		}

		err = execCode(cfg, &params)
		if err != nil {
			log.Printf("Faild to exec code: %+v", err)
			http.Error(w, "Failed to exec code", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("POST /run/browser", func(w http.ResponseWriter, r *http.Request) {
		var params common.BrowserParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Printf("Faild to parse the request body: %+v", err)
			http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
			return
		}

		err = params.Validate()
		if err != nil {
			log.Printf("Faild to validate the request body: %+v", err)
			http.Error(w, "Failed to validate the request body", http.StatusBadRequest)
			return
		}

		err = execBrowser(cfg, &params)
		if err != nil {
			log.Printf("Faild to exec code: %+v", err)
			http.Error(w, "Failed to exec code", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("GET /healthCheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	// Clean up existing socket; This may fail (e.g. the first run)
	syscall.Unlink(cfg.SocketPath)

	ln, err := net.Listen("unix", cfg.SocketPath)
	if err != nil {
		return fmt.Errorf("failed to listen the unix socket: %w", err)
	}
	defer ln.Close()

	err = os.Chmod(cfg.SocketPath, 0600)
	if err != nil {
		return fmt.Errorf("failed to set the unix socket permissions: %w", err)
	}

	return http.Serve(ln, nil)
}
