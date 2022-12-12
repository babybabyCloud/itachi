# itachi

## Configuration Reference

```YAML
environment:
  - name: dev
    variables:
      - variableName1: variable value 1
        variableName2: variable value 2
  - name: qc
    variables:
      - variableName1: variable value 1
        variableName2: variable value 2

scenarios:
  - name: Get value from a remote server
    description: Get a value by using GET request from a remote server, then extract the value from the response
    steps:
      - name: Send a request
        description: Send request to the remote server
        client: 
          - type: HTTP
            method: GET
```

| Key | Value Type | Optional | Choice | Description |
| --- | ---------- | -------- | ------ | ----------- |
| environment | dict | No |   | Define the basic infomations for different environments |
| environment.name | string | No |  | The name of this environment |
| environment.variables | dict | Yes |  | The environment variables that can be used in the suites |
| scenarios | list | No |  | The environment variables that can be used in the scenarios |
| scenarios.name | string | No |  | The name of a scenario |
| scenarios.description | string | Yes |  | The description of a scenario |
| scenarios.steps | list | No |  | The test steps of this scenario |
| scenarios.steps.name | string | No |  | The name of this step |
| scenarios.steps.description | string | Yes |  | The description of this step |
| scenarios.steps.client | dict | No |  | The client that is used in this step |
| scenarios.steps.client.type | string | No | HTTP | The type of a client, only HTTP is supported now |
| scenarios.steps.client.method | string | No | GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE | The type of a client |
