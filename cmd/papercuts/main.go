// Command papercuts appends a timestamped friction log entry to PAPERCUTS.md
// at the root of the current git repository.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var version = "dev"

const usage = `papercuts - log a small friction encountered while working in this repo

Usage:
  papercuts -m <model> "message describing what got in the way"

Flags:
  -m string   model/agent identifier (required)
  -v          print version and exit
`

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "papercuts:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("papercuts", flag.ContinueOnError)
	fs.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	model := fs.String("m", "", "model/agent identifier")
	showVersion := fs.Bool("v", false, "print version and exit")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if *showVersion {
		fmt.Println(version)
		return nil
	}

	if *model == "" {
		fs.Usage()
		return fmt.Errorf("-m <model> is required")
	}

	rest := fs.Args()
	if len(rest) != 1 || strings.TrimSpace(rest[0]) == "" {
		fs.Usage()
		return fmt.Errorf("a message is required")
	}
	message := strings.TrimSpace(rest[0])

	root, err := repoRoot()
	if err != nil {
		return err
	}

	author := gitUserName()
	if author == "" {
		author = currentUser()
	}
	if author == "" {
		author = "unknown"
	}

	return appendEntry(filepath.Join(root, "PAPERCUTS.md"), *model, author, message)
}

// repoRoot walks up from the current directory looking for a .git entry,
// mirroring how git itself locates the repo root. Falls back to cwd.
func repoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting working directory: %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return os.Getwd()
		}
		dir = parent
	}
}

func gitUserName() string {
	out, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func currentUser() string {
	if u := os.Getenv("USER"); u != "" {
		return u
	}
	return os.Getenv("USERNAME")
}

const fileHeader = "# Papercuts\n\nSmall frictions logged by agents while working in this repo: dead-end tool calls, broken links, misleading errors, undocumented setup steps. Not accomplishments and not tracked bugs, just the sand in the gears. Logged via `papercuts`.\n\n"

func appendEntry(path, model, author, message string) error {
	_, statErr := os.Stat(path)
	needsHeader := os.IsNotExist(statErr)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("opening %s: %w", path, err)
	}
	defer f.Close()

	var b strings.Builder
	if needsHeader {
		b.WriteString(fileHeader)
	}
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	fmt.Fprintf(&b, "%s - %s - %s\n\n%s\n\n", timestamp, model, author, message)

	if _, err := f.WriteString(b.String()); err != nil {
		return fmt.Errorf("writing to %s: %w", path, err)
	}
	return nil
}
