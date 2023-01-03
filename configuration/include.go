package configuration

import (
	"fmt"
	"os"
)

const (

	// FILE is the constant of include configuration with type "file"
	FILE = "file"

	// HTTP is the constant of include configuration with type "http"
	HTTP = "http"
)

// IncludeConf is the model of includes in a configuration file
type IncludeConf struct {

	// Type is the key "type"
	Type string `yaml:"type"`

	// Path is the key "path"
	Path string `yaml:"path"`
}

// Includer
type Includer interface {
	read() []byte
}

// FileInclude
type FileInclude struct {
	IncludeConf
}

// read reads a file configuration from local directory
func (inc *FileInclude) read() []byte {
	data, err := os.ReadFile(inc.Path)
	if err != nil {
		// TODO repalce panic with self-defined error message and exit
		panic(err)
	}

	return data
}

// HTTPInclude
type HTTPInclude struct {
	IncludeConf
}

func (inc *HTTPInclude) read() []byte {
	// TODO
	return []byte{}
}

// NewInclude creates a new Includer according to IncludeConf
func NewInclude(conf IncludeConf) Includer {
	switch conf.Type {
	case FILE:
		return &FileInclude{
			IncludeConf: conf,
		}
	case HTTP:
		return &HTTPInclude{
			IncludeConf: conf,
		}
	default:
		// TODO repalce panic with self-defined error message and exit
		panic(fmt.Sprintf("Unrecognized type of includes, possible values are '%s', '%s'", FILE, HTTP))
	}
}
