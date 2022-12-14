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
	variables := Variables{}
	err := yaml.Unmarshal([]byte(data), &variables)

	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := Variables{
		Variables: Variable{
			"variableName1": "variable value 1",
			"variableName2": "variable value 2",
		},
	}
	assert.Equal(t, variables, expected, "Parsed variables should equal to the expected")
}
