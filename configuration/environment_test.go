package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestEnvironment(t *testing.T) {
	const data = `
environments:
  - name: dev
    domain: https://your.dev.domain
    variables:
      variableName1: variable value 1
      variableName2: variable value 2
  - name: qc
    domain: https://your.qc.domain
    variables:
      variableName1: variable value 1
      variableName2: variable value 2
`

	configuration := Configuration{}

	err := yaml.Unmarshal([]byte(data), &configuration)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := Configuration{
		EnvironmentConfig: []EnvironmentConf{
			{
				Name:   "dev",
				Domain: "https://your.dev.domain",
				Variables: VariableConf{
					"variableName1": "variable value 1",
					"variableName2": "variable value 2",
				},
			},
			{
				Name:   "qc",
				Domain: "https://your.qc.domain",
				Variables: VariableConf{
					"variableName1": "variable value 1",
					"variableName2": "variable value 2",
				},
			},
		},
	}

	assert.Equal(t, configuration, expected, "Parsed environments should equal to the expected.")
}
