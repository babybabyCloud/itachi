package configuration

type Environment struct {
	Name      string   `yaml: "name"`
	Domain    string   `yaml: "domain"`
	Variables Variable `yaml: "variables"`
}

// Environments represents the model of environments, this can be used to store configurations of different environments
type Environments struct {
	Environments []Environment `yaml: "environments"`
}
