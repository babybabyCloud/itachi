package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestInclude(t *testing.T) {
	const data = `
includes:
  - type: file
    path: /your/configuration/path
  - type: http
    path: https://raw.githubusercontent.com/babybabyCloud/itachi/master/itachi.yml
`
	configuration := Configuration{}

	err := yaml.Unmarshal([]byte(data), &configuration)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := Configuration{
		IncludeConfig: []IncludeConf{
			{
				Type: "file",
				Path: "/your/configuration/path",
			},
			{
				Type: "http",
				Path: "https://raw.githubusercontent.com/babybabyCloud/itachi/master/itachi.yml",
			},
		},
	}

	assert.Equal(t, expected, configuration, "Parsed includes don't equal to the expected")
}
