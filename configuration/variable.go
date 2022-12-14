package configuration

type Variable map[string]string

type Variables struct {
	Variables Variable `yaml: "variables"`
}
