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

// EnvironmentConfs is the alias of map of EnvironmentConf
type EnvironmentConfs map[string]EnvironmentConf

func (c EnvironmentConfs) merge(other EnvironmentConfs) {
	for key, value := range other {
		c[key] = value
	}
}
