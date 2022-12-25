package podmancomposer

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
