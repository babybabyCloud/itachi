package configuration

import (
	"container/list"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Configuration is a model from the configuration files
type Configuration struct {
	// IncludeConfig represent the model of IncludeConfig.
	IncludeConfig []IncludeConf `yaml:"includes"`

	// EnvironmentConfig represent the model of environments, this can be used to store configurations of different
	// environments
	EnvironmentConfig []EnvironmentConf `yaml:"environments"`

	// VariableConfig represent the model of variables, this can be used to define global environment variables
	VariableConfig VariableConf `yaml:"variables"`

	// ClientConfig represent the model of client.
	ClientConfig []ClientConf `yaml:"clients"`

	// RequestConfig represent the model of requests
	RequestConfig []RequestConf `yaml:"requests"`

	environmentConfig EnvironmentConfs

	clientConfig ClientConfs

	requestConfig RequestConfs
}

// NewConfiguration uses YAML file content to create Configuration
func NewConfiguration(data []byte) *Configuration {
	config := Configuration{}
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		// TODO handle the errer
		fmt.Printf("********************%s\n", err)
	}

	// Environments
	config.environmentConfig = make(EnvironmentConfs, len(config.EnvironmentConfig))
	for _, item := range config.EnvironmentConfig {
		config.environmentConfig[item.Name] = item
	}

	// Variables
	if config.VariableConfig == nil {
		config.VariableConfig = make(VariableConf, 0)
	}

	// Clients
	config.clientConfig = make(ClientConfs, len(config.ClientConfig))
	for _, item := range config.ClientConfig {
		config.clientConfig[item.Name] = item
	}

	// Requests
	config.requestConfig = make(RequestConfs, len(config.RequestConfig))
	for _, item := range config.RequestConfig {
		config.requestConfig[item.Name] = item
	}

	return &config
}

// Start starts to complete the configuration
func (config *Configuration) Start() {
	container := list.New()
	element := container.PushBack(config)
	for element != nil {
		item := element.Value.(*Configuration)
		for _, include := range item.IncludeConfig {
			otherConfig := NewConfiguration(NewInclude(include).read())
			config.merge(otherConfig)
			container.PushBack(otherConfig)
		}
		element = element.Next()
	}
}

func (config *Configuration) merge(otherConfig *Configuration) {

	// Merge EnvironmentConfig
	config.environmentConfig.merge(otherConfig.environmentConfig)

	// Merge VariableConfig
	config.VariableConfig.merge(otherConfig.VariableConfig)

	// Merge ClientConfig
	config.clientConfig.merge(otherConfig.clientConfig)

	// Merge RequestConfig
	config.requestConfig.merge(otherConfig.requestConfig)

}
