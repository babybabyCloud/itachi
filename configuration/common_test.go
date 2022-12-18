package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestVariables(t *testing.T) {
	const data = `
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
		VariableConfig: VariableConf{
			"variableName1": "variable value 1",
			"variableName2": "variable value 2",
		},
	}

	assert.EqualValues(t, expected, configuration, "Parsed variables don't equal to the expected")
}
