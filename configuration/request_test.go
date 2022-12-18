package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestRequestConf(t *testing.T) {
	const data = `
requests:
  - name: Request 1
    method: GET
    path: /your/request/path
    parameters:
      studentName: cloud
      age: 6
    headers:
      Content-Length: 0
  - name: Request 2
    method: POST
    path: /your/request/path
    body:
      type: JSON
      jsonValue: |
        {
          "name: "cloud",
          "age": 6
        }
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
  - name: Request 3
    method: POST
    path: /your/request/path
    body:
      type: TEXT
      textValue: "this is the text body"
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
  - name: Request 4
    method: POST
    path: /your/request/path
    body:
      type: RAW
      rawValue: "This is the raw body"
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
  - name: Request 5
    method: POST
    path: /your/request/path
    body:
      type: BINARY
      binaryValue: "This is the binary body"
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
  - name: Request 6
    method: POST
    path: /your/request/path
    body:
      type: FORM
      formValue:
        name: "cloud"
        address: "unknown"
        family:
          - father
          - mother
          - sister
          - brother
        file: "@filename"
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
  - name: Request 7
    method: POST
    path: /your/request/path
    body:
      type: XML
      xmlValue: <xmlroot><xmlelement>xml value</xmlelement></xmlroot>
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
  - name: Request 8
    method: POST
    path: /your/request/path
    body:
      type: URLENCODED
      urlencodedValue:
        name: cloud
        age: 6
    headers:
      Accept: text/plain
      Accept-Encoding: gzip
`
	configuration := Configuration{}
	err := yaml.Unmarshal([]byte(data), &configuration)
	if err != nil {
		t.Fatal(err)
	}

	expected := Configuration{
		RequestConfig: []RequestConf{
			{
				Name:   "Request 1",
				Method: "GET",
				Path:   "/your/request/path",
				Parameters: Parameter{
					"studentName": "cloud",
					"age":         "6",
				},
				Headers: Header{
					"Content-Length": "0",
				},
			},
			{
				Name:   "Request 2",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type:      "JSON",
					JsonValue: "{\n  \"name: \"cloud\",\n  \"age\": 6\n}\n",
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
			{
				Name:   "Request 3",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type:      "TEXT",
					TextValue: "this is the text body",
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
			{
				Name:   "Request 4",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type:     "RAW",
					RawValue: "This is the raw body",
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
			{
				Name:   "Request 5",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type:        "BINARY",
					BinaryValue: "This is the binary body",
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
			{
				Name:   "Request 6",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type: "FORM",
					FormValue: KeyPairValue{
						"name":    "cloud",
						"address": "unknown",
						"family": []interface{}{
							"father",
							"mother",
							"sister",
							"brother",
						},
						"file": "@filename",
					},
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
			{
				Name:   "Request 7",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type:     "XML",
					XMLValue: "<xmlroot><xmlelement>xml value</xmlelement></xmlroot>",
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
			{
				Name:   "Request 8",
				Method: "POST",
				Path:   "/your/request/path",
				Body: Body{
					Type: "URLENCODED",
					URLEncodedValue: KeyPairValue{
						"name": "cloud",
						"age":  6,
					},
				},
				Headers: Header{
					"Accept":          "text/plain",
					"Accept-Encoding": "gzip",
				},
			},
		},
	}

	assert.EqualValues(t, expected, configuration, "Parsed requests don't equal to the expected")
}
