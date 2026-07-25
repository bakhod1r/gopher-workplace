package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go/format"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// handleFmt mirrors the UI Format button. It formats the submitted source with
// go/format (gofmt) and returns {ok, source, error} — the same shape the wasm
// window.gopherFormat returns.
func (s *server) handleFmt(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Src string `json:"src"`
	}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	out, err := format.Source([]byte(body.Src))
	if err != nil {
		writeJSON(w, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	writeJSON(w, map[string]any{"ok": true, "source": string(out)})
}

// handleVet materializes the challenge like /run and runs `go vet`, returning a
// report ({ok, compileOk, error}). No per-test cases — vet is pass/fail.
func (s *server) handleVet(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRun(r)
	if err != nil {
		writeJSON(w, report{Error: err.Error()})
		return
	}
	chDir, err := s.challengeDir(req.ChallengeID)
	if err != nil {
		writeJSON(w, report{Error: err.Error()})
		return
	}
	tmp, err := os.MkdirTemp("", "gw-vet-*")
	if err != nil {
		writeJSON(w, report{Error: "temp dir: " + err.Error()})
		return
	}
	defer os.RemoveAll(tmp)
	if err := materialize(chDir, tmp, req); err != nil {
		writeJSON(w, report{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), runTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "vet", "./...")
	cmd.Dir = tmp
	cmd.Env = sandboxEnv(tmp)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		if cmd.Process != nil {
			_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return cmd.Process.Kill()
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	runErr := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		writeJSON(w, report{CompileOK: true, Error: "vet timed out"})
		return
	}
	msg := strings.TrimSpace(out.String())
	if runErr != nil {
		writeJSON(w, report{CompileOK: false, Error: msg})
		return
	}
	writeJSON(w, report{OK: true, CompileOK: true, Cases: []caseResult{{Name: "go vet", Pass: true}}})
}

// handleHistory returns recent submissions for a challenge (most recent first).
func (s *server) handleHistory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("challengeId")
	if id == "" || s.store == nil {
		writeJSON(w, []any{})
		return
	}
	rows, err := s.store.history(id, 20)
	if err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, rows)
}

func decodeBody(r *http.Request, v any) error {
	if r.Method != http.MethodPost {
		return errors.New("POST required")
	}
	return json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(v)
}
