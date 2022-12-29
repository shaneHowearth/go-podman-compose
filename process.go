package podmancomposer

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Input - received from user.
type Input struct {
	Files         []string
	Project       string
	Profile       string
	Verbose       bool
	LogLevel      string
	Ansi          bool
	Version       bool
	Host          string
	TLS           bool
	TLSCACert     string
	TLSCert       string
	TLSKey        string
	TLSVerify     bool
	SkipHostname  bool
	ProjectDir    string
	Compatibility bool
	Args          []string
}

func ProcessCommands(input Input) error {
	const OwnerRW = 0600

	for _, composeFilename := range input.Files {
		// Ensure that the file exists.
		if _, err := os.Stat(composeFilename); errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("%q cannot be found", composeFilename)
		}

		// Read it into memory.
		file, err := os.OpenFile(composeFilename, os.O_RDONLY, OwnerRW)
		if errors.Is(err, fs.ErrPermission) {
			return fmt.Errorf("do not have permissions to read %q", composeFilename)
		}

		_ = file // TODO
	}

	return nil
}
