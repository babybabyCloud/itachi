package configuration

// EnvironmentConf is the model of environments in a configuration file
type EnvironmentConf struct {

	// Name is the key "name"
	Name string `yaml:"name"`

	// Domain is the key "domain"
	Domain string `yaml:"domain"`

	// Variables are the key "variables"
	Variables VariableConf `yaml:"variables"`
}
