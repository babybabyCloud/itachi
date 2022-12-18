package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestClientConf(t *testing.T) {
	const data = `
client:
  - name: Client 1
    auth:
        method: Basic
        value: username:password
    domain: https://your.domain
    timeout: 30m
    headers:
        header_key_1: header value 1
        header_key_2: header value 2
  - name: Client 2
    auth:
        method: Bearer
        value: token
    domain: https://your.domain
    timeout: 30m
    headers:
        header_key_1: header value 1
        header_key_2: header value 2
    `
	configuration := Configuration{}
	err := yaml.Unmarshal([]byte(data), &configuration)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expected := Configuration{
		ClientConfig: []ClientConf{
			{
				Name: "Client 1",
				Auth: AuthConf{
					Method: "Basic",
					Value:  "username:password",
				},
				Domain:  "https://your.domain",
				Timeout: "30m",
				Headers: Header{
					"header_key_1": "header value 1",
					"header_key_2": "header value 2",
				},
			},
			{
				Name: "Client 2",
				Auth: AuthConf{
					Method: "Bearer",
					Value:  "token",
				},
				Domain:  "https://your.domain",
				Timeout: "30m",
				Headers: Header{
					"header_key_1": "header value 1",
					"header_key_2": "header value 2",
				},
			},
		},
	}

	assert.EqualValues(t, expected, configuration, "Parsed clients don't equal to the expected")
}
