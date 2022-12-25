package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pborman/getopt/v2"
)

// Input - received from user
type Input struct {
	Files         []string
	Project       string
	Profile       string
	Verbose       bool
	LogLevel      string
	Ansi          bool
	Version       bool
	Host          string
	Tls           bool
	TlsCACert     string
	TlsCert       string
	TlsKey        string
	TlsVerify     bool
	SkipHostname  bool
	ProjectDir    string
	Compatibility bool
	Args          []string
}

func main() {
	// Create flags for input
	helpFlag := getopt.BoolLong("help", 'h', "display help")
	files := getopt.ListLong("file", 'f', "")
	projectName := getopt.StringLong("project-name", 'p', "")
	profileName := getopt.StringLong("profile", 0, "")
	verboseFlag := getopt.BoolLong("verbose", 0, "")
	logLevel := getopt.StringLong("log-level", 0, "")
	noANSI := getopt.BoolLong("no-ansi", 0, "")
	version := getopt.BoolLong("version", 'v', "")
	host := getopt.StringLong("host", 'H', "")
	tlsFlag := getopt.BoolLong("tls", 0, "")
	tlsCACert := getopt.StringLong("tlscacert", 0, "")
	tlsCert := getopt.StringLong("tlscert", 0, "")
	tlsKey := getopt.StringLong("tlskey", 0, "")
	tlsVerifyFlag := getopt.BoolLong("tlsverify", 0, "")
	skipHostname := getopt.BoolLong("skip-hostname-check", 0, "")
	projectDirectory := getopt.StringLong("project-directory", 0, "")
	compatibilityFlag := getopt.BoolLong("compatibility", 0, "")

	err := getopt.Getopt(nil)
	if err != nil {
		// Assume a bad flag has been passed in.
		fmt.Println(help())
		os.Exit(0)
	}

	if *helpFlag {
		fmt.Print(help())
		os.Exit(0)
	}

	// Sanitise input.

	if len(*files) == 0 {
		// set default if nothing supplied.
		*files = []string{"docker-compose.yml"}
	}

	if *projectName == "" {
		// default project name is the current directory if nothing supplied.
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		dirRune := []rune(dir)
		if dirRune[len(dirRune)-1] == os.PathSeparator {
			dir = string(dirRune[:len(dirRune)-1])
		}

		dirs := strings.Split(dir, string(os.PathSeparator))
		*projectName = dirs[len(dirs)-1]
	}

	if *logLevel != "" && strings.ToUpper(*logLevel) != "DEBUG" && strings.ToUpper(*logLevel) != "INFO" && strings.ToUpper(*logLevel) != "WARNING" && strings.ToUpper(*logLevel) != "ERROR" && strings.ToUpper(*logLevel) != "CRITICAL" {
		fmt.Println(help())
		os.Exit(0)
	}

	// Collate input into a structure that can be passed around.
	input := Input{
		files:         *files,
		project:       *projectName,
		profile:       *profileName,
		verbose:       *verboseFlag,
		logLevel:      *logLevel,
		ansi:          *noANSI,
		version:       *version,
		host:          *host,
		tls:           *tlsFlag,
		tlsCACert:     *tlsCACert,
		tlsCert:       *tlsCert,
		tlsKey:        *tlsKey,
		tlsVerify:     *tlsVerifyFlag,
		skipHostname:  *skipHostname,
		projectDir:    *projectDirectory,
		compatibility: *compatibilityFlag,
		args:          getopt.Args(),
	}
}

func help() string {
	// https://docs.docker.com/compose/reference/
	return `Define and run multi-container applications with Docker.

Usage:
  go-podman-compose [-f <arg>...] [--profile <name>...] [options] [COMMAND] [ARGS...]
  go-podman-compose -h|--help

Options:
  -f, --file FILE             Specify an alternate compose file
                              (default: docker-compose.yml)
  -p, --project-name NAME     Specify an alternate project name
                              (default: directory name)
  --profile NAME              Specify a profile to enable
  --verbose                   Show more output
  --log-level LEVEL           DEPRECATED and not working from 2.0 - Set log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)
  --no-ansi                   Do not print ANSI control characters
  -v, --version               Print version and exit
  -H, --host HOST             Daemon socket to connect to

  --tls                       Use TLS; implied by --tlsverify
  --tlscacert CA_PATH         Trust certs signed only by this CA
  --tlscert CLIENT_CERT_PATH  Path to TLS certificate file
  --tlskey TLS_KEY_PATH       Path to TLS key file
  --tlsverify                 Use TLS and verify the remote
  --skip-hostname-check       Don't check the daemon's hostname against the
                              name specified in the client certificate
  --project-directory PATH    Specify an alternate working directory
                              (default: the path of the Compose file)
  --compatibility             If set, Compose will attempt to convert deploy
                              keys in v3 files to their non-Swarm equivalent

Commands:
  build              Build or rebuild services
  bundle             Generate a Docker bundle from the Compose file
  config             Validate and view the Compose file
  create             Create services
  down               Stop and remove containers, networks, images, and volumes
  events             Receive real time events from containers
  exec               Execute a command in a running container
  help               Get help on a command
  images             List images
  kill               Kill containers
  logs               View output from containers
  pause              Pause services
  port               Print the public port for a port binding
  ps                 List containers
  pull               Pull service images
  push               Push service images
  restart            Restart services
  rm                 Remove stopped containers
  run                Run a one-off command
  scale              Set number of containers for a service
  start              Start services
  stop               Stop services
  top                Display the running processes
  unpause            Unpause services
  up                 Create and start containers
  version            Show the Docker Compose version information`

}
