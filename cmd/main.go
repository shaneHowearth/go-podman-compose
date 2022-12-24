package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pborman/getopt/v2"
)

func main() {
	helpFlag := getopt.BoolLong("help", 'h', "display help")
	files := getopt.ListLong("file", 'f', "")
	projectName := getopt.StringLong("project-name", 'p', "")
	profileName := getopt.StringLong("profile", 0, "")
	verboseFlag := getopt.BoolLong("verbose", 0, "")
	err := getopt.Getopt(nil)
	if err != nil {
		fmt.Println(help())
		os.Exit(0)
	}
	if *helpFlag {
		fmt.Print(help())
	}
	if len(*files) == 0 {
		// default
		*files = []string{"docker-compose.yml"}
	}
	if *projectName == "" {
		// default
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		// Note: Could use strings.TrimSuffix here, but converting to []rune and
		// back isn't expensive IMO
		dirRune := []rune(dir)
		if dirRune[len(dirRune)-1] == os.PathSeparator {
			dir = string(dirRune[:len(dirRune)-1])
		}
		dirs := strings.Split(dir, string(os.PathSeparator))
		*projectName = dirs[len(dirs)-1]
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
