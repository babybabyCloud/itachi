package configuration

// Configuration is a model from the configuration files
type Configuration struct {
	// IncludeConfig represent the model of IncludeConfig.
	IncludeConfig []IncludeConf `yaml: "includes`

	// EnvironmentConfig represent the model of environments, this can be used to store configurations of different
	// environments
	EnvironmentConfig []EnvironmentConf `yaml: "environments"`

	// VariableConfig represent the model of variables, this can be used to define global environment variables
	VariableConfig VariableConf `yaml: "variables"`
}
