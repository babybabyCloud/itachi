package configuration

// AuthConf Method
const (
	BASIC  = "Basic"
	BEARER = "Bearer"
)

// ClientConf is the model of client in a configuration file
type ClientConf struct {

	// Name is the name of a client
	Name string `yaml:"name"`

	// Auth is the configuration to specify which authoriation to use
	Auth AuthConf `yaml:"auth"`

	// // Domain is the key "domain"
	Domain string `yaml:"domain"`

	// Timeout is the key "timeout"
	Timeout string `yaml:"timeout"`

	// Headers is the key "headers"
	Headers Header `yaml:"headers"`
}

// AuthConf is the model of client.auth in a configuration file
type AuthConf struct {
	// Method specifies which type of authorization to use, only supports "Basic" and "Bearer" now
	// Please refer to https://datatracker.ietf.org/doc/html/rfc7617 and https://datatracker.ietf.org/doc/html/rfc6750
	// for more details
	Method string `yaml:"method"`

	// Value is the credential to use
	Value string `yaml:"value"`
}
