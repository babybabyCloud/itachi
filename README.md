# Table of Contents
- [Table of Contents](#table-of-contents)
- [itachi](#itachi)
  - [Configuration reference](#configuration-reference)
  - [Command Line Reference](#command-line-reference)
    - [Run tests](#run-tests)
      - [Run all tests](#run-all-tests)
      - [Run all tests in specific configurations](#run-all-tests-in-specific-configurations)
      - [Run all tests in a specific environment](#run-all-tests-in-a-specific-environment)
      - [Run tests with variables](#run-tests-with-variables)
      - [Run tests filtering with tags](#run-tests-filtering-with-tags)
    - [Subcommand](#subcommand)
      - [run](#run)
        - [Format](#format)
        - [options](#options)
        - [file](#file)
      - [check](#check)


# itachi

This project aims to build a HTTP automation testing tool. The users can easily use the YAML file as the configuration. 

---
## Configuration reference
```YAML
includes:
  - type: file
    path: /your/configuration/path
  - type: http
    path: https://raw.githubusercontent.com/babybabyCloud/itachi/master/itachi.yml

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

variables:
  variableName1: variable value 1
  variableName2: variable value 2

client:
  - name: Client name
    auth:
      - method: The authorization method of this client
        value: The authorization value
    type: HTTP
    method: GET
    path: /your/request/path
    parameters:
      - parameter_name_1: parameter value 1
        parameter_name_2: parameter value 2
        parameter_name_3: parameter value 3
    body:
      - type: JSON
        values:
          - json_value: |
              {
                "json key 1": "json value 1",
                "json key 2": "json value 2"
              }
          - text_value: This is the text body
          - raw_value: This is a raw body
          - binary_value: This is a binary body
          - form_value:
              - form_key_1: form value 1
                form_key_2: form value 2
                form_key_3:
                  - form value 3 in a list
                  - form value 3 in a list, the 2nd item
                  - form value 3 in a list, the 3rd item
                form_key_4:
                  file: The file path
          - xml_value: |
              <xmlroot>
                <xmlelement>
                  xml value
                </xmlelement>
              </xmlroot>
          - urlencoded_value:
              - urlencoded_key_1: urlencoded value 1
              - urlencoded_key_2: urlencoded value 2
    headers:
      - header_key_1: header value 1
        header_key_2: header value 2

scenarios:
  - name: Get value from a remote server
    description: Get a value by using GET request from a remote server, then extract the value from the response
    steps:
      - name: Send a request
        description: Send request to the remote server
        client:
          - reference: Client name

```
---
| Key | Value Type | Optional | Choice | Description |
| --- | ---------- | -------- | ------ | ----------- |
| includes | list | Yes |  | To include another configuration file |
| includes.type | string | No | file, http | The type to be included |
| includes.path | string | No |  | The path to be included |
| environments | dict | No |   | Define the basic infomations for different environments |
| environments.name | string | No |  | The name of this environment |
| environments.domain | string | No |  | The base URL with the prefix protocal of your remote sever for test. For example, http://your.doman |
| environments.variables | dict | Yes |  | The environment variables that can be used in the suites |
| variables | dict | Yes |  | The global variables that can be used in the suites |
| client | list | Yes |  | The HTTP client used in the following test steps, this can be referred in scenarios.steps.client |
| client.auth | dict | Yes |  | Specify the authorization part of a client |
| client.auth.method | string | No | Basic, Bearer | Specify the authorization type |
| client.auth.value | string | No |  | Specify the authorization value. For "Basic" type, it requires this value in the format of "username:password", it will be encoded using Base64. For Bearer type, it requires only the authorization value, it will be added "Bearer " to the head of this value. |
| client.name | string | Yes |  | The name of an HTTP client, this is used in the scenarios.steps.client.refer. This is not requires in section "scenarios.steps.client" |
| client.type | string | No | HTTP | The type of a client, only HTTP is supported now |
| client.method | string | No | GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE | The type of a client |
| client.path | string | No |  | The path of your request to access. It supports including the path variable. For example, set the path to "/pet/{id}, the value of id should be set in client.paremeters |
| client.parameters | dict | Yes |  | The HTTP requests parameters |
| client.body | dict | Yes |  | The HTTP requests body |
| client.body.type | string | No | JSON, TEXT, RAW, BINARY, FORM, URLENCODED,XML | The type of HTTP requests body |
| client.header | dict | No |  | The headers of HTTP requests body |
| scenarios | list | No |  | The environment variables that can be used in the scenarios |
| scenarios.name | string | No |  | The name of a scenario |
| scenarios.description | string | Yes |  | The description of a scenario |
| scenarios.steps | list | No |  | The test steps of this scenario |
| scenarios.steps.name | string | No |  | The name of this step |
| scenarios.steps.description | string | Yes |  | The description of this step |
| scenarios.steps.client |  |  |  | The same the configuration of client part. |

## Command Line Reference

### Run tests
#### Run all tests
```bash
$: itachi run
```

#### Run all tests in specific configurations
```bash
$: itachi run configuration1.yaml configuration2.yaml
```

#### Run all tests in a specific environment
```bash
$: itachi run --env dev
```

#### Run tests with variables
```bash
$: itachi run --var "name=value" --var "name=value"
```

#### Run tests filtering with tags
```bash
$: itachi run --tag=tag1 --tag=tag2
```

---

### Subcommand
#### run
##### Format 
```bash
$: itachi run [options] [file...]
```
##### options
* --env=name Specify the environment to use
* --var="key=value" Add or overwrite the value of a variable
* --tag=name Specify the scenarios or steps with the name of tags to run

##### file
Specify which configuration file to run. If not specify, the current directory must contain "main.yaml" or "main.yml"

#### check