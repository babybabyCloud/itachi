package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationInclude(t *testing.T) {
	data := []byte(`
includes:
  - type: file
    path: /file/path
  - type: http
    path: https://domain.com/file/path
environments:
  - name: env 1
    domain: domain 1
clients:
  - name: client 1
    auth:
      method: Basic
      value: some value
requests:
  - name: request 1
    method: GET
`)
	config := NewConfiguration(data)

	expectedEnvironment := []EnvironmentConf{
		{
			Name:   "env 1",
			Domain: "domain 1",
		},
	}
	expectedClient := []ClientConf{
		{
			Name: "client 1",
			Auth: AuthConf{
				Method: "Basic",
				Value:  "some value",
			},
		},
	}
	expectedRequest := []RequestConf{
		{
			Name:   "request 1",
			Method: "GET",
		},
	}

	expected := Configuration{
		IncludeConfig: []IncludeConf{
			{
				Type: FILE,
				Path: "/file/path",
			},
			{
				Type: HTTP,
				Path: "https://domain.com/file/path",
			},
		},
		EnvironmentConfig: expectedEnvironment,
		ClientConfig:      expectedClient,
		RequestConfig:     expectedRequest,
		VariableConfig:    make(VariableConf, 0),
		environmentConfig: EnvironmentConfs{
			"env 1": expectedEnvironment[0],
		},
		clientConfig: ClientConfs{
			"client 1": expectedClient[0],
		},
		requestConfig: RequestConfs{
			"request 1": expectedRequest[0],
		},
	}

	assert.EqualValues(t, expected, *config)
}

func TestMerge(t *testing.T) {
	config1Env1 := EnvironmentConf{
		Name:   "Env 1",
		Domain: "https://domain1",
		Variables: VariableConf{
			"Variable 1": "Variable 1 value",
			"Variable 2": "Variable 2 value",
		},
	}
	config1Env2 := EnvironmentConf{
		Name:   "Env 2",
		Domain: "https://domain2",
		Variables: VariableConf{
			"Variable a": "Variable a value",
			"Variable b": "Variable b value",
		},
	}
	config1Client1 := ClientConf{
		Name: "client 1",
		Auth: AuthConf{
			Method: "BASIC",
			Value:  "Auth value 1",
		},
		Domain:  "http://client1.domain",
		Timeout: "10s",
		Headers: Header{
			"Header 1": "Header 1 value",
			"Header 2": "Header 2 value",
		},
	}
	config1Client2 := ClientConf{
		Name: "client 2",
		Auth: AuthConf{
			Method: "Bearer",
			Value:  "Auth value ",
		},
		Domain:  "http://client2.domain",
		Timeout: "10m",
		Headers: Header{
			"Header 1": "Header 1 value",
			"Header 2": "Header 2 value",
		},
	}
	config1Requet1 := RequestConf{
		Name:   "Request 1",
		Method: "GET",
		Path:   "/request/1/path",
		Parameters: Parameter{
			"Parameter 1": "Parameter value 1",
			"Parameter 2": "Parameter value 2",
		},
		Body: Body{
			Type: "JSON",
			JsonValue: `{
				"json key 1": "json value 1",
				"json key 2": "json value 2",
			}`,
		},
	}
	config1Requet2 := RequestConf{
		Name:   "Request 2",
		Method: "POST",
		Path:   "/request/2/path",
		Parameters: Parameter{
			"Parameter 1": "Parameter value 1",
			"Parameter 2": "Parameter value 2",
		},
		Body: Body{
			Type:      "Text",
			TextValue: "Text value",
		},
	}
	config1 := &Configuration{
		EnvironmentConfig: []EnvironmentConf{
			config1Env1,
			config1Env2,
		},
		VariableConfig: VariableConf{
			"Variable key 1": "Variable key 1 value",
			"Variable key 2": "Variable key 2 value",
		},
		ClientConfig: []ClientConf{
			config1Client1,
			config1Client2,
		},
		RequestConfig: []RequestConf{
			config1Requet1,
			config1Requet2,
		},

		environmentConfig: EnvironmentConfs{
			config1Env1.Name: config1Env1,
			config1Env2.Name: config1Env2,
		},

		clientConfig: ClientConfs{
			config1Client1.Name: config1Client1,
			config1Client2.Name: config1Client2,
		},

		requestConfig: RequestConfs{
			config1Requet1.Name: config1Requet1,
			config1Requet2.Name: config1Requet2,
		},
	}

	config2Env1 := EnvironmentConf{
		Name:   "Env 1",
		Domain: "https://domain11",
		Variables: VariableConf{
			"Variable 1": "Variable 1 value",
			"Variable 2": "Variable 2 value",
		},
	}
	config2Env2 := EnvironmentConf{
		Name:   "Env 3",
		Domain: "https://domain21",
		Variables: VariableConf{
			"Variable a": "Variable a value",
			"Variable b": "Variable b value",
		},
	}
	config2Client1 := ClientConf{
		Name: "client 1",
		Auth: AuthConf{
			Method: "BASIC",
			Value:  "Auth value 11",
		},
		Domain:  "http://client1.domain",
		Timeout: "10s",
		Headers: Header{
			"Header 1": "Header 1 value",
			"Header 2": "Header 2 value",
		},
	}
	config2Client2 := ClientConf{
		Name: "client 3",
		Auth: AuthConf{
			Method: "Bearer",
			Value:  "Auth value 22",
		},
		Domain:  "http://client2.domain",
		Timeout: "10m",
		Headers: Header{
			"Header 1": "Header 1 value",
			"Header 2": "Header 2 value",
		},
	}
	config2Requet1 := RequestConf{
		Name:   "Request 1",
		Method: "GET",
		Path:   "/request/11/path",
		Parameters: Parameter{
			"Parameter 1": "Parameter value 1",
			"Parameter 2": "Parameter value 2",
		},
		Body: Body{
			Type: "JSON",
			JsonValue: `{
				"json key 1": "json value 1",
				"json key 2": "json value 2",
			}`,
		},
	}
	config2Requet2 := RequestConf{
		Name:   "Request 3",
		Method: "POST",
		Path:   "/request/22/path",
		Parameters: Parameter{
			"Parameter 1": "Parameter value 1",
			"Parameter 2": "Parameter value 2",
		},
		Body: Body{
			Type:      "Text",
			TextValue: "Text value",
		},
	}
	config2 := Configuration{
		EnvironmentConfig: []EnvironmentConf{
			config2Env1,
			config2Env2,
		},
		VariableConfig: VariableConf{
			"Variable key 1": "Variable key 11 value",
			"Variable key 3": "Variable key 3 value",
		},
		ClientConfig: []ClientConf{
			config2Client1,
			config2Client2,
		},
		RequestConfig: []RequestConf{
			config2Requet1,
			config2Requet2,
		},

		environmentConfig: EnvironmentConfs{
			config2Env1.Name: config2Env1,
			config2Env2.Name: config2Env2,
		},

		clientConfig: ClientConfs{
			config2Client1.Name: config2Client1,
			config2Client2.Name: config2Client2,
		},

		requestConfig: RequestConfs{
			config2Requet1.Name: config2Requet1,
			config2Requet2.Name: config2Requet2,
		},
	}

	config1.merge(&config2)

	expected := Configuration{
		EnvironmentConfig: []EnvironmentConf{
			config1Env1,
			config1Env2,
		},
		VariableConfig: VariableConf{
			"Variable key 1": "Variable key 11 value",
			"Variable key 2": "Variable key 2 value",
			"Variable key 3": "Variable key 3 value",
		},
		ClientConfig: []ClientConf{
			config1Client1,
			config1Client2,
		},
		RequestConfig: []RequestConf{
			config1Requet1,
			config1Requet2,
		},

		environmentConfig: EnvironmentConfs{
			config1Env1.Name: config2Env1,
			config1Env2.Name: config1Env2,
			config2Env2.Name: config2Env2,
		},

		clientConfig: ClientConfs{
			config1Client1.Name: config2Client1,
			config1Client2.Name: config1Client2,
			config2Client2.Name: config2Client2,
		},

		requestConfig: RequestConfs{
			config1Requet1.Name: config2Requet1,
			config1Requet2.Name: config1Requet2,
			config2Requet2.Name: config2Requet2,
		},
	}

	assert.EqualValues(t, expected, *config1)
}

func TestConfigurationStart(t *testing.T) {
	data, err := os.ReadFile("../test_data/one.yaml")
	if err != nil {
		t.Fatal(err)
	}

	config := NewConfiguration(data)
	config.Start()

	expected := VariableConf{
		"variable": "value",
	}

	assert.EqualValues(t, expected, config.VariableConfig)
}
