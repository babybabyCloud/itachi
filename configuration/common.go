package configuration

// VariableConf is the key "variables"
type VariableConf map[string]string

func (v VariableConf) merge(other VariableConf) {
	for key, value := range other {
		v[key] = value
	}
}

// Header is the key "headers"
type Header map[string]string

// Parameter is the key "parameter"
type Parameter map[string]string
